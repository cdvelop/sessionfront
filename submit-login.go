package sessionfrontend

import (
	"syscall/js"
)

func (s *sessionFrontend) submitLoginForm(t js.Value, btn []js.Value) interface{} {

	const this = "submitLoginForm "

	// s.Log("info. enviando data al backend:", s.Form.FormData)

	s.SendOneRequest("POST", "create", s.Form.ObjectName, s.Form.FormData, func(result []map[string]string, err string) {

		if err != "" {
			s.UserMessage(err)
			return
		}

		s.Log("RESULTADO SESIÓN:", result)

		if len(result) != 1 {
			s.UserMessage("error se esperaba data para inicio de sesión")
			return
		}

		// SI NO HAY DATA DE ARRANQUE HTML NO DETENGO EL FLUJO
		s.Log(s.FrontendLoadHtmlBootData(result[0]["boot"]))

		// DECODIFICAMOS LA SESIÓN PARA ALMACENARLA EN MEMORIA
		err = s.DecodeStruct([]byte(result[0]["session"]), &s.Actual)
		if err != "" {
			s.UserMessage(this + err)
			return
		}

		// s.Log("SESION STORE", s.Actual.Id_session)

		// EJECUTAMOS LA CONSTRUCCIÓN DE LA UI
		err = s.BuildFrontendUI()
		if err != "" {
			s.UserMessage(this + err)
			return
		}
		//

	})

	return nil

}
