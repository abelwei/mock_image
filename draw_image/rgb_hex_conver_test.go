package draw_image

import "testing"

func TestNewRgbHexConver(t *testing.T) {

	conver := NewRgbHexConver()

	err, rgb := conver.hex2rgb("668B8B")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("rgb:%+v", rgb)

	err, hex := conver.rgb2hex(RGB{
		Red:   222,
		Green: 33,
		Blue:  243,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("hex:%s", hex)

}
