package main

import (
	"encoding/json"
	"fluorescence/shading"
	"io/ioutil"
)

type Parameters struct {
	ImageWidth           int            `json:"image_width"`
	ImageHeight          int            `json:"image_height"`
	FileType             string         `json:"file_type"`
	FileDirectory        string         `json:"file_directory"`
	FileName             string         `json:"file_name"`
	AntialiasSampleCount int            `json:"antialias_sample_count"`
	BackgroundColor      *shading.Color `json:"background_color"`
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
	return &parameters, nil
}
