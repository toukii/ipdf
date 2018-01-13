package imagick

import (
	"fmt"
	"gopkg.in/gographics/imagick.v2/imagick"
	// "gopkg.in/gographics/imagick.v3/imagick"
	"os"
	"path/filepath"
	"strings"
)

func PdfToImg(filePath, dirToSave string) (imgLists []string, err error) {
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return
	}

	if _, err = os.Stat(dirToSave); os.IsNotExist(err) {
		return
	}

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// count page of pdf
	if err = mw.PingImage(filePath); err != nil {
		return
	}
	pdfPages := mw.GetNumberImages()

	fileName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

	// convert pdf to image page by page
	for i := uint(0); i < pdfPages; i++ {
		pageName := fmt.Sprintf("%s[%v]", filePath, i)
		imgName := fmt.Sprintf("%s-%v.png", fileName, i)
		imgPath := filepath.Join(dirToSave, imgName)

		// clear resources associated with the wand
		mw.Clear()

		err = mw.ReadImage(pageName)
		if err != nil {
			return
		}

		mwc := mw.Clone()
		err = mwc.WriteImage(imgPath)
		if err != nil {
			return
		}

		imgLists = append(imgLists, imgName)
	}

	return
}
