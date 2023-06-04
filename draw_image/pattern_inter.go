package draw_image

type PatternInter interface {
	SetParam(paramSs []string)
	SaveFile(filePath string) error
	ResponseWriter() (error, []byte)
}
