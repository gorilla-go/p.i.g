package Map

func Merge(maps ...map[string]interface{}) map[string]interface{} {
	c := make(map[string]interface{})
	for _, mapItem := range maps {
		for k, v := range mapItem {
			c[k] = v
		}
	}
	return c
}
