package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct{}

// Get 接收者用不到名字 可以省略名字
func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
