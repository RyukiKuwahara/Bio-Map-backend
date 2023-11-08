package models

type SessionData struct {
	SessionId string
}

type BadgeInfo struct {
	BadgeId   int    `json:"badge_id"`
	ImageData string `json:"image_data"`
}

type MypageResponse struct {
	Name       string      `json:"name"`
	Posts      []NewPost   `json:"posts"`
	BadgesData []BadgeInfo `json:"badges"`
}
