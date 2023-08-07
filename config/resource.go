package config

import (
	"io/fs"
	"log"
)

type prefixer interface {
	prefix() []string
}

func NewConfig(fs fs.FS) *Settings {
	env := &environment{
		variable: make(map[string]string),
	}
	env.read()

	s := &Settings{
		flags:       varFlags,
		environment: env,
		fs:          fs,
	}

	cfgPath := ""
	if cfgPathPtr, ok := s.flags.GetFlag(configFile); ok {
		cfgPath = *cfgPathPtr
	}
	if len(cfgPath) == 0 {
		cfgPath, _ = s.environment.get(configFile, nil)
	}
	if len(cfgPath) == 0 {
		cfgPath = defaultConfigFile
	}

	if err := s.readSettings(cfgPath); err != nil {
		log.Fatal(err)
		panic(err)
	}
	if err := s.readOauthConfig("oauth"); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return s
}
