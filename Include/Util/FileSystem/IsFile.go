package FileSystem

import (
	"os"
)

func IsFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}
