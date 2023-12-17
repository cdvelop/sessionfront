package sessionfrontend

func (s sessionFrontend) UserSessionNumber() (number, err string) {
	return s.Actual.Session_number, ""
}
