package main

import (
	"log"

	"github.com/abelwei/mock_image/draw_image"
	"github.com/fogleman/gg"
)

func main() {
	_, dir := draw_image.NewLoadEnv().GetGenerateDir()
	im, err := gg.LoadImage("res/baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(512, 512)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	dc.Clip()
	dc.DrawImage(im, 0, 0)
	dc.SavePNG(dir+"/mask.png")
}
