package config

import (
	"log"
	"os"
)

type prefixer interface {
	prefix() []string
}

func NewConfig() *Settings {
	env := &environment{
		variable: make(map[string]string),
	}
	env.read()
	s := &Settings{
		flags:       varFlags,
		environment: env,
	}
	dir := s.getFlagOrEnv(configPath, defaultConfigPath)
	s.fs = os.DirFS(dir)

	cfgPath := s.getFlagOrEnv(configFile, defaultConfigFile)
	if err := s.readSettings(cfgPath); err != nil {
		log.Fatal(err)
		panic(err)
	}
	oauthPath := s.getFlagOrEnv(oauthFile, defaultOauthFile)
	if err := s.readOauthConfig(oauthPath); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return s
}

func (s *Settings) getFlagOrEnv(name, defaultValue string) (cfgPath string) {
	if cfgPathPtr, ok := s.flags.GetFlag(name); ok {
		cfgPath = *cfgPathPtr
	}
	if len(cfgPath) == 0 {
		cfgPath, _ = s.environment.get(name, nil)
	}
	if len(cfgPath) == 0 {
		cfgPath = defaultValue
	}
	return
}
