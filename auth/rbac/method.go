package rbac

type Method int64

const (
	GET Method = iota
	POST
	PUT
	DELETE
	PATCH
	OPTIONS
	HEAD
	CONNECT
	TRACE
)

func (m Method) Get() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	case PATCH:
		return "PATCH"
	case OPTIONS:
		return "OPTIONS"
	case HEAD:
		return "HEAD"
	case CONNECT:
		return "CONNECT"
	case TRACE:
		return "TRACE"
	}
	return "unknown"
}
