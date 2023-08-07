package config

import "strconv"

func (s *Settings) DB() *db {
	return s.cfg.Db
}

func (db *db) prefix() []string {
	return []string{"db"}
}

func (db *db) GetDbUser() string {
	if s, ok := db.get("user", db); ok {
		return s
	}
	return db.User
}
func (db *db) GetDbName() string {
	if s, ok := db.get("name", db); ok {
		return s
	}
	return db.Name
}
func (db *db) GetDbPassword() string {
	if s, ok := db.get("password", db); ok {
		return s
	}
	return db.Password
}
func (db *db) GetDbHost() string {
	if s, ok := db.get("host", db); ok {
		return s
	}
	return db.Host
}
func (db *db) GetDbPort() int {
	if s, ok := db.get("port", db); ok {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return i
	}
	return db.Port
}

func (db *db) GetDbSchema() string {
	if s, ok := db.get("schema", db); ok {
		return s
	}
	return db.Schema
}
func (db *db) GetDbSSL() string {
	if s, ok := db.get("ssl", db); ok {
		return s
	}
	return db.Ssl
}
func (db *db) GetDbDriver() string {
	if s, ok := db.get("driver", db); ok {
		return s
	}
	return db.Driver
}
