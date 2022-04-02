package Component

type RequestMethod uint8

const (
	GET RequestMethod = iota
	POST
	PUT
	DELETE
)

func (m RequestMethod) ToString() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	}

	return ""
}
