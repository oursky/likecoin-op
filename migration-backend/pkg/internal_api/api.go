package internal_api

import (
	"net/http"
)

type InternalAPI struct {
	HTTPClient *http.Client
	APIUrlBase string
}
