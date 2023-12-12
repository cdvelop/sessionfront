package sessionfrontend

import "github.com/cdvelop/model"

func (s sessionFrontend) AddHeaderAuthentication() model.Header {

	return model.Header{
		Name:    s.H,
		Content: "",
	}

}
