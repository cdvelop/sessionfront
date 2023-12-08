package sessionfrontend

import (
	"syscall/js"
)

func (s *sessionFrontend) submitLoginForm(t js.Value, btn []js.Value) interface{} {

	s.Submit()

	return nil

}
