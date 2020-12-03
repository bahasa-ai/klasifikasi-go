package klasifikasi

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
				Auth:   clientAuth,
				Client: clientModel,
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
