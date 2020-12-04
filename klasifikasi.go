package klasifikasi

import (
	"errors"
	"time"
)

var Cfg config = config{
	BaseUrl: "https://api.klasifikasi.com",
}

type Klasifikasi struct {
	modelMapping map[string]ModelMapping
}

var instance *Klasifikasi

func Build(credentials []ClientBuildParams) *Klasifikasi {

	if instance == nil {

		_modelMapping := map[string]ModelMapping{}

		for _, data := range credentials {

			clientAuth := requestToken(data)
			clientModel := getModelData(clientAuth)

			_modelMapping[clientModel.Model.PublicId] = ModelMapping{
				Auth:       clientAuth.Auth,
				Client:     clientModel.Model,
				ClientData: data,
			}

			instance = &Klasifikasi{
				modelMapping: _modelMapping,
			}
		}

	}

	return instance

}

func (ins *Klasifikasi) GetModels() map[string]ModelMapping {
	return instance.modelMapping
}

func (ins *Klasifikasi) Classify(publicId, query string) (ClassifyResponse, error) {
	var result ClassifyResponse
	var err error
	model, exist := ins.modelMapping[publicId]
	if !exist {
		return result, errors.New("Model not found !")
	}

	expired := time.Unix(0, int64(model.Auth.ExpiredAfter)*int64(time.Millisecond))

	if time.Now().After(expired) {
		clientAuth := requestToken(model.ClientData)
		model.Auth = clientAuth.Auth
	}

	result, err = classify(model.Auth, publicId, query)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ins *Klasifikasi) Logs(publicId string, params LogsParams) (LogsResponse, error) {
	var result LogsResponse
	var err error
	model, exist := ins.modelMapping[publicId]
	if !exist {
		return result, errors.New("Model not found !")
	}

	expired := time.Unix(0, int64(model.Auth.ExpiredAfter)*int64(time.Millisecond))

	if time.Now().After(expired) {
		clientAuth := requestToken(model.ClientData)
		model.Auth = clientAuth.Auth
	}

	result, err = logs(model.Auth, publicId, params)
	if err != nil {
		return result, err
	}
	return result, nil
}
