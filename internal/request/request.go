package request

import (
	"errors"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

var METHODS = map[string]int{
	"GET":     1,
	"POST":    2,
	"PUT":     3,
	"PATCH":   4,
	"DELETE":  5,
	"HEAD":    6,
	"OPTIONS": 7,
	"CONNECT": 8,
	"TRACE":   9,
}

/*
	POST /coffee HTTP/1.1
*/

func parseRequestLine(request string) (*RequestLine, error) {
	rawRequestLine := strings.Split(request, "\r\n")[0]
	if rawRequestLine == "" {
		return &RequestLine{}, errors.New("wrong request format")
	}
	requestLineParts := strings.Split(rawRequestLine, " ")
	if len(requestLineParts) < 3 {
		return &RequestLine{}, errors.New("wrong request line format")
	}

	method := requestLineParts[0]
	if METHODS[method] == 0 {
		return &RequestLine{}, errors.New("invalid method format")
	}
	requestTarget := requestLineParts[1]
	httpVersion := requestLineParts[2]
	httpVersionParts := strings.Split(httpVersion, "/")
	if len(httpVersionParts) != 2 {
		return &RequestLine{}, errors.New("invalid Http Version")
	}

	rl := &RequestLine{
		HttpVersion:   httpVersionParts[1], // 1.1
		RequestTarget: requestTarget,
		Method:        method,
	}

	return rl, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return &Request{}, err
	}

	reqLine, err := parseRequestLine(string(data))
	if err != nil {
		return &Request{}, err
	}

	return &Request{RequestLine: *reqLine}, nil
}
