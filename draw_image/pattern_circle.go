package draw_image

import (
	"bytes"
	"github.com/fogleman/gg"
	"strconv"
	"strings"
)

type PatternCircle struct {
	err    error
	args   []string
	width  int
	height int
	rgb    RGB
	param  circleParam
}

type circleParam struct {
	centreX float64
	centreY float64
	radius  float64
}

func NewPatternCircle(args []string) *PatternCircle {
	return &PatternCircle{args: args}
}

func (self *PatternCircle) SetParam(paramSs []string) {
	for _, param := range paramSs {
		prmSs := strings.Split(param, "=")
		if len(prmSs) > 1 {
			switch prmSs[0] {
			case "x":
				centreX, err := strconv.ParseFloat(prmSs[1], 64)
				if err != nil {
					self.err = err
					return
				}
				self.param.centreX = centreX
				break
			case "y":
				centreY, err := strconv.ParseFloat(prmSs[1], 64)
				if err != nil {
					self.err = err
					return
				}
				self.param.centreY = centreY
				break
			case "radius":
				radius, err := strconv.ParseFloat(prmSs[1], 64)
				if err != nil {
					self.err = err
					return
				}
				self.param.radius = radius
				break
			default:

			}
		}
	}
}

func (self *PatternCircle) SaveFile(filePath string) error {
	if err, ggContext := self.parse().settingDraw(); err == nil {
		ggContext.SavePNG(filePath)
		return nil
	} else {
		return err
	}
}

func (self *PatternCircle) ResponseWriter() (error, []byte) {
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

func (self *PatternCircle) parse() *PatternCircle {
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

func (self *PatternCircle) settingDraw() (error, *gg.Context) {
	if self.err != nil {
		return self.err, nil
	}
	self.setDefaultIfNull()
	ggContext := gg.NewContext(self.width, self.height)
	ggContext.DrawCircle(self.param.centreX, self.param.centreY, self.param.radius)
	ggContext.SetRGB255(int(self.rgb.Red), int(self.rgb.Green), int(self.rgb.Blue))
	ggContext.Fill()
	return nil, ggContext
}

// 如果存在空值就设置默认值
func (self *PatternCircle) setDefaultIfNull() {
	if self.width == 0 {
		self.width = 100
	}
	if self.height == 0 {
		self.height = 100
	}
}
