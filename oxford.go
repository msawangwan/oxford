package oxford

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ErrInvalidWord ..
var ErrInvalidWord = errors.New("Invalid Word")

// LanguageCode ..
type LanguageCode string

// EN is the language code for english
const EN LanguageCode = "en"

// Oxford is the client interface for calling the dictionary api
type Oxford struct {
	*http.Client
	*config
	lang LanguageCode
}

// New creates a new oxford client
func New(configPath string, lang LanguageCode) (*Oxford, error) {
	contents, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	cfg := &config{}

	if err = json.Unmarshal(contents, &cfg); err != nil {
		return nil, err
	}

	return &Oxford{
		Client: &http.Client{},
		config: cfg,
		lang:   lang,
	}, nil
}

// Exists => GET /inflections/{source_lang}/{word_id}
func (ox *Oxford) Exists(word string) (bool, error) {
	path := fmt.Sprintf("/inflections/%s/%s", ox.lang, strings.ToLower(word))
	uri := fmt.Sprintf("%s%s", ox.BaseURL, path)

	log.Printf("GET %s", path)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return false, err
	}

	ox.setAuth(req)

	res, err := ox.Do(req)
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 400 {
		return false, ErrInvalidWord
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	var payload WordExistsResponse

	if err = json.Unmarshal(body, &payload); err != nil {
		return false, err
	}

	log.Printf("%v", payload)

	return true, nil
}

func (ox *Oxford) setAuth(r *http.Request) {
	r.Header.Set("app_key", ox.Credentials.AppKey)
	r.Header.Set("app_id", ox.Credentials.AppID)
}

type config struct {
	BaseURL     string      `json:"base_url,omitempty"`
	Credentials credentials `json:"credentials,omitempty"`
}

type credentials struct {
	AppKey string `json:"app_key,omitempty"`
	AppID  string `json:"app_id,omitempty"`
}
