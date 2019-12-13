package main

import (
	"encoding/json"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/bvh"
	"fluorescence/geometry/primitive/disk"
	"fluorescence/geometry/primitive/plane"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/geometry/primitive/rectangle"
	"fluorescence/geometry/primitive/sphere"
	"fluorescence/geometry/primitive/triangle"
	"fluorescence/shading"
	"fluorescence/shading/material"
	"fmt"
	"io/ioutil"
	"reflect"
)

type Parameters struct {
	ImageWidth      int            `json:"image_width"`
	ImageHeight     int            `json:"image_height"`
	FileType        string         `json:"file_type"`
	FileDirectory   string         `json:"file_directory"`
	FileName        string         `json:"file_name"`
	GammaCorrection int            `json:"gamma_correction"`
	SampleCount     int            `json:"sample_count"`
	MaxBounces      int            `json:"max_bounces"`
	UseBVH          bool           `json:"use_bvh"`
	BackgroundColor *shading.Color `json:"background_color"`
	TMin            float64        `json:"t_min"`
	TMax            float64        `json:"t_max"`
	Scene           *Scene         `json:"scene"`
}

type Scene struct {
	CameraName      string              `json:"camera_name"`
	Camera          *Camera             `json:"-"`
	ObjectMaterials []*ObjectMaterial   `json:"objects"`
	Objects         primitive.Primitive `json:"-"`
}

type ObjectMaterial struct {
	ObjectName   string `json:"object_name"`
	MaterialName string `json:"material_name"`
}

type CameraData struct {
	Name   string  `json:"name"`
	Camera *Camera `json:"data"`
}

type ObjectData struct {
	Name     string      `json:"name"`
	TypeName string      `json:"type"`
	Data     interface{} `json:"data"`
}

type MaterialData struct {
	Name     string      `json:"name"`
	TypeName string      `json:"type"`
	Data     interface{} `json:"data"`
}

func LoadConfigs(parametersFileName, camerasFileName, objectsFileName, materialsFileName string) (*Parameters, error) {

	totalCameras, err := loadCameras(camerasFileName)
	if err != nil {
		return nil, err
	}
	totalObjects, err := loadObjects(objectsFileName)
	if err != nil {
		return nil, err
	}
	totalMaterials, err := loadMaterials(materialsFileName)
	if err != nil {
		return nil, err
	}
	parameters, err := loadParameters(parametersFileName)
	if err != nil {
		return nil, err
	}

	selectedCamera, exists := totalCameras[parameters.Scene.CameraName]
	if !exists {
		return nil, fmt.Errorf("Selected Camera (%s) not in %s", parameters.Scene.CameraName, camerasFileName)
	}
	parameters.Scene.Camera = selectedCamera
	err = parameters.Scene.Camera.Setup(parameters)
	if err != nil {
		return nil, err
	}

	boundedSceneObjects := &primitivelist.PrimitiveList{}
	unboundedSceneObjects := &primitivelist.PrimitiveList{}
	for _, om := range parameters.Scene.ObjectMaterials {
		selectedObject, exists := totalObjects[om.ObjectName]
		if !exists {
			return nil, fmt.Errorf("Selected Object (%s) not in %s", om.ObjectName, objectsFileName)
		}
		selectedMaterial, exists := totalMaterials[om.MaterialName]
		if !exists {
			return nil, fmt.Errorf("Selected Material (%s) not in %s", om.MaterialName, materialsFileName)
		}
		if reflect.TypeOf(selectedMaterial) == reflect.TypeOf(&material.Dielectric{}) {
			objectType := reflect.TypeOf(selectedObject)
			triangleType := reflect.TypeOf(triangle.EmptyTriangle())
			rectangleType := reflect.TypeOf(rectangle.EmptyRectangle())
			planeType := reflect.TypeOf(plane.EmptyPlane())
			diskType := reflect.TypeOf(disk.EmptyDisk())
			hollowDiskType := reflect.TypeOf(disk.EmptyHollowDisk())
			if objectType == triangleType ||
				objectType == rectangleType ||
				objectType == planeType ||
				objectType == diskType ||
				objectType == hollowDiskType {
				return nil, fmt.Errorf("Cannot attach refractive or volumetric materials (%s) to non-closed geometry (%s)",
					om.MaterialName, om.ObjectName)
			}
		}
		newPrimitive := selectedObject.Copy()
		newPrimitive.SetMaterial(selectedMaterial)
		_, isBounded := newPrimitive.BoundingBox(0, 0)
		if isBounded {
			boundedSceneObjects.List = append(boundedSceneObjects.List, newPrimitive)
		} else {
			unboundedSceneObjects.List = append(unboundedSceneObjects.List, newPrimitive)
		}
	}

	if parameters.UseBVH {
		sceneBVH, err := bvh.NewBVH(boundedSceneObjects)
		if err != nil {
			return nil, err
		}
		if len(unboundedSceneObjects.List) == 0 {
			parameters.Scene.Objects = sceneBVH
		} else {
			rootNode := &primitivelist.PrimitiveList{
				List: append(unboundedSceneObjects.List, sceneBVH),
			}
			parameters.Scene.Objects = rootNode
		}
	} else {
		parameters.Scene.Objects = &primitivelist.PrimitiveList{
			List: append(boundedSceneObjects.List, unboundedSceneObjects.List...),
		}
	}

	return parameters, nil
}

func loadCameras(fileName string) (map[string]*Camera, error) {
	camerasBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var camerasData []*CameraData
	err = json.Unmarshal(camerasBytes, &camerasData)
	if err != nil {
		return nil, err
	}
	camerasMap := map[string]*Camera{}
	for _, cd := range camerasData {
		camerasMap[cd.Name] = cd.Camera
	}
	return camerasMap, nil
}

func loadObjects(fileName string) (map[string]primitive.Primitive, error) {
	objectsBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var objectsData []*ObjectData
	err = json.Unmarshal(objectsBytes, &objectsData)
	if err != nil {
		return nil, err
	}
	objectsMap := map[string]primitive.Primitive{}
	for _, o := range objectsData {
		switch o.TypeName {
		case "Disk":
			var dd disk.DiskData
			dataBytes, err := json.Marshal(o.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &dd)
			newDisk, err := disk.NewDisk(&dd)
			if err != nil {
				return nil, err
			}
			objectsMap[o.Name] = newDisk
		case "HollowDisk":
			var hdd disk.HollowDiskData
			dataBytes, err := json.Marshal(o.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &hdd)
			newHollowDisk, err := disk.NewHollowDisk(&hdd)
			if err != nil {
				return nil, err
			}
			objectsMap[o.Name] = newHollowDisk
		case "Plane":
			var pd plane.PlaneData
			dataBytes, err := json.Marshal(o.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &pd)
			newPlane, err := plane.NewPlane(&pd)
			if err != nil {
				return nil, err
			}
			objectsMap[o.Name] = newPlane
		case "Rectangle":
			var rd rectangle.RectangleData
			dataBytes, err := json.Marshal(o.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &rd)
			newRectangle, err := rectangle.NewRectangle(&rd)
			if err != nil {
				return nil, err
			}
			objectsMap[o.Name] = newRectangle
		case "Sphere":
			var sd sphere.SphereData
			dataBytes, err := json.Marshal(o.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &sd)
			newSphere, err := sphere.NewSphere(&sd)
			if err != nil {
				return nil, err
			}
			objectsMap[o.Name] = newSphere
		case "Triangle":
			var td triangle.TriangleData
			dataBytes, err := json.Marshal(o.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &td)
			newTriangle, err := triangle.NewTriangle(&td)
			if err != nil {
				return nil, err
			}
			objectsMap[o.Name] = newTriangle
		default:
			return nil, fmt.Errorf("Type (%s) not a valid primitive type", o.TypeName)
		}
	}
	return objectsMap, nil
}

func loadMaterials(fileName string) (map[string]material.Material, error) {
	materialsBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var materialsData []MaterialData
	err = json.Unmarshal(materialsBytes, &materialsData)
	if err != nil {
		return nil, err
	}
	materialsMap := map[string]material.Material{}
	for _, m := range materialsData {
		switch m.TypeName {
		case "Lambertian":
			var l material.Lambertian
			dataBytes, err := json.Marshal(m.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &l)
			materialsMap[m.Name] = &l
		case "Metal":
			var mtl material.Metal
			dataBytes, err := json.Marshal(m.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &mtl)
			materialsMap[m.Name] = &mtl
		case "Dielectric":
			var d material.Dielectric
			dataBytes, err := json.Marshal(m.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &d)
			materialsMap[m.Name] = &d
		default:
			return nil, fmt.Errorf("Type (%s) not a valid material type", m.TypeName)
		}
	}
	return materialsMap, nil
}

func loadParameters(fileName string) (*Parameters, error) {
	parametersBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var parameters Parameters
	err = json.Unmarshal(parametersBytes, &parameters)
	if err != nil {
		return nil, err
	}
	return &parameters, nil
}
