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
	param  rectParam
}

type rectParam struct {
}

func NewPatternRect(args []string) *PatternRect {
	return &PatternRect{args: args}
}

func (self *PatternRect) SetParam(paramSs []string) {
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
				err, rgb := NewRgbHexConver().Hex2rgb(argSs[1])
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
	self.setDefaultIfNull()
	ggContext := gg.NewContext(self.width, self.height)
	ggContext.DrawRectangle(0, 0, float64(self.width), float64(self.height))
	ggContext.SetRGB255(int(self.rgb.Red), int(self.rgb.Green), int(self.rgb.Blue))
	ggContext.Fill()
	return nil, ggContext
}

// 如果存在空值就设置默认值
func (self *PatternRect) setDefaultIfNull() {
	if self.width == 0 {
		self.width = 100
	}
	if self.height == 0 {
		self.height = 100
	}
}
