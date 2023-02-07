package draw_image

type PatternInter interface {
	SaveFile(filePath string) error
	ResponseWriter() (error, []byte)
}
