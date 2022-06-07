package Log

import (
	"fmt"
	"os"
	"php-in-go/Include/Http"
	FileSystem2 "php-in-go/Include/Util/FileSystem"
	"time"
)

type Log struct {
	LogPath string
}

func (l *Log) StartLogManager() {

}

func (l *Log) CloseLogManager() {

}

func (l *Log) Log(request *Http.Request, response *Http.Response) {
	filePrefix := time.Now().Format("0601")
	rootFile, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file := fmt.Sprintf("%s/%s%s/", rootFile, l.LogPath, filePrefix)

	if FileSystem2.IsFile(file) == false {
		FileSystem2.Mkdir(file)
	}
	f, err := os.OpenFile(
		fmt.Sprintf(
			"%s%s.log",
			file,
			time.Now().Format("02"),
		),
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0755,
	)

	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()
	s := fmt.Sprintf(
		"%s [%d] %s %s\n %s\n %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		response.Code,
		request.Method,
		request.RequestURI,
		response.ErrorMessage,
		response.ErrorStack,
	)
	_, err = f.Write([]byte(s))
	if err != nil {
		panic(err)
	}
}
