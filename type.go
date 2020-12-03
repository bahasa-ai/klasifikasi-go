package klasifikasi

import "time"

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
	Auth   TokenData
	Client Model
}

type ClassifyResponse struct {
	Result []TagResponse `json:"result"`
}

type TagResponse struct {
	Label string  `json:"label"`
	Score float32 `json:"score"`
}

type LogsParams struct {
	StartedAt time.Time
	EndedAt   time.Time
	Take      int
	Skip      int
}

type LogsResponse struct {
	Logs []Logs `json:"histories"`
}

type Logs struct {
	CreatedAt   string        `json:"createdAt"`
	UpdatedAt   string        `json:"updatedAt"`
	Id          int           `json:"id"`
	Query       string        `json:"query"`
	ModelResult []TagResponse `json:"modelResult"`
}
