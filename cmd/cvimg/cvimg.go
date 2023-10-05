package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/watarun54/go-commandline-tool/cmd/cvimg/convert"
)

var (
	extension string
	imagepath string
	destpath  string
)

func main() {
	flag.StringVar(&extension, "e", "jpeg", "拡張子の指定")
	flag.StringVar(&imagepath, "f", "", "変換対象のファイルのパスの指定")
	flag.StringVar(&destpath, "d", "", "変換後のファイル名の指定")
	flag.Parse()

	err := convert.ValidateExtension(extension)
	if err != nil {
		fmt.Println("[Error]", err)
		os.Exit(1)
	}

	err = convert.ValidateFilepath(imagepath)
	if err != nil {
		fmt.Println("[Error]", err)
		os.Exit(1)
	}

	err = convert.ValidateDestpath(destpath)
	if err != nil {
		fmt.Println("[Error]", err)
		os.Exit(1)
	}

	f := filepath.Ext(imagepath)
	err = convert.ValidateFileExistence(f)
	if err != nil {
		fmt.Println("[Error]", err)
		os.Exit(1)
	}

	err = convert.Execute(extension, imagepath, destpath)
	if err != nil {
		fmt.Println("[Error]", err)
		os.Exit(1)
	}
}
