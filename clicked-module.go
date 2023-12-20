package sessionfrontend

func (s sessionFrontend) ClickedModuleEvent() {
	// s.Log("MODULO ", s.Form.ObjectName, "ha sido cliqueado realizar foco en input")
	s.Log(s.FormInputFocus(s.Form.ObjectName))
}
