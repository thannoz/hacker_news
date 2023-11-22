package form

import "regexp"

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type errs map[string][]string

func (e errs) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errs) First(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
