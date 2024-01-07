package sessionfrontend

import (
	"github.com/cdvelop/model"
)

func (d sessionFrontend) runBootActions() {
	const t = "runBootActions error "
	// d.Log("CORRIENDO ACTIONS DATA DE ARRANQUE")

	// Establece el contenido del elemento meta a una cadena vacía
	js_value, err := d.SelectContent(model.SelectDomOptions{
		QuerySelector: "meta[name='JsonBootActions']",
		GetContent:    "content",
		SetAfter:      true,
		StringReturn:  true,
	})
	if err != "" {
		d.Log(t + err)
		return
	}

	content, ok := js_value.(string)
	if !ok {
		d.Log(t + "contenido seleccionado no retorno como string")
		return
	}

	// d.Log("CONTENDIDO JSON BOOT ok 2:", content)
	if content != "none" {

		err := d.FrontendLoadHtmlBootData(content)
		d.UserMessage(err)
		d.Log(err)

	}

}

func (d sessionFrontend) FrontendLoadHtmlBootData(data string) (err string) {
	const this = "FrontendLoadHtmlBootData error "

	if data == "" {
		return this + "sin data para cargar al sistema"
	}

	resp, err := d.DecodeResponses([]byte(data))
	if err != "" {
		return this + err
	}

	// fmt.Println("DATA SESSION", resp)

	d.addBootDataToLocalDB(resp...)

	for _, o := range d.GetAllObjects(false) {

		if o.FrontHandler.NotifyBootData != nil {
			o.FrontHandler.NotifyBootDataIsLoaded()
		}
	}

	return
}
