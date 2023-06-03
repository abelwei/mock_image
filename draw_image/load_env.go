package draw_image

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type LoadEnv struct{
	Err error
}

func NewLoadEnv() *LoadEnv {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Errorf("Load Env Err: %s", err)
	}
	return &LoadEnv{
		Err: err,
	}
}

func (self *LoadEnv) GetGenerateDir() (err error, generateDir string) {
	if self.Err!=nil {
		return self.Err, generateDir
	}
	generateDir = os.Getenv("generate_dir")

	return
}