package draw_image

import "testing"

func TestNewRgbHexConver(t *testing.T) {

	conver := NewRgbHexConver()

	err, rgb := conver.Hex2rgb("98F5FF")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("rgb:%+v", rgb)

	err, hex := conver.Rgb2hex(RGB{
		Red:   222,
		Green: 33,
		Blue:  243,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("hex:%s", hex)

}
