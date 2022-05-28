package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

//EncodeJSON serializes the response as a JSON object to the ResponseWriter
func EncodeJSON(w http.ResponseWriter, v interface{}, code int) error {
	//Any HTTP/1.1 message with an entity-body should include a Content-Type header field defining the media type of said body
	//Since there is no content, then there is no reason to specify a Content-Type header
	if code == http.StatusNoContent {
		w.WriteHeader(code)
		return nil
	}

	var jsonData []byte

	var err error
	switch v := v.(type) {
	case []byte:
		jsonData = v
	case io.Reader:
		jsonData, err = ioutil.ReadAll(v)
	default:
		jsonData, err = json.Marshal(v)
	}

	if err != nil {
		return err
	}

	//Set the content type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//Write the status code to the response and context
	w.WriteHeader(code)

	//Send the result back to the client
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
