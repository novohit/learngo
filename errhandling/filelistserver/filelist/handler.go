package filelist

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type myUserError string

func (e myUserError) Error() string {
	return e.Message()
}

func (e myUserError) Message() string {
	return string(e)
}

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return myUserError("path must start with " + prefix)
	}

	path := request.URL.Path  // /list/errhandling/fib.txt
	path = path[len(prefix):] // /errhandling/fib.txt
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
