package net

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/omnis-org/omnis-rest-api/pkg/model"
	"github.com/omnis-org/omnis-server/config"
	log "github.com/sirupsen/logrus"
)

func InitDefaultTransport() {
	if config.GetConfig().RestApi.InsecureSkipVerify {
		log.Warning("http : insecure skip verify")
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
}

func get(rootPath string, pathS string, i interface{}) ([]byte, error) {
	var url string
	// Generic way of get protocol with API OmnIS
	url = fmt.Sprintf("%s/%s", config.GetRestApiStringUrl(), rootPath)
	switch v := i.(type) {
	case int32:
		url = fmt.Sprintf("%s/%s/%d", url, pathS, v)
	case string:
		url = fmt.Sprintf("%s/%s/%s", url, pathS, v)
	default:
		url = fmt.Sprintf("%s/%s", url, pathS)
	}

	res, err := http.Get(config.GetRestApiScheme() + path.Clean(url))

	if err != nil {
		return nil, fmt.Errorf("Get failed <- %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Error rest api : %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll failed <- %v", err)
	}
	return body, nil
}

func postBytes(rootPath string, pathS string, jsonB []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/%s", config.GetRestApiStringUrl(), rootPath, pathS)
	res, err := http.Post(config.GetRestApiScheme()+path.Clean(url), "application/json", bytes.NewBuffer(jsonB))
	if err != nil {
		return nil, fmt.Errorf("http.Post failed <- %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Error rest api : %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll failed <- %v", err)
	}

	return body, nil
}

func put(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest failed <- %v", err)
	}
	req.Header.Set("Content-Type", contentType)

	return http.DefaultClient.Do(req)
}

func delete(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest failed <- %v", err)
	}
	req.Header.Set("Content-Type", contentType)

	return http.DefaultClient.Do(req)
}

func putBytes(rootPath string, pathS string, id int32, jsonB []byte) error {
	url := fmt.Sprintf("%s/%s/%s/%d", config.GetRestApiStringUrl(), rootPath, pathS, id)
	res, err := put(config.GetRestApiScheme()+path.Clean(url), "application/json", bytes.NewBuffer(jsonB))
	if err != nil {
		return fmt.Errorf("put failed <- %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Error rest api : %s", res.Status)
	}

	return nil
}

func deleteID(rootPath string, pathS string, id int32) error {
	url := fmt.Sprintf("%s/%s/%s/%d", config.GetRestApiStringUrl(), rootPath, pathS, id)
	res, err := delete(config.GetRestApiScheme()+path.Clean(url), "application/json", nil)
	if err != nil {
		return fmt.Errorf("delete failed <- %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Error rest api : %s", res.Status)
	}

	return nil
}

func insertObject(rootPath string, o model.Object, apiPath string) (int32, error) {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal failed <- %v", err)
	}

	body, err := postBytes(rootPath, apiPath, jsonBytes)
	if err != nil {
		return 0, fmt.Errorf("postBytes failed <- %v", err)
	}

	var jsonID model.IdJSON
	err = json.Unmarshal(body, &jsonID)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal failed <- %v", err)
	}
	id32 := int32(jsonID.Id)

	return id32, nil
}

func updateObject(rootPath string, o model.Object, apiPath string, id int32) error {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return fmt.Errorf("json.Marshal failed <- %v", err)
	}

	err = putBytes(rootPath, apiPath, id, jsonBytes)
	if err != nil {
		return fmt.Errorf("putBytes failed <- %v", err)
	}

	return nil
}

func getObjects(rootPath string, apiPath string, i interface{}, objects model.Objects) error {
	data, err := get(rootPath, apiPath, i)
	if err != nil {
		return fmt.Errorf("get failed <- %v", err)
	}

	err = json.Unmarshal(data, &objects)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed <- %v", err)
	}
	return nil
}

func getObject(rootPath string, apiPath string, i interface{}, object model.Object) error {
	data, err := get(rootPath, apiPath, i)
	if err != nil {
		return fmt.Errorf("get failed <- %v", err)
	}

	err = json.Unmarshal(data, &object)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed <- %v", err)
	}
	return nil
}

func deleteObject(rootPath string, apiPath string, id int32) error {
	err := deleteID(rootPath, apiPath, id)
	if err != nil {
		return fmt.Errorf("deleteID failed <- %v", err)
	}

	return nil
}
