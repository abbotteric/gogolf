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
func pos_at_t(t float64, dv float64, iv float64, ic float64) float64 {
	return 0.5*dv*t*t + iv*t + ic
}

func real_space_to_img_space(p worldPoint, s worldSize, img image.RGBA) image.Point {
	return image.Point{int(math.Round(float64(img.Rect.Max.X) * (p.x / s.width))), img.Rect.Max.Y - int(math.Round(float64(img.Rect.Max.Y)*(p.y/s.height)))}
}

/*
	Set up some constants
*/
const ts = 0.1
const G = -9.8
const scale = 10.0

var world = worldSize{100, 59}
var iv = Iv{10, 30}
var ip = worldPoint{0, 0}

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

	pos := ip
	start := true
	for start || (t > 0 && pos.y >= 0) {
		pos.x = pos_at_t(t, 0, iv.x, ip.x)
		pos.y = pos_at_t(t, G, iv.y, ip.y)
		imgSpace := real_space_to_img_space(pos, world, *img)
		img.Set(imgSpace.X, imgSpace.Y, white)
		fmt.Printf("%f - %f,%f - %d,%d\n", t, pos.x, pos.y, imgSpace.X, imgSpace.Y)
		t += ts
		start = false
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)

}
