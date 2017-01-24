package model

import (
	"net/http"
	"io"
)

func CreateRequest(method string, url string, buffer io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, buffer);
	if (err != nil) {
		return nil, err;
	}

	return req, err

}
