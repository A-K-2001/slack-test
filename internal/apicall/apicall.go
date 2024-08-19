package apiCall

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)


type APIRequest struct {
    Method  string
    URL     string
    Body    interface{}
    Params  map[string]string
    Headers map[string]string
}

func MakeAPICall(req APIRequest) (*http.Response, error) {
    client := &http.Client{}

    // Add query 
    baseURL, err := url.Parse(req.URL)
    if err != nil {
        return nil, err
    }
    
    if req.Params != nil {
        params := url.Values{}
        for key, value := range req.Params {
            params.Add(key, value)
        }
        baseURL.RawQuery = params.Encode()
    }

    var bodyReader io.Reader
    if req.Body != nil {
        jsonBody, err := json.Marshal(req.Body)
        if err != nil {
            return nil, err
        }
        bodyReader = bytes.NewBuffer(jsonBody)
    }

    httpReq, err := http.NewRequest(req.Method, baseURL.String(), bodyReader)
    if err != nil {
        return nil, err
    }

    // fixed headers
    httpReq.Header.Add("Content-Type", "application/json")
    httpReq.Header.Add("Authorization", "Bearer ") // Replace with actual token

    // custom headers
    for key, value := range req.Headers {
        httpReq.Header.Add(key, value)
    }

    return client.Do(httpReq)
}