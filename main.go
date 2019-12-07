package main

import (
	"fluorescence/geometry"
	"fmt"
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	// get settings
	parameters, err := LoadParameters("./config/parameters.json")
	if err != nil {
		fmt.Printf("Error loading parameters data: %s\n", err.Error())
		return
	}
	// get objects
	objects, err := LoadObjects("./config/objects.json")
	if err != nil {
		fmt.Printf("Error loading settings data: %s\n", err.Error())
		return
	}

	// create image
	img := image.NewRGBA64(image.Rect(0, 0, parameters.ImageWidth, parameters.ImageHeight))

	startTime := time.Now()
	for y := 0; y < parameters.ImageHeight; y++ {
		for x := 0; x < parameters.ImageWidth; x++ {
			colorAccumulator := geometry.ZERO.Copy()
			for s := 0; s < parameters.AntialiasSampleCount; s++ {
				u := (float64(x) + rand.Float64()) / float64(parameters.ImageWidth)
				v := (float64(y) + rand.Float64()) / float64(parameters.ImageHeight)

				ray := objects.Camera.GetRay(u, v)

				tempColor := colorOf(parameters, objects, ray)
				colorAccumulator.AddInPlace(tempColor)
			}
			color := colorAccumulator.DivideFloat64(float64(parameters.AntialiasSampleCount)).ToColor()
			img.SetRGBA64(x, parameters.ImageHeight-y-1, *color.ToRGBA64())
		}
	}
	endTime := time.Now()
	fmt.Printf("Fill pixel time: %.3fms\n", 0.000001*float64(endTime.UnixNano()-startTime.UnixNano()))

	// create file

	file, err := getImageFile(parameters)
	if err != nil {
		fmt.Printf("Error creating image file: %s\n", err.Error())
		return
	}
	defer file.Close()

	// encode image to file
	err = png.Encode(file, img)
	if err != nil {
		fmt.Printf("Error encoding to image file: %s\n", err.Error())
		return
	}
}

func colorOf(parameters *Parameters, objects *Objects, r *geometry.Ray) *geometry.Vector {
	var minRayHit *geometry.RayHit
	minT := math.MaxFloat64
	hitSomething := false
	for _, g := range objects.Total {
		rayHit, wasHit := g.Intersection(r, 0, 100000.0)
		if wasHit && rayHit.T < minT {
			hitSomething = true
			minRayHit = rayHit
			minT = rayHit.T
		}
	}
	if hitSomething {
		return &geometry.Vector{
			X: (1 + minRayHit.NormalAtHit.X) / 2,
			Y: (1 + minRayHit.NormalAtHit.Y) / 2,
			Z: (1 + minRayHit.NormalAtHit.Z) / 2,
		}
	}
	return &geometry.Vector{
		X: parameters.BackgroundColor.Red,
		Y: parameters.BackgroundColor.Green,
		Z: parameters.BackgroundColor.Blue,
	}
}

func getImageFile(parameters *Parameters) (*os.File, error) {
	filename := fmt.Sprintf("%s%s_%s.%s", parameters.FileDirectory, parameters.FileName, time.Now().Format("2006-01-02_T150405"), parameters.FileType)
	os.MkdirAll(parameters.FileDirectory, os.ModePerm)
	return os.Create(filename)
}
