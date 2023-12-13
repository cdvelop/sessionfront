package sessionfrontend

import "github.com/cdvelop/model"

func (s sessionFrontend) AddHeaderAuthentication() model.Header {

	// s.Log("info AddHeaderAuthentication:", s.HeaderAuthName)

	if s.Actual != nil {
		return model.Header{
			Name:    s.HeaderAuthName,
			Content: s.Actual.Session_encode,
		}
	}

	return model.Header{
		Name:    s.HeaderAuthName,
		Content: "",
	}

}
