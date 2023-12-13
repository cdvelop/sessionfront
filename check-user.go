package sessionfrontend

import (
	"github.com/cdvelop/model"
)

func (s *sessionFrontend) FrontendCheckUser(result func(u *model.User, err string)) {

	if s.Actual == nil {
		result(nil, "error no hay usuario registrado")
		return
	}

	user, err := s.DecodeUser(s.Actual.Session_encode)

	// s.Log("info USER:", user.Name)

	result(user, err)
}
