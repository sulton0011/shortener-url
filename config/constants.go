package config

import "time"

const (
	// AccessTokenExpiresInTime ...
	AccessTokenExpiresInTime time.Duration = 20 * 365 * 24 * 60 * time.Minute // 175200h 0m 0s

	DatabaseQueryTimeLayout = `"YYYY-MM-DD"T"HH24:MI:SS"."MS"Z"TZ"` // "YYYY-MM-DD"T"HH24:MI:SS"."MS"Z"TZ"
	DatabaseTimeLayout      = time.RFC3339                          // 2006-01-02T15:04:05Z07:00
	FormatDateYM            = "2006-01"                             // 2006-01
	FormatDateYMD           = "2006-01-02"                          // 2006-01-02
	ErrEnvNodFound          = "No .env file found"
	SizeQrCode = 256 // size 256

	ExpiresDateUrl = 1 * 24 * 60 * time.Minute
)

const (
	SrvsUrl     = "UrlService"
	SrvsSession = "SessionService"
	SrvsUser    = "UserService"
)

const (
	PkgSecurity = "SecurityPkg"
)

const (
	StrgUrl  = "UrlStorage"
	StrgUser = "UserStorage"
)
