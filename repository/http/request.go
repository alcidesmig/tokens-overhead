package http

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
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

var mode = os.Getenv("MODE")

func (h *httpRequestImpl) Request(token, address string) error {
	req, err := http.NewRequest(http.MethodGet, address, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("%s", token))
	dump, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		if mode == "unique" {
			log.Printf("HTTP request size: %dbytes\n", len(dump))
		}
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return nil
}
