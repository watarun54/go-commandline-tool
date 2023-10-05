package convert

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

const (
	JPEG = "jpeg"
	JPG  = "jpg"
	GIF  = "gif"
	PNG  = "png"
)

func ValidateExtension(ext string) error {
	switch ext {
	case JPEG, JPG, GIF, PNG:
		return nil
	default:
		return errors.New("サポートされていない拡張子です：" + ext)
	}
}

func ValidateFilepath(imagepath string) error {
	switch imagepath {
	case "":
		return errors.New("変換対象のファイルのパスを指定してください")
	default:
		if f, err := os.Stat(imagepath); os.IsNotExist(err) || f.IsDir() {
			return errors.New("ファイルが存在しません：" + imagepath)
		} else {
			return nil
		}
	}
}

func ValidateDestpath(destpath string) error {
	switch destpath {
	case "":
		return errors.New("変換後のファイル名を指定してください")
	default:
		return nil
	}
}

func ValidateFileExistence(imagepath string) error {
	switch imagepath {
	case "." + JPEG, "." + JPG, "." + GIF, "." + PNG:
		return nil
	default:
		return errors.New("指定したファイルが対応していません：" + imagepath)
	}
}

func Execute(extension string, imagepath string, destpath string) error {
	exFile, err := os.Open(imagepath)
	defer exFile.Close()
	if err != nil {
		return errors.New("os.Openに失敗しました")
	}

	output, err := os.Create(destpath)
	defer output.Close()
	if err != nil {
		return errors.New("os.Createに失敗しました")
	}

	img, _, Err := image.Decode(exFile)
	if Err != nil {
		return errors.New("image.Decodeに失敗しました")
	}

	switch extension {
	case JPEG, JPG:
		err = jpeg.Encode(output, img, nil)
		if err != nil {
			return errors.New("jpeg.Encodeに失敗しました")
		}
		fmt.Println("変換が正常に完了しました")
		return nil
	case GIF:
		err = gif.Encode(output, img, nil)
		if err != nil {
			return errors.New("gif.Encodeに失敗しました")
		}
		fmt.Println("変換が正常に完了しました")
		return nil
	case PNG:
		err = png.Encode(output, img)
		if err != nil {
			return errors.New("png.Encodeに失敗しました")
		}
		fmt.Println("変換が正常に完了しました")
		return nil
	}
	return nil
}
