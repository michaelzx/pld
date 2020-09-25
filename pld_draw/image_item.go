package pld_draw

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/michaelzx/pld/pld_fs"
	"image"
)

type ImageItem struct {
	ImagePath string
	Width     int
	Height    int
	X         float64
	Y         float64
}

func (ii *ImageItem) Draw(ggc *gg.Context) (endX, endY float64) {
	if !pld_fs.Exists(ii.ImagePath) {
		panic(ii.ImagePath + "-->不存在")
	}
	img, err := gg.LoadImage(ii.ImagePath)
	if err != nil {
		panic(ii.ImagePath + "-->读取失败")
	}
	var imgResize *image.NRGBA
	if ii.Height == 0 {
		imgResize = imaging.Resize(img, ii.Width, 0, imaging.Lanczos)
	} else {
		imgResize = imaging.Fill(img, ii.Width, ii.Width, imaging.Center, imaging.Lanczos)
	}
	ggc.DrawImage(imgResize, int(ii.X), int(ii.Y))
	endX = ii.X + float64(imgResize.Rect.Max.X)
	endY = ii.Y + float64(imgResize.Rect.Max.Y)
	return
}
