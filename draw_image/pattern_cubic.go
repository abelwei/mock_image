package draw_image

import (
	"bytes"
	"github.com/fogleman/gg"
	"strconv"
	"strings"
)

type PatternCubic struct {
	err    error
	args   []string
	width  int
	height int
	rgb    RGB
	param  cubicParam
}

type cubicParam struct {
}

func NewPatternCubic(args []string) *PatternCubic {
	return &PatternCubic{args: args}
}

func (self *PatternCubic) SetParam(paramSs []string) {
	for _, param := range paramSs {
		prmSs := strings.Split(param, "=")
		if len(prmSs) > 1 {
			switch prmSs[0] {

			default:

			}
		}
	}
}

func (self *PatternCubic) SaveFile(filePath string) error {
	if err, ggContext := self.parse().settingDraw(); err == nil {
		ggContext.SavePNG(filePath)
		return nil
	} else {
		return err
	}
}

func (self *PatternCubic) ResponseWriter() (error, []byte) {
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

func (self *PatternCubic) parse() *PatternCubic {
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

func (self *PatternCubic) settingDraw() (error, *gg.Context) {
	if self.err != nil {
		return self.err, nil
	}
	self.setDefaultIfNull()
	ggContext := gg.NewContext(self.width, self.height)
	ggContext.SetRGB255(int(self.rgb.Red), int(self.rgb.Green), int(self.rgb.Blue))
	ggContext.Clear()

	ggContext.Translate(float64(self.width)/2, float64(self.height)/2)
	//ggContext.Scale(20, 20)
	ggContext.Scale(float64(self.width)/25, float64(self.height)/25)

	var x0, y0, x1, y1, x2, y2, x3, y3 float64
	x0, y0 = -10, 0
	x1, y1 = -8, -8
	x2, y2 = 8, 8
	x3, y3 = 10, 0

	ggContext.MoveTo(x0, y0)
	ggContext.CubicTo(x1, y1, x2, y2, x3, y3)
	ggContext.SetRGBA(0, 0, 0, 0.2)
	ggContext.SetLineWidth(8)
	ggContext.FillPreserve()
	ggContext.SetRGB(0, 0, 0)
	ggContext.SetDash(16, 24)
	ggContext.Stroke()

	ggContext.MoveTo(x0, y0)
	ggContext.LineTo(x1, y1)
	ggContext.LineTo(x2, y2)
	ggContext.LineTo(x3, y3)
	ggContext.SetRGBA(1, 0, 0, 0.4)
	ggContext.SetLineWidth(2)
	ggContext.SetDash(4, 8, 1, 8)
	ggContext.Stroke()

	return nil, ggContext
}

// 如果存在空值就设置默认值
func (self *PatternCubic) setDefaultIfNull() {
	if self.width == 0 {
		self.width = 100
	}
	if self.height == 0 {
		self.height = 100
	}
}
