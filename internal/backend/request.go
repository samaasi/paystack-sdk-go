package backend

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

const (
	userAgent = "paystack-sdk-go/1.0.0"
)

// RequestOptions contains options for the request
type RequestOptions struct {
	Headers map[string]string
}

// NewRequest creates a new HTTP request with the necessary headers
func NewRequest(method, urlStr string, apiKey string, bodyBytes []byte, opts *RequestOptions) (*http.Request, error) {
	var buf io.Reader
	if bodyBytes != nil {
		buf = bytes.NewReader(bodyBytes)
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

		tag := field.Tag.Get("query")
		if tag == "" {
			tag = field.Name
		}
		if tag == "-" {
			continue
		}

		isPointer := value.Kind() == reflect.Ptr
		if isPointer {
			if value.IsNil() {
				continue
			}
			value = value.Elem()
		}

		tagParts := strings.Split(tag, ",")
		tagName := tagParts[0]
		omitempty := false
		for _, part := range tagParts[1:] {
			if part == "omitempty" {
				omitempty = true
			}
		}

		if omitempty && !isPointer && value.IsZero() {
			continue
		}

		switch value.Kind() {
		case reflect.String:
			values.Add(tagName, value.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			values.Add(tagName, strconv.FormatInt(value.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			values.Add(tagName, strconv.FormatUint(value.Uint(), 10))
		case reflect.Float32, reflect.Float64:
			values.Add(tagName, strconv.FormatFloat(value.Float(), 'f', -1, 64))
		case reflect.Bool:
			values.Add(tagName, strconv.FormatBool(value.Bool()))
		}
	}

	return values.Encode(), nil
}
