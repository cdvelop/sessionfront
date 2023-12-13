package sessionfrontend

import (
	"syscall/js"

	"github.com/cdvelop/model"
	"github.com/cdvelop/sessionhandler"
)

func AddAuthAdapter(h *model.Handlers, c *sessionhandler.Config) (err string) {

	s, err := sessionhandler.Add(h, c)
	if err != "" {
		return err
	}

	f := &sessionFrontend{
		Actual:  nil,
		Session: s,
	}

	h.AuthFrontendAdapter = f
	h.FrontendBootDataUser = f

	js.Global().Set("submitLoginForm", js.FuncOf(f.submitLoginForm))

	return
}
