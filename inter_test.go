package intersdk

import (
	"testing"
)

func TestInterInstance(t *testing.T) {
	_ = New(
		WithClientID("client_id"),
		WithClientSecret("client"),
		WithKeyFilePath("key.key"),
		WithCertFilePath("cert.crt"),
	)
}
