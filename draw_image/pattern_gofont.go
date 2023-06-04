package draw_image

import (
	"bytes"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font/gofont/goregular"
	"strconv"
	"strings"
)

type PatternGofont struct {
	err    error
	args   []string
	width  int
	height int
	rgb    RGB
	param  gofontParam
}

type gofontParam struct {
	char, color string
	size        int
}

func NewPatternGofont(args []string) *PatternGofont {
	return &PatternGofont{args: args}
}

func (self *PatternGofont) SetParam(paramSs []string) {
	for _, param := range paramSs {
		prmSs := strings.Split(param, "=")
		if len(prmSs) > 1 {
			switch prmSs[0] {
			case "char":
				self.param.char = prmSs[1]
				break

			case "color":
				self.param.color = prmSs[1]
				break

			case "size":
				size, err := strconv.Atoi(prmSs[1])
				if err != nil {
					self.err = err
					return
				}
				self.param.size = size
				break
			default:

			}
		}
	}
}

func (self *PatternGofont) SaveFile(filePath string) error {
	if err, ggContext := self.parse().settingDraw(); err == nil {
		ggContext.SavePNG(filePath)
		return nil
	} else {
		return err
	}
}

func (self *PatternGofont) ResponseWriter() (error, []byte) {
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

func (self *PatternGofont) parse() *PatternGofont {
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

func (self *PatternGofont) settingDraw() (error, *gg.Context) {
	if self.err != nil {
		return self.err, nil
	}
	self.setDefaultIfNull()
	ggContext := gg.NewContext(self.width, self.height)
	ggContext.SetRGB255(int(self.rgb.Red), int(self.rgb.Green), int(self.rgb.Blue))
	ggContext.Clear()

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		logrus.Errorf("truetype.Parse Err:%s", err.Error())
		return err, ggContext
	}

	face := truetype.NewFace(font, &truetype.Options{Size: float64(self.param.size)})
	ggContext.SetFontFace(face)
	ggContext.SetHexColor(self.param.color)
	ggContext.DrawStringAnchored(self.param.char, float64(self.width)/2, float64(self.height)/2, 0.5, 0.5)

	return nil, ggContext
}

// 如果存在空值就设置默认值
func (self *PatternGofont) setDefaultIfNull() {
	if self.width == 0 {
		self.width = 100
	}
	if self.height == 0 {
		self.height = 100
	}
}
