package constants

import (
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/people/v1"
)

const (
	AuthPrefix            = "Auth"
	RedisPrefix           = "Redis"
	CorsPrefix            = "CORS"
	OauthLoginURLDuration = "loginURLDuration"
	ConfigFilePath        = "./cfg/calendar.json"
)

// services
const (
	GoogleService = "google.com"
)

const (
	ApiPrefix = "/api"
	GasPrefix = "/gas"

	CredentialsFileName = "./cfg/credentials.json"
	AccesTokenSecret    = "access_token_secret"
	RefreshTokenSecret  = "refresh_token_secret"
	DSN                 = "dsn"
	CertFileFormat      = "./cfg/%s.cer"
	KeyFileFormat       = "./cfg/%s.key"
)

// default values
const (
	RequestTimeout time.Duration = 30 * time.Second
)

// scope functions
var (
	AppScopes = func() []string {
		return []string{
			calendar.CalendarReadonlyScope,
			calendar.CalendarEventsScope,
			people.ContactsScope,
			people.UserinfoProfileScope,
			people.UserEmailsReadScope,
			people.UserPhonenumbersReadScope,
		}
	}
)
