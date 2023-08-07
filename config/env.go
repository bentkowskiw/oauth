package config

import (
	"os"
	"strings"
)



func (e *environment) read() {

	// get env if exists
	for _, env := range os.Environ() {
		envData := strings.SplitN(env, "=", 2)
		if len(envData) == 2 {
			e.variable[envData[0]] = envData[1]
		}
	}
}

func (e *environment) get(key string, prefixer prefixer) (string, bool) {

	k := strings.Builder{}
	if prefixer != nil {

		for _, p := range prefixer.prefix() {
			k.WriteString(strings.ToUpper(p))
			k.WriteRune('_')
		}
	}
	k.WriteString(strings.ToUpper(key))
	v, ok := e.variable[k.String()]
	if ok && len(v) > 0 {
		return v, true
	}
	return "", false
}
