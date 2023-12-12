package sessionfrontend

import (
	"syscall/js"
)

func (s *sessionFrontend) submitLoginForm(t js.Value, btn []js.Value) interface{} {

	const this = "submitLoginForm error "

	// s.Log("DATA:", s.Form.FormData)

	s.Log(s.Session_status)

	if s.current_session != nil {

		s.Log("hay usuario cargado :", s.current_session.Id_session)

	} else {

		s.Log("no hay usuario en local. enviando data al backend:", s.Form.FormData)

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

			// SI NO HAY DATA DE ARRANQUE NO DETENGO EL FLUJO
			err = s.FrontendLoadBootData(result[0]["boot"])
			if err != "" {
				s.Log(err)
			}

			// DECODIFICAMOS LA SESIÓN PARA ALMACENARLA EN MEMORIA
			err = s.DecodeStruct([]byte(result[0]["session"]), s.current_session)
			if err != "" {
				s.UserMessage(this + err)
				return
			}

			// EJECUTAMOS LA CONSTRUCCIÓN DE LA UI
			err = s.BuildFrontendUI()
			if err != "" {
				s.UserMessage(this + err)
				return
			}
			//

		})

	}

	return nil

}
