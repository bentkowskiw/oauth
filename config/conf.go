package config

import (
	"encoding/json"
	"io/fs"
	"log"
)

const (
	configFile        = "config"
	defaultConfigFile = "cfg/config.json"
)

// ReadSettings Reads values from the provided file and stores all keys
func (cfg *Settings) readSettings(filePath string) (err error) {
	content, err := cfg.readFile(filePath)
	if err != nil {
		return
	}
	log.Println("Configuration successfully loaded from file: ", filePath)

	// Deserialize 1
	if err := cfg.UnmarshalJSON(content); err != nil {
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

	cfg.cfg.Auth.environment = cfg.environment

	return
}

func (cfg *Settings) readOauthConfig(filePath string) (err error) {
	content, err := cfg.readFile(filePath)
	if err != nil {
		return
	}
	cfg.oAuth = &oAuth{
		cfg: cfg.cfg,
	}
	return cfg.oAuth.UnmarshalJSON(content)
}

func (c *Settings) UnmarshalJSON(b []byte) error {
	c.cfg = &cfg{}
	if err := json.Unmarshal(b, c.cfg); err != nil {
		return err
	}
	c.cfg.Auth.client = c.cfg.Client
	c.cfg.Auth.server = c.cfg.Server
	c.cfg.Auth.cors = c.cfg.CORS
	c.cfg.Auth.environment = c.environment
	return nil
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

func (cfg *Settings) readFile(fileName string) (b []byte, err error) {
	file, err := cfg.fileExists(fileName)
	if err != nil {
		log.Fatal("NO " + fileName + " FOUND")
		return
	}
	fi, err := file.Stat()
	if err != nil {
		return
	}
	b = make([]byte, fi.Size())
	_, err = file.Read(b)
	return
}
