package Config

import "strings"

type Loader map[string]interface{}

func (r Loader) Load(s string) interface{} {
	if _, ok := r[s]; !ok {
		panic("configuration not found: " + s)
	}
	return r[s]
}

func (r Loader) LoadPath(s string) Loader {
	empty := make(map[string]interface{})
	for k, v := range r {
		if strings.HasPrefix(k, s) {
			k = strings.Replace(k, s+".", "", -1)
			empty[k] = v
		}
	}
	return empty
}
