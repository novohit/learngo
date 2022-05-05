package mock

type Retriever struct {
	Contents string
}

// mockretriever实现了两个接口，既可以Get() 又可以 Post()

// Get mock.Retriever结构实现了Retriever的接口
func (r Retriever) Get(url string) string {
	return r.Contents
}

// Post 实现了Post接口
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
