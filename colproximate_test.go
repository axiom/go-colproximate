package colproximate

import "testing"
import "image/color"
import "github.com/axiom/rgbtxt"

func TestBlackAndWhite(t *testing.T) {
	pal := XTerm256

	if approx := pal.Convert(color.White).Hex(); approx != "#ffffff" {
		t.Errorf("Convert = %v, want %v", approx, "#ffffff")
	}

	if approx := pal.Convert(color.Black).Hex(); approx != "#000000" {
		t.Errorf("Convert = %v, want %v", approx, "#000000")
	}

	if approx := pal.Convert(rgbtxt.HotPink).Hex(); approx != "#ff69b4" {
		t.Errorf("Convert = %v, want %v", approx, "#ff69b4")
	}
}
