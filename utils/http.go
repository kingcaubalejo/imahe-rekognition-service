package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Response(data interface{}, responseStatus int, w http.ResponseWriter) {
	jData, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	ErrorChecker(w.Write(jData))
}

func HttpReq(req *http.Request) (map[string]interface{}, string) {
	body, err := ioutil.ReadAll(req.Body)
	if err == nil {
		jsonData := make(map[string]interface{})
		ErrorChecker(0, json.Unmarshal(body, &jsonData))
		if len(jsonData) != 0 {
			return jsonData, ""
		}
		return nil, "Request parameters not found or empty."
	}
	return nil, "Invalid request."
}