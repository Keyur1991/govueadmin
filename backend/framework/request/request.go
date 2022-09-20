package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestFactory interface {
	Request() Request
}

func Factory(r *http.Request) (Request, error) {

	cp := r.Header.Get("Content-Type")

	if cp == "application/json" {
		rf := &JSON{r}
		return rf.Request(), nil
	} else if cp == "application/x-www-form-urlencoded" {
		rf := &UrlEncoded{r}
		return rf.Request(), nil
	}

	return nil, fmt.Errorf("Unknown content type")
}

type JSON struct {
	r *http.Request
}

type UrlEncoded struct {
	r *http.Request
}

func (j JSON) Request() Request {
	return &JsonRequest{
		r: j.r,
	}
}

func (u UrlEncoded) Request() Request {
	return &UrlEncodedRequest{
		r: u.r,
	}
}

type Request interface {
	Parse() error
	Get(key string) string
	GetSlice(key string) []string
}

type JsonRequest struct {
	r    *http.Request
	body map[string]interface{}
}

type UrlEncodedRequest struct {
	r *http.Request
}

func (jr *JsonRequest) Parse() error {
	return json.NewDecoder((*jr).r.Body).Decode(&jr.body)
}

func (jr *JsonRequest) Get(key string) string {
	// get key value string
	if val, ok := jr.body[key]; ok {
		return val.(string)
	}
	return ""
}

func (jr *JsonRequest) GetSlice(key string) []string {
	// get key value slice string
	if val, ok := jr.body[key]; ok {
		return val.([]string)
	}
	return nil
}

func (jr *JsonRequest) GetStruct(key string) interface{} {
	// get key value struct
	if val, ok := jr.body[key]; ok {
		return val
	}
	return nil
}

func (ur *UrlEncodedRequest) Parse() error {
	return ur.r.ParseForm()
}

func (ur *UrlEncodedRequest) Get(key string) string {
	// get key value string
	return ur.r.FormValue(key)
}

func (ur *UrlEncodedRequest) GetSlice(key string) []string {
	// get key valye slice string
	if val, ok := ur.r.Form[key]; ok {
		return val
	}
	return nil
}
