package config

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
)

const (
	configFile        = "config"
	defaultConfigFile = "/cfg/config.json"
)

// ReadSettings Reads values from the provided file and stores all keys
func (cfg *Settings) readSettings(filePath string) (err error) {

	file, err := cfg.fileExists(filePath)
	if err != nil {
		log.Fatal("NO CONFIG FOUND")
		os.Exit(1)
	}
	var content []byte
	if _, err = file.Read(content); err != nil {
		panic(err)
	}
	log.Println("Configuration successfully loaded from file: ", filePath)

	// Deserialize 1
	if err := cfg.Unmarshal(content); err != nil {
		return err
	}

	// Complete config
	cfg.generic = make(map[string]interface{})

	// Deserialize 2
	if err = json.Unmarshal(content, &cfg.generic); err != nil {
		return
	}

	cfg.cfg.Redis.environment = cfg.environment

	cfg.cfg.Secure.environment = cfg.environment

	cfg.cfg.Server.environment = cfg.environment

	cfg.cfg.Db.environment = cfg.environment

	return
}

func (cfg *Settings) readOauthConfig(filePath string) (err error) {
	file, err := cfg.fileExists(filePath)
	if err != nil {
		log.Fatal("NO OAUTH FOUND")
		os.Exit(1)
	}
	var content []byte
	if _, err = file.Read(content); err != nil {
		panic(err)
	}
	cfg.oAuth = &oAuth{
		oAuthProviders: make(map[string]any),
	}

	return cfg.oAuth.UnmarshalJSON(content)
}

func (c *Settings) Unmarshal(b []byte) error {
	c.cfg = &cfg{}
	return json.Unmarshal(b, c.cfg)
}

func (cfg *Settings) fileExists(filename string) (file fs.File, err error) {
	return cfg.fs.Open(filename)
}

func (s *Settings) DbConfig() *db {
	return s.cfg.Db
}

func (s *Settings) Cors() []string {
	return *s.cfg.CORS
}
