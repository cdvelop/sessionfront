package sessionfrontend

import "syscall/js"

func (s *sessionFrontend) submitLoginForm(this js.Value, btn []js.Value) interface{} {

	s.Log("boton login clikeado:", btn[0])

	return nil

}
