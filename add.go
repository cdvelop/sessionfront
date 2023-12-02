package sessionfrontend

import (
	"syscall/js"

	"github.com/cdvelop/model"
	"github.com/cdvelop/sessionhandler"
)

func AddAuthAdapter(h *model.Handlers) (err string) {

	s, err := sessionhandler.Add(h)
	if err != "" {
		return err
	}

	f := &sessionFrontend{
		Session: s,
	}

	h.AuthAdapter = f

	js.Global().Set("submitLoginForm", js.FuncOf(f.submitLoginForm))

	return
}
