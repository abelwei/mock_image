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

	dslSs := strings.Split(dsl, ",")
	pattern := dslSs[0]
	var args []string
	args = append(args, dslSs[1:]...)
	switch pattern {
	case "rect":
		self.opt = NewPatternRect(args)
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
