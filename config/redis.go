package config

func (s *Settings) Redis() *redis {
	return s.cfg.Redis
}

func (r *redis) prefix() []string {
	return []string{
		"redis",
	}
}
func (r *redis) RedisDsn() string {
	if s, ok := r.get("dsn", r); ok {
		return s
	}
	return r.DSN
}
