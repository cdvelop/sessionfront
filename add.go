package sessionfrontend

import (
	"syscall/js"

	"github.com/cdvelop/model"
	"github.com/cdvelop/sessionhandler"
)

func AddAuthAdapter(h *model.MainHandler, c *sessionhandler.Config) (err string) {

	s, err := sessionhandler.Add(h, c)
	if err != "" {
		return err
	}

	f := &sessionFrontend{
		Actual:  nil,
		Session: s,
	}
	s.Form.FrontHandler.FormNotify = f
	h.SessionFrontendAdapter = f
	h.FrontendBootDataUser = f

	s.Module.FrontendModuleHandlers.ClickedModuleEventAdapter = f

	js.Global().Set("submitLoginForm", js.FuncOf(f.submitLoginForm))

	return
}
