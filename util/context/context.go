package context

type handle struct {
	Property map[string]interface{}
}

var instantiated *handle

//initiation is done and ensured
func Instance() *handle {
	if instantiated == nil {
		instantiated = new(handle)
		instantiated.Property = make(map[string]interface{})
	}
	return instantiated
}

//Get val by key
func (h *handle) Get(key string) string {
	if v, ok := h.Property[key]; ok {
		return v.(string)
	}
	return ""
}

//Set a key value pair
func (h *handle) Set(key string, value string) {
	h.Property[key] = value
	//fmt.Println(h)
}
