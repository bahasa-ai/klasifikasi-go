package klasifikasi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func requestToken(data ClientBuildParams) AuthData {

	payload, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/auth/token", Cfg.BaseUrl), bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var result AuthData
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	return result

}

func getModelData(data AuthData) ClientModel {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/auth/activeClient", Cfg.BaseUrl), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", data.Auth.Token))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var result ClientModel
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	return result

}

func classify(token TokenData, publicId, query string) ClassifyResponse {
	data := map[string]interface{}{"query": query}

	payload, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/classify/%s", Cfg.BaseUrl, publicId), bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var result ClassifyResponse
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	return result

}

func logs(token TokenData, publicId string, params LogsParams) LogsResponse {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/history/%s", Cfg.BaseUrl, publicId), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	q := req.URL.Query()
	q.Add("startedAt", params.StartedAt.Format(time.RFC3339))
	q.Add("endedAt", params.EndedAt.Format(time.RFC3339))
	q.Add("take", fmt.Sprint(params.Take))
	q.Add("skip", fmt.Sprint(params.Skip))
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var result LogsResponse
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	return result
}
