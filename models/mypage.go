package models

type MypageRequest struct {
	SessionId string
}

type Badge struct {
	ImageData string `json:"image_data"`
}

type MypageResponse struct {
	Name       string    `json:"name"`
	Posts      []NewPost `json:"posts"`
	BadgesData []Badge   `json:"badges"`
}
