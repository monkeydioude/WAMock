package routing

import "errors"

type Method string

const (
	PUT    Method = "PUT"
	POST   Method = "POST"
	GET    Method = "GET"
	DELETE Method = "DELETE"
	ALL    Method = "ALL"
	NONE   Method = "NONE"
	PATCH  Method = "PATCH"
)

func (m Method) String() string {
	return string(m)
}

func (m Method) Match(method string) bool {
	return method == m.String()
}

// SeekMethod tries to match a string with the Method type
// If none found, returns a NONE Method and an error
func SeekMethod(method string) (Method, error) {
	methods := []Method{PUT, POST, GET, DELETE, ALL, PATCH}

	for _, m := range methods {
		if Method(method) == m {
			return m, nil
		}
	}
	return NONE, errors.New("no method found")
}
