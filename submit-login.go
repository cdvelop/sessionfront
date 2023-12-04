package sessionfrontend

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (s *sessionFrontend) submitLoginForm(this js.Value, btn []js.Value) interface{} {

	s.Log("boton login clikeado")

	s.Log("DATA:", s.Form.FormData)

	s.Log(s.Session_status)

	// s.CreateObjectsInDB(s.Table, true, map[string]string{
	// 	s.Id_session:     "123",
	// 	s.Session_status: "in",
	// 	s.Session_encode: "xx",
	// })

	s.ReadStringDataAsyncInDB(model.ReadDBParams{
		FROM_TABLE:      s.Table,
		WHERE:           []string{s.Session_status},
		SEARCH_ARGUMENT: "in",
	}, func(result []map[string]string, err string) {
		if err != "" {
			s.UserMessage(err)
			return
		}

		if len(result) == 1 {
			s.Log("hay usuario en db local:", result)

		} else {
			s.Log("no hay usuario en local. enviando data al backend:", s.Form.FormData)

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

				err = s.FrontendLoadBootData(result[0]["boot_data"])
				if err != "" {
					s.UserMessage(err)
					return
				}

			})

		}

	})
	// form_name := f.html_form.Get("name").String()

	return nil

}
