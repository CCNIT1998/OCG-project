package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BodyParser(arg interface{}, w http.ResponseWriter, r *http.Request) error {
	bodyType, err := ioutil.ReadAll(r.Body) //đọc body
	defer r.Body.Close()

	if err != nil {
		return err
	}
	//chuyển json thành model
	err = json.Unmarshal(bodyType, arg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return err
	}
	return nil
}
