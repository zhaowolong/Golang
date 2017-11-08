package main

import (
	"image"
	"image/color"
	"net/http"

	"git.code4.in/mobilegameserver/logging"

	"github.com/disintegration/imaging"
)

func main() {
	dst := imaging.New(100, 100, color.NRGBA{255, 255, 0, 128})
	/*
		files := []string{"image/1.png", "image/2.png", "image/3.png", "image/3.png", "image/3.png", "image/3.png", "image/3.png", "image/3.png", "image/3.png"}

		// load images and make 100x100 thumbnails of them
		var thumbnails []image.Image
		for _, file := range files {
			img, err := imaging.Open(file)
			if err != nil {
				panic(err)
			}
			thumb := imaging.Thumbnail(img, 40, 40, imaging.CatmullRom)
			thumbnails = append(thumbnails, thumb)
		}
	*/
	files := []string{"1.png", "2.png", "3.png", "3.png", "3.png", "3.png", "3.png", "3.png", "3.png"}
	var thumbnails []image.Image
	for _, _ = range files {
		res, _ := http.Get("http://14.17.104.56:8888/img/1000.png")
		//res, _ := http.Get("http://14.17.104.56:8888/img/25995c31685619e3eef2b2d6.jpg")
		img, err := imaging.Decode(res.Body)
		if err != nil {
			logging.Error("image.decode %s", err.Error())
			return
		}
		thumb := imaging.Thumbnail(img, 40, 40, imaging.CatmullRom)
		thumbnails = append(thumbnails, thumb)

	}
	newpic := group2pic(dst, thumbnails)
	err := imaging.Save(newpic, "2.jpg")
	if err != nil {
		panic(err)
	}

	newpic = group3pic(dst, thumbnails)
	err = imaging.Save(newpic, "3.jpg")
	if err != nil {
		panic(err)
	}

	newpic = group4pic(dst, thumbnails)
	err = imaging.Save(newpic, "4.jpg")
	if err != nil {
		panic(err)
	}

	newpic = group5pic(dst, thumbnails)
	err = imaging.Save(newpic, "5.jpg")
	if err != nil {
		panic(err)
	}

	newpic = group6pic(dst, thumbnails)
	err = imaging.Save(newpic, "6.jpg")
	if err != nil {
		panic(err)
	}

	newpic = group7pic(dst, thumbnails)
	err = imaging.Save(newpic, "7.jpg")
	if err != nil {
		panic(err)
	}

	newpic = group8pic(dst, thumbnails)
	err = imaging.Save(newpic, "8.jpg")
	if err != nil {
		panic(err)
	}

	newpic = group9pic(dst, thumbnails)
	err = imaging.Save(newpic, "9.jpg")
	if err != nil {
		panic(err)
	}

}

func group2pic(dst image.Image, group []image.Image) image.Image {
	for i, thumb := range group {
		group[i] = imaging.Thumbnail(thumb, 45, 45, imaging.CatmullRom)
	}
	for i, thumb := range group {
		dst = imaging.Paste(dst, thumb, image.Pt((i+1)*3+i*45, 23))
	}
	return dst
}

func group3pic(dst image.Image, group []image.Image) image.Image {
	for i, thumb := range group {
		group[i] = imaging.Thumbnail(thumb, 45, 45, imaging.CatmullRom)
	}
	for i, thumb := range group {
		if i == 0 {
			dst = imaging.Paste(dst, thumb, image.Pt(28, 3))
		} else {
			dst = imaging.Paste(dst, thumb, image.Pt((i)*3+(i-1)*45, 51))
		}
	}
	return dst
}

func group4pic(dst image.Image, group []image.Image) image.Image {
	for i, thumb := range group {
		group[i] = imaging.Thumbnail(thumb, 45, 45, imaging.CatmullRom)
	}
	for i, thumb := range group {
		if i < 2 {
			dst = imaging.Paste(dst, thumb, image.Pt((i+1)*3+i*45, 3))
		} else {
			dst = imaging.Paste(dst, thumb, image.Pt((i-1)*3+(i-2)*45, 51))
		}
	}
	return dst
}

func group5pic(dst image.Image, group []image.Image) image.Image {
	for i, thumb := range group {
		group[i] = imaging.Thumbnail(thumb, 30, 30, imaging.CatmullRom)
	}
	for i, thumb := range group {
		if i < 2 {
			dst = imaging.Paste(dst, thumb, image.Pt(18+i*33, 20))
		} else {
			dst = imaging.Paste(dst, thumb, image.Pt((i-1)*3+(i-2)*30, 53))
		}
	}
	return dst
}
func group6pic(dst image.Image, group []image.Image) image.Image {
	for i, thumb := range group {
		group[i] = imaging.Thumbnail(thumb, 30, 30, imaging.CatmullRom)
	}
	for i, thumb := range group {
		if i < 3 {
			dst = imaging.Paste(dst, thumb, image.Pt((i+1)*3+i*30, 13))
		} else {
			dst = imaging.Paste(dst, thumb, image.Pt((i-2)*3+(i-3)*30, 46))
		}
	}
	return dst
}
func group7pic(dst image.Image, group []image.Image) image.Image {
	for i, thumb := range group {
		group[i] = imaging.Thumbnail(thumb, 30, 30, imaging.CatmullRom)
	}
	for i, thumb := range group {
		if i < 1 {
			dst = imaging.Paste(dst, thumb, image.Pt(35, 2))
		} else if i < 4 {
			dst = imaging.Paste(dst, thumb, image.Pt(i*3+(i-1)*30, 34))
		} else {
			dst = imaging.Paste(dst, thumb, image.Pt((i-3)*3+(i-4)*30, 66))
		}
	}
	return dst
}
func group8pic(dst image.Image, group []image.Image) image.Image {
	for i, thumb := range group {
		group[i] = imaging.Thumbnail(thumb, 30, 30, imaging.CatmullRom)
	}
	for i, thumb := range group {
		if i < 2 {
			dst = imaging.Paste(dst, thumb, image.Pt(18+i*33, 2))
		} else if i < 5 {
			dst = imaging.Paste(dst, thumb, image.Pt((i-1)*3+(i-2)*30, 34))
		} else {
			dst = imaging.Paste(dst, thumb, image.Pt((i-4)*3+(i-5)*30, 66))
		}
	}
	return dst
}
func group9pic(dst image.Image, group []image.Image) image.Image {
	for i, thumb := range group {
		group[i] = imaging.Thumbnail(thumb, 30, 30, imaging.CatmullRom)
	}
	for i, thumb := range group {
		if i < 3 {
			dst = imaging.Paste(dst, thumb, image.Pt((i+1)*3+i*30, 2))
		} else if i < 6 {
			dst = imaging.Paste(dst, thumb, image.Pt((i-2)*3+(i-3)*30, 34))
		} else {
			dst = imaging.Paste(dst, thumb, image.Pt((i-5)*3+(i-6)*30, 66))
		}
	}
	return dst
}
