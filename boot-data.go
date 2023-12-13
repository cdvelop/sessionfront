package sessionfrontend

func (d sessionFrontend) runBootActions() {
	const t = "runBootActions error "
	// d.Log("CORRIENDO ACTIONS DATA DE ARRANQUE")

	// Establece el contenido del elemento meta a una cadena vacía
	content, err := d.SelectContent("meta[name='JsonBootActions']", "content", true)
	if err != "" {
		d.Log(t + err)
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

	d.addBootDataToLocalDB(resp...)

	for _, o := range d.GetObjects() {

		if o.FrontHandler.NotifyBootData != nil {
			o.FrontHandler.NotifyBootDataIsLoaded()
		}
	}

	return
}
