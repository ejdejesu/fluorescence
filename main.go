package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {

	maxThreads := int64(runtime.NumCPU())
	fmt.Printf("Max Threads: %d\n", maxThreads)
	// maxThreads := int64(runtime.NumCPU() * 10)
	// maxThreads := int64(runtime.NumCPU() * 1000)
	// maxThreads := int64(1)
	// get parameters
	parametersFileName := "./config/parameters.json"
	camerasFileName := "./config/cameras.json"
	objectsFileName := "./config/objects.json"
	materialsFileName := "./config/materials.json"
	texturesFileName := "./config/textures.json"
	fmt.Printf("Loading Config files...\n")
	parameters, err := LoadConfigs(parametersFileName, camerasFileName, objectsFileName, materialsFileName, texturesFileName)
	if err != nil {
		fmt.Printf("Error loading parameters data: %s\n", err.Error())
		return
	}

	// create image
	fmt.Printf("Creating in-mem image...\n")
	img := image.NewRGBA64(image.Rect(0, 0, parameters.ImageWidth, parameters.ImageHeight))

	// fill image
	fmt.Printf("Filling in-mem image...\n")

	// spew.Dump(parameters.Scene.Objects)
	pixelCount := parameters.ImageWidth * parameters.ImageHeight
	doneChan := make(chan int, pixelCount)

	runtime.LockOSThread()

	startTime := time.Now()
	go TraceImage(parameters, img, doneChan, maxThreads)

	doneCount := 0
	printInterval := pixelCount / 1000
	for i := 0; i < pixelCount; i++ {
		<-doneChan
		doneCount++
		if pixelCount > 1000 && doneCount%printInterval == 0 {
			elapsedTime := time.Since(startTime)
			estimatedTime := time.Duration(float64(elapsedTime) * (float64(pixelCount) / float64(doneCount)))
			remainingTime := estimatedTime - elapsedTime
			fmt.Printf("\r\t%5.1f%% - Est. Rem: ~%v,\tTotal: ~%v", 100*float64(doneCount)/float64(pixelCount), remainingTime, estimatedTime)
		}
	}
	// wg.Wait()
	// sem.Release(0)
	totalDuration := time.Since(startTime)
	fmt.Printf("\tTotal time: %v\n", totalDuration)

	// create file
	fmt.Printf("Creating image file...\n")
	file, err := getImageFile(parameters)
	if err != nil {
		fmt.Printf("Error creating image file: %s\n", err.Error())
		return
	}
	defer file.Close()

	// encode image to file
	fmt.Printf("Writing in-mem image to image file...\n")
	err = png.Encode(file, img)
	if err != nil {
		fmt.Printf("Error encoding to image file: %s\n", err.Error())
		return
	}
	fmt.Printf("Done!\n")
	return
}

func getImageFile(parameters *Parameters) (*os.File, error) {
	filename := fmt.Sprintf(
		"%s%s_v%s_%ds_%s.%s",
		parameters.FileDirectory,
		strings.ReplaceAll(parameters.Scene.Name, " ", "_"),
		parameters.Version,
		parameters.SampleCount,
		time.Now().Format("2006-01-02_T150405"),
		parameters.FileType)
	os.MkdirAll(parameters.FileDirectory, os.ModePerm)
	return os.Create(filename)
}
