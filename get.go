package sessionfrontend

import (
	"github.com/cdvelop/model"
)

func (s *sessionFrontend) FrontendCheckUser(result func(u *model.User, err string)) {

	if s.current_session == nil {
		result(&model.User{}, "error no hay usuario registrado")
		return
	}

	result(s.DecodeUser(s.current_session.Session_encode))
}
