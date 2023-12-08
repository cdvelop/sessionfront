package sessionfrontend

import (
	"github.com/cdvelop/model"
)

func (s *sessionFrontend) FrontendCheckUser(result func(u *model.User, err string)) {
	const this = "***TESTANDO FrontendCheckUser "

	s.ReadAsyncDataDB(model.ReadParams{
		FROM_TABLE:      s.Table,
		ID:              "",
		WHERE:           []string{s.Session_status},
		SEARCH_ARGUMENT: "in",
		ORDER_BY:        "",
		SORT_DESC:       false,
	}, func(r model.ReadResult) {

		if r.Error != "" {
			result(nil, this+r.Error)
			return
		}

		result(s.BuildUserFromStoreData(r.DataString))
	})
}

func (s *sessionFrontend) getLoginResult(result []map[string]string, err string) {

	// s.Log("ERROR:", err)
	// s.Log("result:", result)

	// if err != "" {
	// 	s.errorChannel <- err
	// } else {
	// 	// Procesa el resultado y colócalo en el canal
	// 	user := &model.User{
	// 		Token:          "XXXX",
	// 		Id:             "",
	// 		Ip:             "",
	// 		Name:           "FAKE",
	// 		Area:           "",
	// 		AccessLevel:    "",
	// 		LastConnection: "",
	// 	}

	// 	s.resultChannel <- user
	// }

}
