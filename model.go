package sessionfrontend

import (
	"github.com/cdvelop/sessionhandler"
)

type sessionFrontend struct {
	Actual *sessionhandler.SessionStore

	*sessionhandler.Session
}
