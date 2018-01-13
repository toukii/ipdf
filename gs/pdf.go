//PDF go image package driver.
package gs

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const pdfHeader = `%PDF`

func init() {
	image.RegisterFormat("pdf", pdfHeader, Decode, DecodeConfig)
}

//Encode SVG to PNG as image.Image
func decode(input []byte) (image.Image, error) {
	cmd := exec.Command("gs", "-r100x100", "-sDEVICE=png16m", "-dLastPage=1", "-q", `-sOutputFile=%stdout`, "-")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stdin = strings.NewReader(string(input))

	if err := cmd.Run(); err != nil {
		return nil, err
	}
	img, err := png.Decode(&out)
	if err != nil {
		return nil, err
	}
	return img, nil
}

//Decodes the first frame of an SVG file into an image
func Decode(r io.Reader) (image.Image, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return decode(b)
}

//Returns metadata
func DecodeConfig(r io.Reader) (image.Config, error) {
	return image.Config{}, errors.New("Not implemented")
}

func SaveImagePng(filename string, img image.Image) error {
	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	defer fd.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return png.Encode(fd, img)
}
