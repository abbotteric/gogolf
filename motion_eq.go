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

type worldSize struct {
	width  float64
	height float64
}

/*
util functions
*/
func real_space_to_img_space(p worldPoint, s worldSize, img image.RGBA) image.Point {
	return image.Point{int(math.Round(float64(img.Rect.Max.X) * (p.x / s.width))), img.Rect.Max.Y - int(math.Round(float64(img.Rect.Max.Y)*(p.y/s.height)))}
}

func generate_image_file(world worldSize) *image.RGBA {
	width := int(math.Round(world.width * scale))
	height := int(math.Round(world.height * scale))

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	black := color.Black

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, black)
		}
	}

	return img
}

func rpm_2_v_ang(rpm float64) float64 {
	return (rpm * 2 * math.Pi) / 60
}

/*
	Set up some constants
*/
const dt = 0.1
const G = -9.8
const scale = 10.0

/*
main
*/
func main() {
	world := worldSize{200, 50}
	img := generate_image_file(world)
	backspin := 6000

	t := 0.0

	var nb Ball

	start := true
	for start || (t > 0 && nb.pos.y >= 0) {
		if start {
			nb = step(ball, dt, impact, rpm_2_v_ang(float64(backspin)))
		} else {
			nb = step(nb, dt, Vector{0, 0}, rpm_2_v_ang(float64(backspin)))
		}

		imgSpace := real_space_to_img_space(worldPoint(nb.pos), world, *img)
		img.Set(imgSpace.X, imgSpace.Y, color.White)

		fmt.Printf("%f || %f, %f || %f, %f || %d, %d\n", t, nb.pos.x, nb.pos.y, nb.vel.x, nb.vel.y, imgSpace.X, imgSpace.Y)

		t += dt
		start = false
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
