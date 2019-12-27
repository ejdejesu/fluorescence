package main

import (
	"encoding/json"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/box"
	"fluorescence/geometry/primitive/bvh"
	"fluorescence/geometry/primitive/cylinder"
	"fluorescence/geometry/primitive/disk"
	"fluorescence/geometry/primitive/hollowcylinder"
	"fluorescence/geometry/primitive/hollowdisk"
	"fluorescence/geometry/primitive/infinitecylinder"
	"fluorescence/geometry/primitive/plane"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/geometry/primitive/pyramid"
	"fluorescence/geometry/primitive/rectangle"
	"fluorescence/geometry/primitive/sphere"
	"fluorescence/geometry/primitive/transform/rotate"
	"fluorescence/geometry/primitive/transform/translate"
	"fluorescence/geometry/primitive/triangle"
	"fluorescence/geometry/primitive/uncappedcylinder"
	"fluorescence/shading"
	"fluorescence/shading/material"
	"fluorescence/shading/texture"
	"fmt"
	"io/ioutil"
	"reflect"
)

// Parameters holds top-level information about the program's execution and the image's properties
type Parameters struct {
	ImageWidth           int           `json:"image_width"`                // width of the image in pixels
	ImageHeight          int           `json:"image_height"`               // height of the image in pixels
	FileType             string        `json:"file_type"`                  // image file type (png, jpg, etc.)
	FileDirectory        string        `json:"file_directory"`             // folder of image to write
	Version              string        `json:"version"`                    // program version
	GammaCorrection      float64       `json:"gamma_correction"`           // how much gamma correction to perform on the image
	TextureGamma         float64       `json:"texture_gamma"`              // how much counter-gamma correction to apply to image textures
	UseScalingTruncation bool          `json:"use_scaling_truncation"`     // should the program truncate over-magnitude colors by scaling linearly as opposed to clamping?
	SampleCount          int           `json:"sample_count"`               // amount of samples to write
	TileWidth            int           `json:"tile_width"`                 // width of a tile in pixels
	TileHeight           int           `json:"tile_height"`                // height of a tile in pixels
	MaxBounces           int           `json:"max_bounces"`                // amount of reflections to check before giving up
	UseBVH               bool          `json:"use_bvh"`                    // should the program generate and use a Bounding Volume Hierarchy?
	BGColorMagnitude     float64       `json:"background_color_magnitude"` // amount to scale bg color by
	BackgroundColor      shading.Color `json:"background_color"`           // color to return when nothing is intersected
	TMin                 float64       `json:"t_min"`                      // minimum ray "time" to count intersection
	TMax                 float64       `json:"t_max"`                      // maximum ray "time" to count intersection
	SceneFileName        string        `json:"scene_file_name"`            // file name of scene config file
	Scene                *Scene        `json:"-"`                          // Scene reference
}

// Scene holds information about the pictured scene, such as the objects and camera
type Scene struct {
	Name            string              `json:"scene_name"`  // name of the scene
	CameraName      string              `json:"camera_name"` // name of the camera to use
	Camera          *Camera             `json:"-"`           // Camera reference
	ObjectMaterials []*ObjectMaterial   `json:"objects"`     // temporary reference to ObjectMaterials to link geometry to materials
	Objects         primitive.Primitive `json:"-"`           // reference to Objects in the scene
}

// ObjectMaterial is a temporary holding structure to link together geometry objects and materials
type ObjectMaterial struct {
	ObjectName   string `json:"object_name"`
	MaterialName string `json:"material_name"`
}

// CameraData holds a reference to the Camera struct and name
type CameraData struct {
	Name   string  `json:"name"`
	Camera *Camera `json:"data"`
}

// ObjectData holds information about a geometry object
type ObjectData struct {
	Name     string      `json:"name"`
	TypeName string      `json:"type"`
	Data     interface{} `json:"data"`
}

// MaterialData holds information about a material
type MaterialData struct {
	Name                   string      `json:"name"`
	TypeName               string      `json:"type"`
	ReflectanceTextureName string      `json:"reflectance_texture_name"`
	EmittanceTextureName   string      `json:"emittance_texture_name"`
	Data                   interface{} `json:"data"`
}

// TextureData holds information about a texture
type TextureData struct {
	Name     string      `json:"name"`
	TypeName string      `json:"type"`
	Data     interface{} `json:"data"`
}

// LoadConfigs reads and parses the config files for the program
func LoadConfigs(
	parametersFileName,
	camerasFileName,
	objectsFileName,
	materialsFileName,
	texturesFileName string) (*Parameters, error) {

	// load various json config files into their respective structs
	fmt.Printf("\tLoading Parameters...\n")
	parameters, err := loadParameters(parametersFileName)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\tLoading Cameras...\n")
	totalCameras, err := loadCameras(camerasFileName)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\tLoading Objects...\n")
	totalObjects, err := loadObjects(objectsFileName)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\tLoading Textures...\n")
	totalTextures, err := loadTextures(texturesFileName, parameters.TextureGamma)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\tLoading Materials...\n")
	totalMaterials, err := loadMaterials(materialsFileName, texturesFileName, totalTextures)
	if err != nil {
		return nil, err
	}

	// append the correct directory to the scene filename
	parameters.SceneFileName = "./config/scenes/" + parameters.SceneFileName
	// ...and load the scene
	parameters.Scene, err = loadScene(parameters.SceneFileName)
	if err != nil {
		return nil, err
	}

	// select the correct camera and initialize it
	selectedCamera, exists := totalCameras[parameters.Scene.CameraName]
	if !exists {
		return nil, fmt.Errorf("selected Camera (%s) not in %s", parameters.Scene.CameraName, camerasFileName)
	}
	parameters.Scene.Camera = selectedCamera
	err = parameters.Scene.Camera.Setup(parameters)
	if err != nil {
		return nil, err
	}

	// loop over the loosely connected ObjectMaterials and parse the proper materials into the primitives they represent

	// most geometry objects are "bounded" meaning an AABB (Axis-Aligned Bounding Box) can be placed around them.
	boundedSceneObjects := &primitivelist.PrimitiveList{}
	// some geometry objects, however, are infinite in nature, which mean they canned be bounded.
	// a distinction must be made between these to prevent assembling a BVH or other acceleration structure
	// without a bounding box around certain primitives
	unboundedSceneObjects := &primitivelist.PrimitiveList{}
	for _, om := range parameters.Scene.ObjectMaterials {
		// grab the labelled objects and materials
		selectedObject, exists := totalObjects[om.ObjectName]
		if !exists {
			return nil, fmt.Errorf("selected Object (%s) not in %s", om.ObjectName, objectsFileName)
		}
		selectedMaterial, exists := totalMaterials[om.MaterialName]
		if !exists {
			return nil, fmt.Errorf("selected Material (%s) not in %s", om.MaterialName, materialsFileName)
		}

		// this is a check to ensure that materials that have a transmission component (i.e. Dielectrics)
		// are not attached to "open" geometry, such as single-sided triangles and rectangles, so the
		// transmission commponent can be reversed
		// this is an arbitrary restriction that is likely to be removed in the future with the user choosing to self-restrict
		// themselves in a similar manner
		if reflect.TypeOf(selectedMaterial) == reflect.TypeOf(&material.Dielectric{}) {
			if !selectedObject.IsClosed() {
				return nil, fmt.Errorf("cannot attach refractive or volumetric materials (%s) to non-closed geometry (%s)",
					om.MaterialName, om.ObjectName)
			}
		}
		// copy the object so we don't override it's material if it is reused in the scene
		newPrimitive := selectedObject.Copy()
		newPrimitive.SetMaterial(selectedMaterial)
		// added to the cooresponding list based on type
		if newPrimitive.IsInfinite() {
			unboundedSceneObjects.List = append(unboundedSceneObjects.List, newPrimitive)
		} else {
			boundedSceneObjects.List = append(boundedSceneObjects.List, newPrimitive)
		}
	}

	// START MANUAL INSERT

	// for x := 0.5; x < 10.0; x++ {
	// 	for y := 0.5; y < 10.0; y++ {
	// 		for z := -0.5; z > -10.0; z-- {
	// 			newSphere, err := sphere.New(&sphere.Data{
	// 				Center: geometry.Point{
	// 					X: x,
	// 					Y: y,
	// 					Z: z,
	// 				},
	// 				Radius: 0.2,
	// 			})
	// 			if err != nil {
	// 				return nil, err
	// 			}
	// 			newSphere.SetMaterial(totalMaterials["glass"])
	// 			boundedSceneObjects.List = append(boundedSceneObjects.List, newSphere)
	// 		}
	// 	}
	// }

	// END MANUAL INSERT

	// if we are using a BVH ...
	if parameters.UseBVH {
		// ... construct it from the bounded objects ..
		sceneBVH, err := bvh.New(boundedSceneObjects)
		if err != nil {
			return nil, err
		}
		// ... and set it as the root node if no infinite geometry exists
		if len(unboundedSceneObjects.List) == 0 {
			parameters.Scene.Objects = sceneBVH
		} else {
			// but if some infinite geometry exists in the scene, we then
			// establish a new root node as a list of the BVH and the infinite geometry
			rootNode := &primitivelist.PrimitiveList{
				List: append(unboundedSceneObjects.List, sceneBVH),
			}
			parameters.Scene.Objects = rootNode
		}
	} else {
		// if we are not using a BVH, combine the lists into a core list and set it as the root node
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
		if _, ok := camerasMap[cd.Name]; ok {
			return nil, fmt.Errorf("camera (%s) redefined", cd.Name)
		}
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
		if _, ok := objectsMap[o.Name]; ok {
			return nil, fmt.Errorf("object (%s) redefined", o.Name)
		}
		newPrimitive, err := decodeObject(o.TypeName, o.Data)
		if err != nil {
			return nil, err
		}
		objectsMap[o.Name] = newPrimitive
	}
	return objectsMap, nil
}

func decodeObject(typeName string, data interface{}) (primitive.Primitive, error) {
	switch typeName {
	case "Box":
		var b box.Box
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &b)
		newBox, err := b.Setup()
		if err != nil {
			return nil, err
		}
		return newBox, nil
	case "Cylinder":
		var c cylinder.Cylinder
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &c)
		newCylinder, err := c.Setup()
		if err != nil {
			return nil, err
		}
		return newCylinder, nil
	case "HollowCylinder":
		var hc hollowcylinder.HollowCylinder
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &hc)
		newHollowCylinder, err := hc.Setup()
		if err != nil {
			return nil, err
		}
		return newHollowCylinder, nil
	case "InfiniteCylinder":
		var ic infinitecylinder.InfiniteCylinder
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &ic)
		newInfiniteCylinder, err := ic.Setup()
		if err != nil {
			return nil, err
		}
		return newInfiniteCylinder, nil
	case "UncappedCylinder":
		var uc uncappedcylinder.UncappedCylinder
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &uc)
		newUncappedCylinder, err := uc.Setup()
		if err != nil {
			return nil, err
		}
		return newUncappedCylinder, nil
	case "Disk":
		var d disk.Disk
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &d)
		newDisk, err := d.Setup()
		if err != nil {
			return nil, err
		}
		return newDisk, nil
	case "HollowDisk":
		var hd hollowdisk.HollowDisk
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &hd)
		newHollowDisk, err := hd.Setup()
		if err != nil {
			return nil, err
		}
		return newHollowDisk, nil
	case "Plane":
		var p plane.Plane
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &p)
		newPlane, err := p.Setup()
		if err != nil {
			return nil, err
		}
		return newPlane, nil
	case "Pyramid":
		var p pyramid.Pyramid
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &p)
		newPyramid, err := p.Setup()
		if err != nil {
			return nil, err
		}
		return newPyramid, nil
	case "Rectangle":
		var r rectangle.Rectangle
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &r)
		newRectangle, err := r.Setup()
		if err != nil {
			return nil, err
		}
		return newRectangle, nil
	case "Sphere":
		var s sphere.Sphere
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &s)
		newSphere, err := s.Setup()
		if err != nil {
			return nil, err
		}
		return newSphere, nil
	case "Triangle":
		var t triangle.Triangle
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &t)
		newTriangle, err := t.Setup()
		if err != nil {
			return nil, err
		}
		return newTriangle, nil
	case "Translation":
		var t translate.Translation
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &t)
		corePrimitive, err := decodeObject(t.TypeName, t.Data)
		if err != nil {
			return nil, err
		}
		t.Primitive = corePrimitive
		newTranslation, err := (&t).Setup()
		if err != nil {
			return nil, err
		}
		return newTranslation, nil
	case "RotationX":
		var rx rotate.RotationX
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &rx)
		corePrimitive, err := decodeObject(rx.TypeName, rx.Data)
		if err != nil {
			return nil, err
		}
		rx.Primitive = corePrimitive
		newRotationX, err := (&rx).Setup()
		if err != nil {
			return nil, err
		}
		return newRotationX, nil
	case "RotationY":
		var ry rotate.RotationY
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &ry)
		corePrimitive, err := decodeObject(ry.TypeName, ry.Data)
		if err != nil {
			return nil, err
		}
		ry.Primitive = corePrimitive
		newRotationY, err := (&ry).Setup()
		if err != nil {
			return nil, err
		}
		return newRotationY, nil
	case "RotationZ":
		var rz rotate.RotationZ
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &rz)
		corePrimitive, err := decodeObject(rz.TypeName, rz.Data)
		if err != nil {
			return nil, err
		}
		rz.Primitive = corePrimitive
		newRotationZ, err := (&rz).Setup()
		if err != nil {
			return nil, err
		}
		return newRotationZ, nil
	case "Quaternion":
		var q rotate.Quaternion
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(dataBytes, &q)
		corePrimitive, err := decodeObject(q.TypeName, q.Data)
		if err != nil {
			return nil, err
		}
		q.Primitive = corePrimitive
		newRotationZ, err := (&q).Setup()
		if err != nil {
			return nil, err
		}
		return newRotationZ, nil
	default:
		return nil, fmt.Errorf("type (%s) not a valid primitive type", typeName)
	}
}

func loadTextures(fileName string, tGamma float64) (map[string]texture.Texture, error) {
	texturesBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var texturesData []TextureData
	err = json.Unmarshal(texturesBytes, &texturesData)
	if err != nil {
		return nil, err
	}
	texturesMap := map[string]texture.Texture{}
	for _, t := range texturesData {
		if _, ok := texturesMap[t.Name]; ok {
			return nil, fmt.Errorf("texture (%s) redefined", t.Name)
		}
		switch t.TypeName {
		case "Color":
			var c texture.Color
			dataBytes, err := json.Marshal(t.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &c)
			texturesMap[t.Name] = &c
		case "Image":
			var i texture.Image
			dataBytes, err := json.Marshal(t.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &i)
			if i.Magnitude == 0.0 {
				i.Magnitude = 1.0
			}
			if i.Gamma == 0.0 {
				i.Gamma = tGamma
			}
			err = i.Load()
			if err != nil {
				return nil, err
			}
			texturesMap[t.Name] = &i
		default:
			return nil, fmt.Errorf("type (%s) not a valid texture type", t.TypeName)
		}
	}
	return texturesMap, nil
}

func loadMaterials(fileName, texturesFileName string, texturesMap map[string]texture.Texture) (map[string]material.Material, error) {
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
		if _, ok := materialsMap[m.Name]; ok {
			return nil, fmt.Errorf("material (%s) redefined", m.Name)
		}
		switch m.TypeName {
		case "Lambertian":
			var l material.Lambertian
			dataBytes, err := json.Marshal(m.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &l)
			var ok bool
			if m.ReflectanceTextureName == "" {
				l.ReflectanceTexture, ok = texturesMap["default"]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", "default", texturesFileName)
				}
			} else {
				l.ReflectanceTexture, ok = texturesMap[m.ReflectanceTextureName]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", m.ReflectanceTextureName, texturesFileName)
				}
			}

			if m.EmittanceTextureName == "" {
				l.EmittanceTexture, ok = texturesMap["default"]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", "default", texturesFileName)
				}
			} else {
				l.EmittanceTexture, ok = texturesMap[m.EmittanceTextureName]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", m.EmittanceTextureName, texturesFileName)
				}
			}
			materialsMap[m.Name] = &l
		case "Metal":
			var mtl material.Metal
			dataBytes, err := json.Marshal(m.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &mtl)
			var ok bool

			if m.ReflectanceTextureName == "" {
				mtl.ReflectanceTexture, ok = texturesMap["default"]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", "default", texturesFileName)
				}
			} else {
				mtl.ReflectanceTexture, ok = texturesMap[m.ReflectanceTextureName]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", m.ReflectanceTextureName, texturesFileName)
				}
			}

			if m.EmittanceTextureName == "" {
				mtl.EmittanceTexture, ok = texturesMap["default"]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", "default", texturesFileName)
				}
			} else {
				mtl.EmittanceTexture, ok = texturesMap[m.EmittanceTextureName]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", m.EmittanceTextureName, texturesFileName)
				}
			}
			materialsMap[m.Name] = &mtl
		case "Dielectric":
			var d material.Dielectric
			dataBytes, err := json.Marshal(m.Data)
			if err != nil {
				return nil, err
			}
			json.Unmarshal(dataBytes, &d)
			var ok bool
			if m.ReflectanceTextureName == "" {
				d.ReflectanceTexture, ok = texturesMap["default"]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", "default", texturesFileName)
				}
			} else {
				d.ReflectanceTexture, ok = texturesMap[m.ReflectanceTextureName]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", m.ReflectanceTextureName, texturesFileName)
				}
			}
			if m.EmittanceTextureName == "" {
				d.EmittanceTexture, ok = texturesMap["default"]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", "default", texturesFileName)
				}
			} else {
				d.EmittanceTexture, ok = texturesMap[m.EmittanceTextureName]
				if !ok {
					return nil, fmt.Errorf("selected Texture (%s) not in %s", m.EmittanceTextureName, texturesFileName)
				}
			}
			materialsMap[m.Name] = &d
		default:
			return nil, fmt.Errorf("type (%s) not a valid material type", m.TypeName)
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
	parameters.BackgroundColor = parameters.BackgroundColor.MultScalar(parameters.BGColorMagnitude)
	return &parameters, nil
}

func loadScene(fileName string) (*Scene, error) {
	sceneBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var scene Scene
	err = json.Unmarshal(sceneBytes, &scene)
	if err != nil {
		return nil, err
	}
	return &scene, nil
}
