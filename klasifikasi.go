package klasifikasi

import (
	"errors"
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
				Auth:   clientAuth.Auth,
				Client: clientModel.Model,
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
	model, exist := ins.modelMapping[publicId]
	if !exist {
		return result, errors.New("Model not found !")
	}
	result = classify(model.Auth, publicId, query)
	return result, nil
}

func (ins *Klasifikasi) Logs(publicId string, params LogsParams) (LogsResponse, error) {
	var result LogsResponse
	model, exist := ins.modelMapping[publicId]
	if !exist {
		return result, errors.New("Model not found !")
	}
	result = logs(model.Auth, publicId, params)
	return result, nil
}
