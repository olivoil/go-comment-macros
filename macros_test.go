package macros

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	input, err := ioutil.ReadFile("./fixtures/simple.go")
	assert.NoError(t, err)

	m := NewMacro()

	m.Register(func(label string) string {
		return "fmt.Println(\"" + label + "\")"
	})

	m.Register(func(label string) string {
		if 0 == strings.Index(label, "start ") {
			return "metrics.Start(\"" + strings.Replace(label, "start ", "", 1) + "\")"
		}

		if 0 == strings.Index(label, "end ") {
			return "metrics.End(\"" + strings.Replace(label, "end ", "", 1) + "\")"
		}

		return ""
	})

	output := m.Process(string(input))

	expected, err := ioutil.ReadFile("./fixtures/simple_out.go")
	assert.NoError(t, err)

	assert.Equal(t, strings.TrimSpace(output), strings.TrimSpace(string(expected)))
}
