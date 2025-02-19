package internal_api

import (
	"net/http"
)

type InternalAPI struct {
	HTTPClient *http.Client
	APIUrlBase string
	APIKey     string
}

func (i *InternalAPI) attachAPIKey(request *http.Request) {
	request.Header.Set("X-API-KEY", i.APIKey)
}
