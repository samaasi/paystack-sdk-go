package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strconv"
)

const (
	userAgent = "paystack-sdk-go/1.0.0"
)

// RequestOptions contains options for the request
type RequestOptions struct {
	Headers map[string]string
}

// NewRequest creates a new HTTP request with the necessary headers
func NewRequest(method, urlStr string, apiKey string, body interface{}, opts *RequestOptions) (*http.Request, error) {
	var buf io.Reader
	if body != nil {
		var b []byte
		var err error
		b, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, urlStr, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("%s go/%s", userAgent, runtime.Version()))

	if opts != nil {
		for k, v := range opts.Headers {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

// EncodeQueryParams encodes a struct into URL query parameters
func EncodeQueryParams(v interface{}) (string, error) {
	if v == nil {
		return "", nil
	}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return "", nil
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return "", fmt.Errorf("EncodeQueryParams: expected struct, got %s", val.Kind())
	}

	values := url.Values{}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// Get query tag
		tag := field.Tag.Get("query")
		if tag == "" {
			tag = field.Name
		}
		if tag == "-" {
			continue
		}

		// Handle zero values (omitempty behavior logic could be added here if needed)
		if value.IsZero() {
			continue
		}

		switch value.Kind() {
		case reflect.String:
			values.Add(tag, value.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			values.Add(tag, strconv.FormatInt(value.Int(), 10))
		case reflect.Bool:
			values.Add(tag, strconv.FormatBool(value.Bool()))
		}
	}

	return values.Encode(), nil
}
