package draw_image

import (
	"errors"
	"strings"
)

type DrawPattern struct {
	opt PatternInter
}

func NewDrawPattern() *DrawPattern {
	return &DrawPattern{}
}

func (self *DrawPattern) Parse(dsl string) *DrawPattern {

	dslSs := strings.Split(dsl, ":")
	var (
		pattern            string
		patterArgSs, param []string
	)

	if len(dslSs) == 3 {
		pattern = dslSs[0]
		patterArgSs = strings.Split(dslSs[1], ",")
		param = strings.Split(dslSs[2], ",")
	}

	switch pattern {
	case "rect":
		self.opt = NewPatternRect(patterArgSs)
		self.opt.SetParam(param)
		break
	case "beziers":
		self.opt = NewPatternBeziers(patterArgSs)
		self.opt.SetParam(param)
		break
	case "circle":
		self.opt = NewPatternCircle(patterArgSs)
		self.opt.SetParam(param)
		break
	case "crisp":
		self.opt = NewPatternCrisp(patterArgSs)
		self.opt.SetParam(param)
		break
	case "cubic":
		self.opt = NewPatternCubic(patterArgSs)
		self.opt.SetParam(param)
		break
	case "ellipse":
		self.opt = NewPatternEllipse(patterArgSs)
		self.opt.SetParam(param)
		break
	case "gofont":
		self.opt = NewPatternGofont(patterArgSs)
		self.opt.SetParam(param)
		break
	default:

	}

	return self

}

func (self *DrawPattern) SaveDisk(savePath string) (err error) {
	if self.opt == nil {
		return errors.New("当前模式不存在")
	}
	return self.opt.SaveFile(savePath)
}

func (self *DrawPattern) Response() (error, []byte) {
	if self.opt == nil {
		return errors.New("当前模式不存在"), []byte{}
	}
	return self.opt.ResponseWriter()
}
