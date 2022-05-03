package http

import (
	"fmt"
	"net/http"
	"tokens-overhead/repository"
)

type httpRequestImpl struct {
	Client *http.Client
}

func NewHTTPRequester() repository.RequestInterface {
	return &httpRequestImpl{
		Client: &http.Client{},
	}
}

func (h *httpRequestImpl) Request(token, address string) error {
	req, err := http.NewRequest(http.MethodGet, address, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("%s", token))

	resp, err := h.Client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return nil
}
