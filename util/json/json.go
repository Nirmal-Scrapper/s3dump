package json

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type JSON struct {
	Jmap map[string]interface{}
}

//String to byte to json
func ParseString(data string) JSON {
	return Parse([]byte(data))
}

//byte to kind of json type(map)
func Parse(data []byte) (parsed JSON) {
	var f interface{}
	if err := json.Unmarshal(data, &f); err != nil {
		return
	}
	switch f.(type) {
	case []interface{}:
		//      log.Println("Found Array, Not Supported")
		return
	default:
		m := f.(map[string]interface{})
		parsed.Jmap = m
		return parsed
	}
}

//get value by key
func (jobj *JSON) GetString(k string) string {
	data := jobj.Jmap[k]
	if str, ok := data.(string); ok {
		return str
	}
	return ""
}

//Get string array by key
func (jobj *JSON) GetStringArray(k string) []string {
	v := jobj.Jmap[k]
	str := make([]string, 0)
	switch val := v.(type) {
	case []string: //String array
		for _, u := range val {
			str = append(str, fmt.Sprintf("%s", u))
		}
	case []interface{}: //String array
		for _, u := range val {
			switch u.(type) {
			case string:
				str = append(str, fmt.Sprintf("%s", u))
			}
		}
	default:
		log.Println(k, "is of a type I don't know how to handle ", reflect.TypeOf(v))
	}

	return str
}
