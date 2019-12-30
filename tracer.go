package main

import (
	"context"
	"fluorescence/geometry"
	"image"
	"math"
	"math/rand"
	"runtime"
	"time"

	"github.com/go-gl/mathgl/mgl64"
	"golang.org/x/sync/semaphore"
)

// Tile holds information about a section of pixels on the image
type Tile struct {
	Origin mgl64.Vec2 // Top left corner of Tile
	Span   mgl64.Vec2 // Width and Height of Tile
}

// TraceImage is the powerhouse function, driving the raycasting algorith by casting rays into the scene
func TraceImage(params *Parameters, img *image.RGBA64, doneChan chan<- int, maxThreads int64) {

	tiles := getTiles(params, img)

	rand.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	sem := semaphore.NewWeighted(maxThreads)
	runtime.LockOSThread()

	for _, tile := range tiles {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		sem.Acquire(context.Background(), 1)
		go traceTile(params, r, img, doneChan, sem, tile, params.SampleCount)
	}

}

// traceTile iterates over the pixels in a tile and writes the received colors to the image
func traceTile(p *Parameters, rng *rand.Rand, img *image.RGBA64, dc chan<- int, sem *semaphore.Weighted, t Tile, sampleCount int) {
	defer sem.Release(1)
	for y := t.Origin.Y(); y < t.Origin.Y()+t.Span.Y(); y++ {
		for x := t.Origin.X(); x < t.Origin.X()+t.Span.X(); x++ {
			pixelColor := tracePixel(p, int(x), int(y), rng)

			img.SetRGBA64(int(x), p.ImageHeight-int(y)-1, geometry.Vec3ToRGBA64(pixelColor))
			dc <- 1
		}
	}
	// dc <- 1
}

// tracePixel gets the color for a pixel
func tracePixel(p *Parameters, x, y int, rng *rand.Rand) mgl64.Vec3 {
	pixelColor := mgl64.Vec3{}
	for s := 0; s < p.SampleCount; s++ {
		// pick a random spot on the pixel to shoot a ray into
		// this is purely random, NOT stratified
		u := (float64(x) + rng.Float64()) / float64(p.ImageWidth)
		v := (float64(y) + rng.Float64()) / float64(p.ImageHeight)

		ray := p.Scene.Camera.GetRay(u, v, rng)

		tempColor := traceRay(p, ray, rng, 0)
		pixelColor = pixelColor.Add(tempColor)
	}
	if p.UseScalingTruncation {
		return geometry.PowVec3(geometry.ScaleDownVec3(pixelColor.Mul(1.0/float64(p.SampleCount)), 1.0), 1.0/p.GammaCorrection)
	}
	return geometry.PowVec3(geometry.ClampVec3(pixelColor.Mul(1.0/float64(p.SampleCount)), 0.0, 1.0), 1.0/p.GammaCorrection)

}

// traceRay casts in individual ray into the scene
func traceRay(parameters *Parameters, r geometry.Ray, rng *rand.Rand, depth int) mgl64.Vec3 {

	// if we've gone too deep...
	if depth > parameters.MaxBounces {
		// ...just return BLACK
		return geometry.Vec3Zero
	}
	// check if we've hit something
	rayHit, hitSomething := parameters.Scene.Objects.Intersection(r, parameters.TMin, parameters.TMax)
	// if we did not hit something...
	if !hitSomething {
		// ...return the background color
		// TODO: add support for HDR skymaps
		return parameters.BackgroundColor
	}

	mat := rayHit.Material

	// if the surface is BLACK, it's not going to let any incoming light contribute to the outgoing color
	// so we can safely say no light is reflected and simply return the emittance of the material
	if mat.Reflectance(rayHit.U, rayHit.V) == geometry.Vec3Zero {
		return mat.Emittance(rayHit.U, rayHit.V)
	}

	// get the reflection incoming ray
	scatteredRay, wasScattered := rayHit.Material.Scatter(*rayHit, rng)
	// if no ray could have reflected to us, we just return BLACK
	if !wasScattered {
		return geometry.Vec3Zero
	}
	// get the color that came to this point and gave us the outgoing ray
	incomingColor := traceRay(parameters, scatteredRay, rng, depth+1)
	// return the (very-roughly approximated) value of the rendering equation
	return mat.Emittance(rayHit.U, rayHit.V).Add(geometry.MulVec3(mat.Reflectance(rayHit.U, rayHit.V), incomingColor))
}

// getTiles creates and return a grid of tiles on the image
func getTiles(p *Parameters, i *image.RGBA64) []Tile {
	tiles := []Tile{}
	for y := 0; y < p.ImageHeight; y += p.TileHeight {
		for x := 0; x < p.ImageWidth; x += p.TileWidth {
			width := math.Min(float64(p.TileWidth), float64(p.ImageWidth-x))
			height := math.Min(float64(p.TileHeight), float64(p.ImageHeight-y))
			tiles = append(tiles, Tile{
				Origin: mgl64.Vec2{
					float64(x),
					float64(y),
				},
				Span: mgl64.Vec2{
					width,
					height,
				},
			})
		}
	}
	return tiles
}
