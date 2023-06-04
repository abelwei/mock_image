package draw_image

import (
	"bytes"
	"github.com/fogleman/gg"
	"math/rand"
	"strconv"
	"strings"
)

type PatternBeziers struct {
	err    error
	args   []string
	width  int
	height int
	rgb    RGB
	param  beziersParam
}

type beziersParam struct {
	square float64
	row    int
	column int
}

func NewPatternBeziers(args []string) *PatternBeziers {
	return &PatternBeziers{args: args}
}

func (self *PatternBeziers) SetParam(paramSs []string) {
	for _, param := range paramSs {
		prmSs := strings.Split(param, "=")
		if len(prmSs) > 1 {
			switch prmSs[0] {
			case "square":
				square, err := strconv.ParseFloat(prmSs[1], 64)
				if err != nil {
					self.err = err
					return
				}
				self.param.square = square
				break
			case "row":
				row, err := strconv.Atoi(prmSs[1])
				if err != nil {
					self.err = err
					return
				}
				self.param.row = row
				break
			case "column":
				column, err := strconv.Atoi(prmSs[1])
				if err != nil {
					self.err = err
					return
				}
				self.param.column = column
				break
			default:

			}
		}
	}
}

func (self *PatternBeziers) SaveFile(filePath string) error {
	if err, ggContext := self.parse().settingDraw(); err == nil {
		ggContext.SavePNG(filePath)
		return nil
	} else {
		return err
	}
}

func (self *PatternBeziers) ResponseWriter() (error, []byte) {
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

func (self *PatternBeziers) parse() *PatternBeziers {
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

func (self *PatternBeziers) settingDraw() (error, *gg.Context) {
	if self.err != nil {
		return self.err, nil
	}
	self.setDefaultIfNull()
	ggContext := gg.NewContext(self.width, self.height)
	//ggContext.DrawRectangle(0, 0, float64(self.width), float64(self.height))
	//ggContext.SetRGB(float64(self.rgb.Red), float64(self.rgb.Green), float64(self.rgb.Blue))
	ggContext.SetRGB(1, 1, 1)
	//ggContext.Fill()
	H := self.param.column
	S := self.param.square
	W := self.param.row
	ggContext.Clear()
	for j := 0; j < H; j++ {
		for i := 0; i < W; i++ {
			x := float64(i)*S + S/2
			y := float64(j)*S + S/2
			ggContext.Push()
			ggContext.Translate(x, y)
			ggContext.Scale(S/2, S/2)
			if j%2 == 0 {
				self.randomCubic(ggContext)
			} else {
				self.randomQuadratic(ggContext)
			}
			ggContext.Pop()
		}
	}
	return nil, ggContext
}

// 如果存在空值就设置默认值
func (self *PatternBeziers) setDefaultIfNull() {
	if self.width == 0 {
		self.width = 100
	}
	if self.height == 0 {
		self.height = 100
	}
}

func (self *PatternBeziers) random() float64 {
	return rand.Float64()*2 - 1
}

func (self *PatternBeziers) point() (x, y float64) {
	return self.random(), self.random()
}

func (self *PatternBeziers) drawCurve(dc *gg.Context) {
	dc.SetRGBA(0, 0, 0, 0.1)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(12)
	dc.Stroke()
}

func (self *PatternBeziers) drawPoints(dc *gg.Context) {
	dc.SetRGBA(1, 0, 0, 0.5)
	dc.SetLineWidth(2)
	dc.Stroke()
}

func (self *PatternBeziers) randomQuadratic(dc *gg.Context) {
	x0, y0 := self.point()
	x1, y1 := self.point()
	x2, y2 := self.point()
	dc.MoveTo(x0, y0)
	dc.QuadraticTo(x1, y1, x2, y2)
	self.drawCurve(dc)
	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	self.drawPoints(dc)
}

func (self *PatternBeziers) randomCubic(dc *gg.Context) {
	x0, y0 := self.point()
	x1, y1 := self.point()
	x2, y2 := self.point()
	x3, y3 := self.point()
	dc.MoveTo(x0, y0)
	dc.CubicTo(x1, y1, x2, y2, x3, y3)
	self.drawCurve(dc)
	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	dc.LineTo(x3, y3)
	self.drawPoints(dc)
}
