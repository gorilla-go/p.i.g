package Util

import (
	"os"
)

func Mkdir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
