package main

import (
	"github.com/abelwei/mock_image/draw_image"
	"github.com/sirupsen/logrus"
)

func main() {
	_, dir := draw_image.NewLoadEnv().GetGenerateDir()
	drawFormDsl := `rect,w=500,h=200,color=ff0000`
	filePath := dir + "/0t1.png"
	err := draw_image.NewDrawPattern().Parse(drawFormDsl).SaveDisk(filePath)
	if err != nil {
		logrus.Errorf("dsl:%s, path:%s", drawFormDsl, filePath)
		return
	}
	logrus.Infof("ok: %s", filePath)
}
