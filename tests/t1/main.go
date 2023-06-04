package main

import (
	"github.com/abelwei/mock_image/draw_image"
	"github.com/sirupsen/logrus"
)

func main() {
	_, dir := draw_image.NewLoadEnv().GetGenerateDir()
	drawFormDsl := `ellipse:w=500,h=500,color=FF6EB4:`
	//drawFormDsl := `beziers:w=500,h=200,color=000000:square=100,row=5,column=1`
	filePath := dir + "/0t1.png"
	err := draw_image.NewDrawPattern().Parse(drawFormDsl).SaveDisk(filePath)
	if err != nil {
		logrus.Errorf("dsl:%s, path:%s", drawFormDsl, filePath)
		return
	}
	logrus.Infof("ok: %s", filePath)
}
