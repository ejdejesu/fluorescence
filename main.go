package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"
)

type Settings struct {
	ImageWidth    int    `json:"image_width"`
	ImageHeight   int    `json:"image_height"`
	FileType      string `json:"file_type"`
	FileDirectory string `json:"file_directory"`
	FileName      string `json:"file_name"`
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
	// create image
	image := image.NewRGBA64(image.Rect(0, 0, settings.ImageWidth, settings.ImageHeight))

	startTime := time.Now()
	for x := image.Bounds().Min.X; x < image.Bounds().Max.X; x++ {
		for y := image.Bounds().Min.Y; y < image.Bounds().Max.Y; y++ {
			maxBoundsX := image.Bounds().Max.X
			maxBoundsY := image.Bounds().Max.Y

			redVal := float64(x) / float64(maxBoundsX)
			scaledRedVal := uint16(redVal * float64(math.MaxUint16))
			greenVal := float64(y) / float64(maxBoundsY)
			scaledGreenVal := uint16(greenVal * float64(math.MaxUint16))
			blueVal := rand.Float64()
			scaledBlueVal := uint16(blueVal * float64(math.MaxUint16))
			fmt.Println(scaledBlueVal)
			// fmt.Printf("%d, %d, %d\n", scaledRedVal, scaledGreenVal, scaledBlueVal)
			// yVal := uint16((float64(y) / float64(maxBoundsY)) * math.MaxInt16)

			image.SetRGBA64(x, y, color.RGBA64{scaledRedVal, scaledGreenVal, scaledBlueVal, math.MaxUint16})
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

func getImageFile(settings Settings) (*os.File, error) {
	filename := fmt.Sprintf("%s%s_%s.%s", settings.FileDirectory, settings.FileName, time.Now().Format("2006-01-02_T150405"), settings.FileType)
	os.MkdirAll(settings.FileDirectory, os.ModePerm)
	return os.Create(filename)
}
