package httpx

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-errors/errors"
)

// Ensure Response implements the Message interface.
var _ Message = &Response{}

// Response represents an HTTP response.
type Response struct {
	ConnectionID     string      `json:"connection_id,omitempty"`
	ExchangeID       int64       `json:"exchange_id"`
	Created          time.Time   `json:"created,omitempty"`
	Status           string      `json:"status,omitempty"`
	StatusCode       int         `json:"status_code,omitempty"`
	Proto            string      `json:"proto,omitempty"`
	ProtoMajor       int         `json:"proto_major,omitempty"`
	ProtoMinor       int         `json:"proto_minor,omitempty"`
	Header           http.Header `json:"header,omitempty"`
	Body             string      `json:"body,omitempty"`
	ContentLength    int64       `json:"content_length,omitempty"`
	TransferEncoding []string    `json:"transfer_encoding,omitempty"`
	Close            bool        `json:"close,omitempty"`
	Uncompressed     bool        `json:"uncompressed,omitempty"`
	Trailer          http.Header `json:"trailer,omitempty"`
	SessionID        string      `json:"session_id,omitempty"`
}

// GetConnectionID gets a connection ID.
func (r *Response) GetConnectionID() string { return r.ConnectionID }

// GetExchangeID gets an exchange ID.
func (r *Response) GetExchangeID() int64 { return r.ExchangeID }

// SetCreated sets the created timestamp
func (r *Response) SetCreated(created time.Time) { r.Created = created }

// SetSessionID sets the session ID
func (r *Response) SetSessionID(id string) { r.SessionID = id }

// NewResponse creates a new Response.
func NewResponse(b *bufio.Reader, connectionID string, exchangeID int64) (*Response, error) {
	res, err := http.ReadResponse(b, nil)
	if err != nil {
		return nil, errors.Errorf("failed to read response: %w", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Errorf("failed to read response body: %w", err)
	}

	return &Response{
		ConnectionID:     connectionID,
		ExchangeID:       exchangeID,
		Status:           res.Status,
		StatusCode:       res.StatusCode,
		Proto:            res.Proto,
		ProtoMajor:       res.ProtoMajor,
		ProtoMinor:       res.ProtoMinor,
		Header:           res.Header,
		Body:             string(body),
		ContentLength:    res.ContentLength,
		TransferEncoding: res.TransferEncoding,
		Close:            res.Close,
		Uncompressed:     res.Uncompressed,
		Trailer:          res.Trailer,
	}, nil
}