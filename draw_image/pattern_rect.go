package draw_image

import (
	"bytes"
	"github.com/fogleman/gg"
	"strconv"
	"strings"
)

type PatternRect struct {
	err    error
	args   []string
	width  int
	height int
	rgb    RGB
}

func NewPatternRect(args []string) *PatternRect {
	return &PatternRect{args: args}
}

func (self *PatternRect) parse() *PatternRect {
	for _, arg := range self.args {
		argSs := strings.Split(arg, "=")
		if len(argSs) > 1 {
			switch argSs[0] {
			case "w":
				width, err := strconv.Atoi(argSs[1])
				if err != nil {
					self.err = err
					return self
				}
				self.width = width
				break
			case "h":
				height, err := strconv.Atoi(argSs[1])
				if err != nil {
					self.err = err
					return self
				}
				self.height = height
				break
			case "color":
				err, rgb := NewRgbHexConver().hex2rgb(argSs[1])
				if err != nil {
					self.err = err
					return self
				}
				self.rgb = rgb
				break
			default:

			}
		}
	}
	return self
}

func (self *PatternRect) SaveFile(filePath string) error {
	//return self.parse().settingDraw().SavePNG(filePath)
	if err, ggContext := self.parse().settingDraw(); err == nil {
		ggContext.SavePNG(filePath)
		return nil
	} else {
		return err
	}
}

func (self *PatternRect) ResponseWriter() (error, []byte) {
	if err, ggContext := self.parse().settingDraw(); err == nil {
		buffer := bytes.NewBuffer(nil)
		enErr := ggContext.EncodePNG(buffer)
		if enErr != nil {
			return enErr, []byte{}
		}
		bt := buffer.Bytes()
		return nil, bt
	} else {
		return err, []byte{}
	}
}

func (self *PatternRect) settingDraw() (error, *gg.Context) {
	if self.err != nil {
		return self.err, nil
	}
	ggContext := gg.NewContext(self.width, self.height)
	ggContext.DrawRectangle(0, 0, float64(self.width), float64(self.height))
	ggContext.SetRGB(float64(self.rgb.Red), float64(self.rgb.Green), float64(self.rgb.Blue))
	ggContext.Fill()
	return nil, ggContext
}
