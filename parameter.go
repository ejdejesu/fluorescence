package main

import (
	"encoding/json"
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/shading"
	"fluorescence/shading/material"
	"fmt"
	"io/ioutil"
)

type Parameters struct {
	ImageWidth           int            `json:"image_width"`
	ImageHeight          int            `json:"image_height"`
	FileType             string         `json:"file_type"`
	FileDirectory        string         `json:"file_directory"`
	FileName             string         `json:"file_name"`
	GammaCorrection      int            `json:"gamma_correction"`
	AntialiasSampleCount int            `json:"antialias_sample_count"`
	MaxBounces           int            `json:"max_bounces"`
	BackgroundColor      *shading.Color `json:"background_color"`
	TMin                 float64        `json:"t_min"`
	TMax                 float64        `json:"t_max"`
	Scene                *Scene         `json:"scene"`
}

type Scene struct {
	Camera  *Camera  `json:"camera"`
	Objects *Objects `json:"objects"`
}

type Objects struct {
	Total     []primitive.Primitive
	Spheres   []*primitive.Sphere   `json:"spheres"`
	Triangles []*primitive.Triangle `json:"triangles"`
}

func LoadParameters(filename string) (*Parameters, error) {

	parametersData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var parameters Parameters
	err = json.Unmarshal(parametersData, &parameters)
	if err != nil {
		return nil, err
	}
	err = parameters.Scene.Camera.Setup(&parameters)
	if err != nil {
		fmt.Printf("Error setting up camera: %s\n", err.Error())
		return nil, err
	}

	for _, s := range parameters.Scene.Objects.Spheres {
		s.Material = &material.Lambertian{
			Reflectance_: &geometry.Vector{
				X: 0.5,
				Y: 0.5,
				Z: 0.5,
			},
			Emittance_: geometry.ZERO,
		}
		parameters.Scene.Objects.Total = append(parameters.Scene.Objects.Total, s)
	}

	for _, t := range parameters.Scene.Objects.Triangles {
		t.Material = &material.Lambertian{
			Reflectance_: &geometry.Vector{
				X: 0.3,
				Y: 0.3,
				Z: 0.9,
			},
			Emittance_: geometry.ZERO,
		}
		parameters.Scene.Objects.Total = append(parameters.Scene.Objects.Total, t)
	}

	return &parameters, nil
}
