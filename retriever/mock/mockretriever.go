package mock

type Retriever struct {
	Contents string
}

// Get mock.Retriever结构实现了Retriever的接口
func (r Retriever) Get(url string) string {
	return r.Contents
}
