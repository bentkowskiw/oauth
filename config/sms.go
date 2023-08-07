package config

import (
	"encoding/base64"
	"encoding/json"
)

type smsConfig struct{}

func (s *sms) prefix() []string {
	return []string{"sms"}
}

func (sms *sms) GetConfigSMS(locale string) (c *smsConfig, err error) {
	s, ok := sms.get(locale, sms)

	var b []byte
	switch ok {
	case true:
		b, err = base64.StdEncoding.DecodeString(s)
		if err != nil {
			return nil, err
		}
	default:
		b, err = json.Marshal(sms.Config[locale])
		if err != nil {
			return nil, err
		}
	}
	c = &smsConfig{}
	err = json.Unmarshal(b, c)
	return
}
