package main

import (
	"context"
	"fluorescence/geometry"
	"image"
	"math"
	"math/rand"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

type Tile struct {
	Origin geometry.Point
	Span   geometry.Vector
}

func Trace(params *Parameters, img *image.RGBA64, doneChan chan<- int, maxThreads int64) {

	tiles := getTiles(params, img)

	sem := semaphore.NewWeighted(maxThreads)
	runtime.LockOSThread()

	for _, tile := range tiles {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		sem.Acquire(context.Background(), 1)
		go traceTile(params, r, img, doneChan, sem, tile, params.SampleCount)
	}
}

func traceTile(p *Parameters, rng *rand.Rand, img *image.RGBA64, dc chan<- int, sem *semaphore.Weighted, t Tile, sampleCount int) {
	defer sem.Release(1)
	for y := t.Origin.Y; y < t.Origin.Y+t.Span.Y; y++ {
		for x := t.Origin.X; x < t.Origin.X+t.Span.X; x++ {
			colorAccumulator := geometry.VECTOR_ZERO
			for s := 0; s < p.SampleCount; s++ {
				u := (float64(x) + rng.Float64()) / float64(p.ImageWidth)
				v := (float64(y) + rng.Float64()) / float64(p.ImageHeight)

				ray := p.Scene.Camera.GetRay(u, v, rng)

				tempColor := colorOf(p, ray, rng, 0)
				colorAccumulator = colorAccumulator.Add(tempColor)
			}
			colorAccumulator = colorAccumulator.DivScalar(float64(p.SampleCount)).Clamp(0, 1).Pow(1.0 / float64(p.GammaCorrection))
			color := colorAccumulator.ToColor()

			img.SetRGBA64(int(x), p.ImageHeight-int(y)-1, color.ToRGBA64())
			dc <- 1
		}
	}
}

func colorOf(parameters *Parameters, r geometry.Ray, rng *rand.Rand, depth int) geometry.Vector {

	backgroundColor := geometry.VectorFromColor(parameters.BackgroundColor)

	if depth > parameters.MaxBounces {
		return geometry.VECTOR_ZERO
	}
	rayHit, hitSomething := parameters.Scene.Objects.Intersection(r, parameters.TMin, parameters.TMax)
	if !hitSomething {
		return backgroundColor
	}

	mat := rayHit.Material

	if mat.Reflectance() == geometry.VECTOR_ZERO {
		return mat.Emittance()
	}

	scatteredRay, wasScattered := rayHit.Material.Scatter(*rayHit, rng)
	if !wasScattered {
		return geometry.VECTOR_ZERO
	}
	incomingColor := colorOf(parameters, scatteredRay, rng, depth+1)
	return mat.Emittance().Add(mat.Reflectance().MultVector(incomingColor))
}

func getTiles(p *Parameters, i *image.RGBA64) []Tile {
	tiles := []Tile{}
	for y := 0; y < p.ImageHeight; y += p.TileHeight {
		for x := 0; x < p.ImageWidth; x += p.TileWidth {
			width := math.Min(float64(p.TileWidth), float64(p.ImageWidth-x))
			height := math.Min(float64(p.TileHeight), float64(p.ImageHeight-y))
			tiles = append(tiles, Tile{
				Origin: geometry.Point{
					X: float64(x),
					Y: float64(y),
				},
				Span: geometry.Vector{
					X: width,
					Y: height,
				},
			})
		}
	}
	return tiles
}
