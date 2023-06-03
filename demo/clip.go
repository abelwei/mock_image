package main

import (
	"github.com/abelwei/mock_image/draw_image"
	"github.com/fogleman/gg"
)

func main() {
	_, dir := draw_image.NewLoadEnv().GetGenerateDir()
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(350, 500, 300)
	dc.Clip()
	dc.DrawCircle(650, 500, 300)
	dc.Clip()
	dc.DrawRectangle(0, 0, 1000, 1000)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG(dir+"/clip.png")
}
