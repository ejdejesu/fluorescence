package main

import (
	"encoding/json"
	"fluorescence/geometry/primitive"
	"fmt"
	"io/ioutil"
)

type Objects struct {
	Camera *Camera `json:"camera"`

	Total   []primitive.Primitive
	Spheres []*primitive.Sphere `json:"spheres"`
}

func LoadObjects(filename string) (*Objects, error) {
	objectsData, err := ioutil.ReadFile("./config/objects.json")
	if err != nil {
		fmt.Printf("Error reading objects file: %s\n", err.Error())
		return nil, err
	}
	var objects Objects
	err = json.Unmarshal(objectsData, &objects)
	if err != nil {
		fmt.Printf("Error unmarshalling objects data: %s\n", err.Error())
		return nil, err
	}

	err = objects.Camera.Setup()
	if err != nil {
		fmt.Printf("Error setting up camera: %s\n", err.Error())
		return nil, err
	}

	for _, s := range objects.Spheres {
		objects.Total = append(objects.Total, s)
	}

	return &objects, nil
}
