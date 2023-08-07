package config

import (
	"flag"
	"fmt"
)

func flagKeys() []string {
	return []string{
		configFile,
	}
}

var varFlags *flags

func init() {
	varFlags = &flags{
		flags: make(map[string]*string),
	}
	varFlags.declareFlags()
}

func (env *flags) GetFlag(k string) (*string, bool) {
	value, ok := env.flags[k]
	if !ok || value == nil || len(*value) == 0 {
		return nil, false
	}
	return value, true
}

func (f *flags) declareFlags() {

	for _, k := range flagKeys() {
		f.flags[k] = flag.String(k, "", fmt.Sprintf("config file: '%s' ", k))
	}
	flag.Parse()
}
