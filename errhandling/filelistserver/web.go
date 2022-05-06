package main

import (
	"learngo/errhandling/filelistserver/filelist"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

// 函数式编程 统一错误处理
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	logger, _ := zap.NewProduction()
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			r := recover()
			if r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			logger.Warn("Error handling request: " + err.Error())
			code := http.StatusOK

			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelist.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
