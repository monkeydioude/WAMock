package routing

import (
	"log"
	"regexp"
)

type Route struct {
	Method   Method `json:"-"`
	Path     string `json:"-"`
	Request  any    `json:"request"`
	Response any    `json:"response"`
}

func (r *Route) hydrateMethodPathFromRegRes(regRes []string) *Route {
	r.Method = ALL
	r.Path = "/"
	lenRegRes := len(regRes)
	// matched only one of the 2 parts. Must be a path, so trying to find out
	if lenRegRes == 1 {
		if len(regRes[0]) == 0 {
			return r
		}
		if _, err := SeekMethod(regRes[0]); err == nil {
			log.Println("Found a method, not a route")
			return r
		}
		r.Path = regRes[0]
	}
	if lenRegRes == 2 {
		if m, err := SeekMethod(regRes[0]); err == nil {
			r.Method = m
		}
		if regRes[1] != "" {
			r.Path = regRes[1]
		}
	}
	return r
}

func getMethodPathSlice(methodPath string) []string {
	reg, err := regexp.Compile("^([a-zA-Z]*)?(/.*)?$")
	if err != nil {
		log.Fatalf("NewFromMethodPathStr: %s\n", err)
	}
	strSlice := reg.FindStringSubmatch(methodPath)
	if len(strSlice) > 0 {
		strSlice = strSlice[1:]
	} else {
		strSlice = []string{}
	}
	return strSlice
}

func (r *Route) HydrateWithPath(path string) {
	r.hydrateMethodPathFromRegRes(getMethodPathSlice(path))
}

// NewFromMethodPathStr creates a Route struct from a string.
// This string comes from the config's JSON and can be:
// - a combination of an HTTP Method and a Route, such as "PUT/api/item"
// - just a route, such as "/api/items", which forces the Route struct's Method property to "ALL", accepting every HTTP Method
// - an empty string, which forces Route struct's Method to "ALL" and Path to "/"
func NewFromMethodPathStr(methodPath string) Route {
	route := Route{}
	route.hydrateMethodPathFromRegRes(getMethodPathSlice(methodPath))
	return route
}
