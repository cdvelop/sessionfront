package sessionfrontend

import "github.com/cdvelop/model"

func (d sessionFrontend) addBootDataToLocalDB(responses ...model.Response) {

	for _, r := range responses {

		// d.Log("domUpdate .... buscando objeto", r.Object)

		object, err := d.MainHandlerGetObjectByName(r.Object)
		if err != "" {
			d.Log(err)
			continue
		}

		if r.Action == "create" {
			err := d.CreateObjectsInDB(object.Table, false, r.Data)
			if err != "" {
				d.Log("error addBootDataToLocalDB", err)
				continue
			}
		}
	}
}
