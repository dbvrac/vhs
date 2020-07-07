package http

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	_http "net/http"
	"net/url"
)

// Request represents an HTTP request.
type Request struct {
	Method           string              `json:"method,omitempty"`
	URL              *url.URL            `json:"url,omitempty"`
	Proto            string              `json:"proto,omitempty"`
	ProtoMajor       int                 `json:"proto_major,omitempty"`
	ProtoMinor       int                 `json:"proto_minor,omitempty"`
	Header           map[string][]string `json:"header,omitempty"`
	Body             string              `json:"body,omitempty"`
	ContentLength    int64               `json:"content_length,omitempty"`
	TransferEncoding []string            `json:"transfer_encoding,omitempty"`
	Host             string              `json:"host,omitempty"`
	Trailer          map[string][]string `json:"trailer,omitempty"`
	RemoteAddr       string              `json:"remote_addr,omitempty"`
	RequestURI       string              `json:"request_uri,omitempty"`
}

// NewRequest creates a new Request.
func NewRequest(b []byte) (*Request, error) {
	rdr := bufio.NewReader(bytes.NewReader(b))

	req, err := _http.ReadRequest(rdr)
	if err != nil {
		return nil, fmt.Errorf("failed to read request: %w", err)
	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	return &Request{
		Method:           req.Method,
		URL:              req.URL,
		Proto:            req.Proto,
		ProtoMajor:       req.ProtoMajor,
		ProtoMinor:       req.ProtoMinor,
		Header:           req.Header,
		Body:             string(body),
		ContentLength:    req.ContentLength,
		TransferEncoding: req.TransferEncoding,
		Host:             req.Host,
		Trailer:          req.Trailer,
		RemoteAddr:       req.RemoteAddr,
		RequestURI:       req.RequestURI,
	}, nil
}