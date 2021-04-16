package storage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Headers struct {
	Authorization string `bson:"authorization" json:"authorization"`
}

type Body struct {
}

type Request struct {
	Url     string  `bson:"url" json:"url"`
	Method  string  `bson:"method" json:"method"`
	Headers Headers `bson:"headers" json:"headers"`
	//Body    Body    `bson:"body" json:"body"`
}

type ResponseBody struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func MakeHttpRequest(requestObject Request, responseBody *interface{}) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(requestObject.Method, requestObject.Url, nil)
	if err != nil {
		return nil, err
	}

	if requestObject.Headers.Authorization != "" {
		v := reflect.ValueOf(requestObject.Headers)
		typeOfHeaders := v.Type()
		for i := 0; i < v.NumField(); i++ {
			req.Header.Add(typeOfHeaders.Field(i).Name, v.Field(i).String())
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyData, &responseBody)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
