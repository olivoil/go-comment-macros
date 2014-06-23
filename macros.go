package macros

import (
	"regexp"
	"strings"
)

//
// Macro constructor.
//

func NewMacro() *Macro {
	return &Macro{":", []Handler{}}
}

//
// Macro
//

type Macro struct {
	prefix   string
	handlers []Handler
}

//
// Handler function type.
//

type Handler func(label string) string

//
// Register a new handler.
//

func (m *Macro) Register(handler Handler) {
	m.handlers = append(m.handlers, handler)
}

//
// Change default prefix.
//

func (m *Macro) Prefix(val string) {
	m.prefix = val
}

//
// Process `src` and return parsed string.
//

func (m *Macro) Process(src string) string {
	re := regexp.MustCompile(`\n([ \t]*\/\/.*)`)

	res := re.ReplaceAllStringFunc(src, func(str string) string {

		// save indent
		indent := strings.Split(str, "//")[0]

		// strip
		s := strings.Replace(str, "//", "", 1)
		s = strings.TrimSpace(s)

		// check prefix
		if strings.Index(s, m.prefix) == -1 {
			return str
		}

		// normalize
		s = strings.Replace(s, m.prefix, "", 1)
		s = strings.TrimSpace(s)

		// apply handlers
		ret := make([]string, len(m.handlers))
		for i := range m.handlers {
			tmp := m.handlers[i](s)
			if len(tmp) != 0 {
				ret[i] = indent + tmp
			}
		}

		return strings.Join(ret, "")
	})

	return res
}
