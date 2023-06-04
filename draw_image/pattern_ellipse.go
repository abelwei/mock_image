package draw_image

import (
	"bytes"
	"github.com/fogleman/gg"
	"strconv"
	"strings"
)

type PatternEllipse struct {
	err    error
	args   []string
	width  int
	height int
	rgb    RGB
	param  ellipseParam
}

type ellipseParam struct {
}

func NewPatternEllipse(args []string) *PatternEllipse {
	return &PatternEllipse{args: args}
}

func (self *PatternEllipse) SetParam(paramSs []string) {
	for _, param := range paramSs {
		prmSs := strings.Split(param, "=")
		if len(prmSs) > 1 {
			switch prmSs[0] {

			default:

			}
		}
	}
}

func (self *PatternEllipse) SaveFile(filePath string) error {
	if err, ggContext := self.parse().settingDraw(); err == nil {
		ggContext.SavePNG(filePath)
		return nil
	} else {
		return err
	}
}

func (self *PatternEllipse) ResponseWriter() (error, []byte) {
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

func (self *PatternEllipse) parse() *PatternEllipse {
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

func (self *PatternEllipse) settingDraw() (error, *gg.Context) {
	if self.err != nil {
		return self.err, nil
	}
	self.setDefaultIfNull()
	ggContext := gg.NewContext(self.width, self.height)
	ggContext.SetRGB255(int(self.rgb.Red), int(self.rgb.Green), int(self.rgb.Blue))
	ggContext.Clear()

	ggContext.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		ggContext.Push()
		ggContext.RotateAbout(gg.Radians(float64(i)),
			float64(self.width)/2, float64(self.height)/2)
		ggContext.DrawEllipse(float64(self.width)/2, float64(self.height)/2,
			float64(self.width)*7/16, float64(self.height)/8)
		ggContext.Fill()
		ggContext.Pop()
	}

	return nil, ggContext
}

// 如果存在空值就设置默认值
func (self *PatternEllipse) setDefaultIfNull() {
	if self.width == 0 {
		self.width = 100
	}
	if self.height == 0 {
		self.height = 100
	}
}
