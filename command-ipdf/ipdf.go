package ipdf

import (
	"fmt"
	"os"
	"path/filepath"

	cr "github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toukii/goutils"
	"github.com/toukii/ipdf/gs"
	"github.com/toukii/ipdf/imagick"
)

var Command = &cobra.Command{
	Use:   "ipdf",
	Short: "pdf 2 image",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		size := len(args)
		if size > 0 {
			viper.Set("pdf", args[0])
		} else {
			return
		}
		if err := Excute(); err != nil {
			cr.Red("%+v", err)
		}
	},
}

const (
	ModeGs      = "gs"
	ModeImagick = "imagick"
)

func init() {
	Command.PersistentFlags().StringP("image", "i", "png", "image:png|jpeg")
	Command.PersistentFlags().StringP("mode", "m", "gs", "mode:gs|imagick")
	Command.PersistentFlags().StringP("output", "o", ".", "output dir")

	viper.BindPFlag("image", Command.PersistentFlags().Lookup("image"))
	viper.BindPFlag("mode", Command.PersistentFlags().Lookup("mode"))
}

func Excute() error {
	if viper.GetString("mode") == ModeGs {
		pdfile := viper.GetString("pdf")
		fd, err := os.Open(pdfile)
		if goutils.CheckNoLogErr(err) {
			return err
		}

		img, err := gs.Decode(fd)
		if goutils.CheckNoLogErr(err) {
			return err
		}
		output := filepath.Join(viper.GetString("output"), pdfile+"."+viper.GetString("image"))
		return gs.SaveImagePng(output, img)
	} else {
		imgs, err := imagick.PdfToImg(viper.GetString("pdf"), viper.GetString("output"))
		if goutils.CheckNoLogErr(err) {
			return err
		}
		fmt.Println(imgs)
	}
	return nil
}
