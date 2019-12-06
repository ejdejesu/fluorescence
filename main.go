package main

import (
	"encoding/json"
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/shading"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"time"
)

type Settings struct {
	ImageWidth      int            `json:"image_width"`
	ImageHeight     int            `json:"image_height"`
	FileType        string         `json:"file_type"`
	FileDirectory   string         `json:"file_directory"`
	FileName        string         `json:"file_name"`
	BackgroundColor *shading.Color `json:"background_color"`
}

type Objects struct {
	Total   []primitive.Geometry
	Spheres []*primitive.Sphere `json:"spheres"`
}

func main() {
	// get settings
	settingsData, err := ioutil.ReadFile("./config/settings.json")
	if err != nil {
		fmt.Printf("Error reading settings file: %s\n", err.Error())
		return
	}
	var settings Settings
	err = json.Unmarshal(settingsData, &settings)
	if err != nil {
		fmt.Printf("Error unmarshalling settings data: %s\n", err.Error())
		return
	}

	// get objects
	objectsData, err := ioutil.ReadFile("./config/objects.json")
	if err != nil {
		fmt.Printf("Error reading objects file: %s\n", err.Error())
		return
	}
	var objects Objects
	err = json.Unmarshal(objectsData, &objects)
	if err != nil {
		fmt.Printf("Error unmarshalling objects data: %s\n", err.Error())
		return
	}

	for _, s := range objects.Spheres {
		objects.Total = append(objects.Total, s)
	}

	// create image
	image := image.NewRGBA64(image.Rect(0, 0, settings.ImageWidth, settings.ImageHeight))

	origin := &geometry.Point{0.0, 0.0, 0.0}
	topLeft := &geometry.Point{-2.0, 1.0, -1.0}
	horizontal := &geometry.Vector{4.0, 0.0, 0.0}
	vertical := &geometry.Vector{0.0, -2.0, 0.0}
	startTime := time.Now()
	for y := 0; y < settings.ImageHeight; y++ {
		for x := 0; x < settings.ImageWidth; x++ {

			u := float64(x) / float64(settings.ImageWidth)
			v := float64(y) / float64(settings.ImageHeight)

			ray := &geometry.Ray{
				Origin:    origin,
				Direction: origin.To(topLeft.AddVector(horizontal.MultiplyFloat64(u)).AddVector(vertical.MultiplyFloat64(v))).Unit(),
			}

			color := color(settings, objects, ray)

			image.SetRGBA64(x, y, *color.ToRGBA64())
		}
	}
	endTime := time.Now()
	fmt.Printf("Fill pixel time: %.3fms\n", 0.000001*float64(endTime.UnixNano()-startTime.UnixNano()))

	// create file

	file, err := getImageFile(settings)
	if err != nil {
		fmt.Printf("Error creating image file: %s\n", err.Error())
		return
	}
	defer file.Close()

	// encode image to file
	err = png.Encode(file, image)
	if err != nil {
		fmt.Printf("Error encoding to image file: %s\n", err.Error())
		return
	}
}

func color(settings Settings, objects Objects, r *geometry.Ray) *shading.Color {
	for _, g := range objects.Total {
		_, normal, _, _, hit := g.Intersection(r, 0, 100000.0)
		if hit {
			return &shading.Color{(1 + normal.X) / 2, (1 + normal.Y) / 2, (1 + normal.Z) / 2, 1.0}
		}
	}
	return settings.BackgroundColor
}

func getImageFile(settings Settings) (*os.File, error) {
	filename := fmt.Sprintf("%s%s_%s.%s", settings.FileDirectory, settings.FileName, time.Now().Format("2006-01-02_T150405"), settings.FileType)
	os.MkdirAll(settings.FileDirectory, os.ModePerm)
	return os.Create(filename)
}
