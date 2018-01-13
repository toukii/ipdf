package gs

import (
	"image/png"
	"os"
	"testing"
)

const dataDir = "testData/"
const input = dataDir + "pdf.pdf"
const output = input + ".png"

func TestPDF(t *testing.T) {
	if _, err := os.Stat(input); err != nil {
		t.Fatal(err)
	}
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	img, err := Decode(file)
	if err != nil {
		t.Fatal(err)
	}
	out, err := os.Create(output)
	if err != nil {
		t.Fatal(err)
	}
	if err := png.Encode(out, img); err != nil {
		t.Fatal(err)
	}
}
