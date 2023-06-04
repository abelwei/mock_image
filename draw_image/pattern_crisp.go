package draw_image

import (
	"bytes"
	"github.com/fogleman/gg"
	"strconv"
	"strings"
)

type PatternCrisp struct {
	err    error
	args   []string
	width  int
	height int
	rgb    RGB
	param  crispParam
}

type crispParam struct {
	minor, major int
}

func NewPatternCrisp(args []string) *PatternCrisp {
	return &PatternCrisp{args: args}
}

func (self *PatternCrisp) SetParam(paramSs []string) {
	for _, param := range paramSs {
		prmSs := strings.Split(param, "=")
		if len(prmSs) > 1 {
			switch prmSs[0] {
			case "minor":
				minor, err := strconv.Atoi(prmSs[1])
				if err != nil {
					self.err = err
					return
				}
				self.param.minor = minor
				break

			case "major":
				major, err := strconv.Atoi(prmSs[1])
				if err != nil {
					self.err = err
					return
				}
				self.param.major = major
				break

			default:

			}
		}
	}
}

func (self *PatternCrisp) SaveFile(filePath string) error {
	if err, ggContext := self.parse().settingDraw(); err == nil {
		ggContext.SavePNG(filePath)
		return nil
	} else {
		return err
	}
}

func (self *PatternCrisp) ResponseWriter() (error, []byte) {
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

func (self *PatternCrisp) parse() *PatternCrisp {
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

func (self *PatternCrisp) settingDraw() (error, *gg.Context) {
	if self.err != nil {
		return self.err, nil
	}
	self.setDefaultIfNull()
	ggContext := gg.NewContext(self.width, self.height)
	ggContext.SetRGB255(int(self.rgb.Red), int(self.rgb.Green), int(self.rgb.Blue))
	ggContext.Clear()

	// minor grid
	for x := self.param.minor; x < self.width; x += self.param.minor {
		fx := float64(x) + 0.5
		ggContext.DrawLine(fx, 0, fx, float64(self.height))
	}
	for y := self.param.minor; y < self.height; y += self.param.minor {
		fy := float64(y) + 0.5
		ggContext.DrawLine(0, fy, float64(self.width), fy)
	}
	ggContext.SetLineWidth(1)
	ggContext.SetRGBA(0, 0, 0, 0.25)
	ggContext.Stroke()

	// major grid
	for x := self.param.major; x < self.width; x += self.param.major {
		fx := float64(x) + 0.5
		ggContext.DrawLine(fx, 0, fx, float64(self.height))
	}
	for y := self.param.major; y < self.height; y += self.param.major {
		fy := float64(y) + 0.5
		ggContext.DrawLine(0, fy, float64(self.width), fy)
	}
	ggContext.SetLineWidth(2)
	ggContext.SetRGBA(0, 0, 0, 0.5)
	ggContext.Stroke()
	return nil, ggContext
}

// 如果存在空值就设置默认值
func (self *PatternCrisp) setDefaultIfNull() {
	if self.width == 0 {
		self.width = 100
	}
	if self.height == 0 {
		self.height = 100
	}
}
