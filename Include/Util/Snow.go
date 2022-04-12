package Util

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func Uuid() string {
	uuid := fmt.Sprintf("%d%f", time.Now().UnixNano(), rand.Float64())
	md5Obj := md5.New()
	_, err := io.WriteString(md5Obj, uuid)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", md5Obj.Sum(nil))
}
