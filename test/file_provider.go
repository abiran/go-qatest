package test

import (
	"encoding/json"
	"github.com/federicoleon/golang-restclient/rest"
	"net/http"
)

func GetFile(fileURL string) (*Header, *ApiError) {
	response := rest.Get(fileURL)
	if response == nil || response.Response == nil {
		return nil, &ApiError{
			Status:  http.StatusInternalServerError,
			Message: "invalid restclient response when trying to get json file",
		}
	}

	if response.StatusCode > 299 {
		var apiErr ApiError
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
			return nil, &ApiError{
				Status:  http.StatusInternalServerError,
				Message: "invalid error interface when getting json file",
			}
		}
		return nil, &apiErr
	}

	var result Header
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &ApiError{
			Status:  http.StatusInternalServerError,
			Message: "error when trying to unmarshal data for json file",
		}
	}
	return &result, nil
}
