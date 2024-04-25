package intersdk

import "errors"

var (
	ErrTlsCertificateNil = errors.New("tls certificate not provided")
	ErrOauthFailed       = errors.New("oauth failed")
)
