package config

import (
	"encoding/json"
	"errors"
	"log"
	"time"
)

func (s *Settings) OAuth() *oAuth {
	return s.oAuth
}

func (o *oAuth) ConfigData(provider string) ([]byte, error) {
	i, ok := o.oAuthProviders[provider]
	if !ok {
		return nil, errors.New("provider not found")
	}
	return json.Marshal(i)
}
func (o *oAuth) SessionDuration(d time.Duration) time.Duration {
	if o.sessionDuration == 0 {
		return d
	}
	return o.sessionDuration
}

func (o *oAuth) UnmarshalJSON(b []byte) (err error) {
	s := struct {
		AuthProviders   map[string]any `json:"providers"`
		SessionDuration string         `json:"sessionDuration"`
	}{
		AuthProviders: make(map[string]any),
	}
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}
	o.oAuthProviders = s.AuthProviders
	o.sessionDuration, err = time.ParseDuration(s.SessionDuration)
	if err != nil {
		log.Println("invalid oauth session duration, will take default value")
		err = nil
	}
	return
}
