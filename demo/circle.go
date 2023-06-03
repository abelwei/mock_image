package main

import (
	"github.com/abelwei/mock_image/draw_image"
	"github.com/fogleman/gg"
)

func main() {
	_, dir := draw_image.NewLoadEnv().GetGenerateDir()
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	// dc.DrawRectangle(500, 500, 200, 200)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG(dir+"/circle.png")
}
