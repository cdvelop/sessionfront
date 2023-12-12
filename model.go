package sessionfrontend

import (
	"github.com/cdvelop/sessionhandler"
)

type sessionFrontend struct {
	current_session *sessionhandler.SessionStore

	*sessionhandler.Session
}
