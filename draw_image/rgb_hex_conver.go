package draw_image

import (
	"errors"
	"fmt"
	"strconv"
)

type RgbHexConver struct {
	//rgb RGB
	//hex string
}

func NewRgbHexConver() *RgbHexConver {
	return &RgbHexConver{}
}

type RGB struct {
	Red, Green, Blue int64
}

//type HEX struct {
//	str string
//}

func (self *RgbHexConver) t2x(t int64) string {
	result := strconv.FormatInt(t, 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

/// rgb值转16进制字符值
func (self *RgbHexConver) rgb2hex(rgb RGB) (error, string) {
	if err := self.checkRgbSafe(rgb.Red); err != nil {
		return err, ""
	}
	if err := self.checkRgbSafe(rgb.Green); err != nil {
		return err, ""
	}
	if err := self.checkRgbSafe(rgb.Blue); err != nil {
		return err, ""
	}
	r := self.t2x(rgb.Red)
	g := self.t2x(rgb.Green)
	b := self.t2x(rgb.Blue)
	return nil, r + g + b
}

func (self *RgbHexConver) checkRgbSafe(i64Check int64) error {
	if i64Check > 0 && i64Check < 256 {
		return nil
	}
	return fmt.Errorf("[%d]rgb三个值的取值范围都只能0-255", i64Check)
	//return errors.New("[]rgb三个值的取值范围都只能0-255")
}

/// 16进制字符值转rgb值
func (self *RgbHexConver) hex2rgb(hex string) (error, RGB) {
	checkErr := self.checkHexSafe(hex)
	if checkErr != nil {
		return checkErr, RGB{}
	}
	r, _ := strconv.ParseInt(hex[:2], 16, 10)
	g, _ := strconv.ParseInt(hex[2:4], 16, 18)
	b, _ := strconv.ParseInt(hex[4:], 16, 10)
	return nil, RGB{r, g, b}
}

/// 检测16进制值是否安全正确
func (self *RgbHexConver) checkHexSafe(hex string) error {
	if hex == "" {
		return errors.New("16进制值不能为空")
	}
	if len(hex) != 6 {
		return errors.New("16进制值只能等于6位数字符")
	}
	return nil
}
