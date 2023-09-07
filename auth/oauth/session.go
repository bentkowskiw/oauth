package oauth

import (
	"encoding/json"
	"fmt"
)

func (s oAuthSession) MarshalBinary() ([]byte, error) {
	return json.Marshal(s.sessionData)
}
func (s *oAuthSession) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, &s.sessionData)
}
func (s oAuthSession) Key() string {
	return fmt.Sprintf("%s-%s", s.ProviderName, s.SessionUUID)
}

func newOAuthSession(providerName string, sessionId SessionUUID) *oAuthSession {
	return &oAuthSession{
		sessionData: sessionData{
			ProviderName: providerName,
			SessionUUID:  sessionId,
		},
	}
}
