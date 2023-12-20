package sessionfrontend

import (
	"syscall/js"

	"github.com/cdvelop/sessionhandler"
)

func (s *sessionFrontend) submitLoginForm(t js.Value, btn []js.Value) any {
	s.submitForm()
	return nil
}

func (s *sessionFrontend) KeyEnter() {
	s.submitForm()
	// s.Log("KeyEnter sessionFrontend", s.Form.ObjectName)
}

func (s *sessionFrontend) submitForm() {
	s.ButtonLoginDisabled(true)
	s.DevicePeripherals.KeyboardClientDisable(true)

	const e = "submitLoginForm "

	// s.Log("info. enviando data al backend:", s.Form.FormData)

	s.SendOneRequest("POST", "create", s.Form.ObjectName, s.Form.FormData, func(result []map[string]string, err string) {
		if err != "" {
			s.UserMessage(err)
			return
		}
		// s.Log("RESULTADO SESIÓN:", result)

		if len(result) != 1 {
			s.UserMessage("error se esperaba data para inicio de sesión")
			return
		}

		// DECODIFICAMOS LA SESIÓN PARA ALMACENARLA EN MEMORIA
		s.Actual = &sessionhandler.SessionStore{
			Id_session:     result[0][s.Id_session],
			Session_number: result[0][s.Session_number],
			Session_encode: result[0][s.Session_encode],
		}

		// SI NO HAY DATA DE ARRANQUE HTML NO DETENGO EL FLUJO
		s.Log(s.FrontendLoadHtmlBootData(result[0]["boot"]))
		// s.Log("SESIÓN STORE", s.Actual.Id_session)

		// limpiamos el formulario login
		err = s.FormClean()
		if err != "" {
			s.Log(e + err)
		}

		// EJECUTAMOS LA CONSTRUCCIÓN DE LA UI
		err = s.BuildFrontendUI()
		if err != "" {
			s.UserMessage(e + err)
			return
		}

	})
}
