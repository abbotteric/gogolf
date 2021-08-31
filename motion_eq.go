package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

/*
types
*/
type worldPoint struct {
	x float64
	y float64
}

type Iv struct {
	x float64
	y float64
}

type worldSize struct {
	width  float64
	height float64
}

/*
util functions
*/
// func pos_at_t(t float64, dv float64, iv float64, ic float64) float64 {
// 	return 0.5*dv*t*t + iv*t + ic
// }

func real_space_to_img_space(p worldPoint, s worldSize, img image.RGBA) image.Point {
	return image.Point{int(math.Round(float64(img.Rect.Max.X) * (p.x / s.width))), img.Rect.Max.Y - int(math.Round(float64(img.Rect.Max.Y)*(p.y/s.height)))}
}

/*
	Set up some constants
*/
const dt = 0.1
const G = -9.8
const scale = 10.0

var world = worldSize{200, 50}

/*
main
*/
func main() {
	width := int(math.Round(world.width * scale))
	height := int(math.Round(world.height * scale))

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	black := color.Black
	white := color.White

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, black)
		}
	}

	t := 0.0

	var nb Ball

	start := true
	for start || (t > 0 && nb.pos.y >= 0) {
		if start {
			nb = step(ball, dt, impact)
		} else {
			nb = step(nb, dt, Vector{0, 0})
		}

		imgSpace := real_space_to_img_space(worldPoint(nb.pos), world, *img)
		img.Set(imgSpace.X, imgSpace.Y, white)

		fmt.Printf("%f || %f, %f || %f, %f || %d, %d\n", t, nb.pos.x, nb.pos.y, nb.vel.x, nb.vel.y, imgSpace.X, imgSpace.Y)

		t += dt
		start = false
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
