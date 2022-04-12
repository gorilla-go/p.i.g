package Log

import (
	"fmt"
	"os"
	"php-in-go/Config"
	"php-in-go/Include/Http"
	Util2 "php-in-go/Include/Util"
	"time"
)

type Log struct {
}

func (l *Log) StartLogManager() {

}

func (l *Log) CloseLogManager() {

}

func (l *Log) Log(request *Http.Request, response *Http.Response) {
	fmt.Printf(
		"%s [%d] %s %s  %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		response.Code,
		request.Method,
		request.RequestURI,
		response.ErrorMessage,
	)
	filePrefix := time.Now().Format("0601")
	config := Config.App()
	path := config["logPath"].(string)
	rootFile, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file := fmt.Sprintf("%s/%s%s/", rootFile, path, filePrefix)

	if Util2.IsFile(file) == false {
		Util2.Mkdir(file)
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
