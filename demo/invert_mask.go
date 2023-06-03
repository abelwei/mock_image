package main

import (
	"github.com/abelwei/mock_image/draw_image"
	"github.com/fogleman/gg"
)

func main() {
	_, dir := draw_image.NewLoadEnv().GetGenerateDir()
	dc := gg.NewContext(1024, 1024)
	dc.DrawCircle(512, 512, 384)
	dc.Clip()
	dc.InvertMask()
	dc.DrawRectangle(0, 0, 1024, 1024)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG(dir+"/invert_mask.png")
}
