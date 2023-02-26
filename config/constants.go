package config

import "time"

const (
	// AccessTokenExpiresInTime ...
	AccessTokenExpiresInTime time.Duration = 20 * 365 * 24 * 60 * time.Minute

	DatabaseQueryTimeLayout = `"YYYY-MM-DD"T"HH24:MI:SS"."MS"Z"TZ"` // "YYYY-MM-DD"T"HH24:MI:SS"."MS"Z"TZ"
	DatabaseTimeLayout      = time.RFC3339                          // 2006-01-02T15:04:05Z07:00
	FormatDateYM            = "2006-01"                             // 2006-01
	FormatDateYMD           = "2006-01-02"                          // 2006-01-02
	ErrEnvNodFound          = "No .env file found"
)

const (
	SrvsUrl     = "UrlService"
	SrvsSession = "SessionService"
	SrvsUser    = "UserService"
)

const (
	// PkgCronJob     = "CronJobPkg"
	// PkgIntegration = "IntegrationPkg"
	// PkgTelegram    = "TelegramPkg"
	PkgSecurity = "SecurityPkg"
)

const (
	StrgUrl  = "UrlStorage"
	StrgUser = "UserStorage"
)
