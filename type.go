package klasifikasi

type ClientBuildParams struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type config struct {
	BaseUrl string
}

type AuthData struct {
	Auth  TokenData `json:"auth"`
	Error string    `json:"error"`
}

type TokenData struct {
	Token        string `json:"token"`
	ExpiredAfter int    `json:"expiredAfter"`
}

type ClientModel struct {
	Model Model `json:"model"`
}

type Model struct {
	Name     string `json:"name"`
	Tags     []Tag  `json:"tags"`
	PublicId string `json:"publicId"`
}

type Tag struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	DescriptionWeight int    `json:"descriptionWeight"`
}

type ModelMapping struct {
	Auth   AuthData
	Client ClientModel
}
