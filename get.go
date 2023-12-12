package sessionfrontend

import (
	"github.com/cdvelop/model"
)

func (s *sessionFrontend) FrontendCheckUser(result func(u *model.User, err string)) {

	result(s.DecodeUser(s.current_session.Session_encode))
}
