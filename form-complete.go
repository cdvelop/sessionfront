package sessionfrontend

import "syscall/js"

func (s *sessionFrontend) NotifyFormIsOK() {

	// s.Log("sessionFrontend NotifyFormIsOK:", s.Form.FormData)
	// ENCENDEMOS EL BOTÓN LOGIN
	s.ButtonLoginDisabled(false)
	// ENCENDEMOS LA ESCUCHA DEL TECLADO
	s.DevicePeripherals.KeyboardClientDisable(false)
}

func (s *sessionFrontend) NotifyFormERR() {
	// s.Log("sessionFrontend NotifyFormERR:", s.Form.FormData)
	s.ButtonLoginDisabled(true)
	s.DevicePeripherals.KeyboardClientDisable(true)
}

func (s sessionFrontend) ButtonLoginDisabled(disabled bool) {
	const e = "ButtonLoginDisabled "
	// s.Log(" ACTIVAR BOTÓN LOGIN:", s.Form.ObjectName)
	button, err := s.GetHtmlModuleContent(`button[data-object_name="` + s.Form.ObjectName + `"]`)

	if err != "" {
		s.Log(e + err)
		return
	}

	// Habilita el botón
	button.(js.Value).Set("disabled", disabled)
}
