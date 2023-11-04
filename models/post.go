package models

type Post struct {
	PostId      int
	SpeciesName string
	ImagePath   string
	Explain     string
	Lat         float64
	Lng         float64
}

type NewPost struct {
	PostId      int     `json:"id"`
	SpeciesName string  `json:"name"`
	ImageData   string  `json:"image_data"`
	Explain     string  `json:"explain"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}

type SearchRequest struct {
	Name string
}

type PostRequest struct {
	SessionId   string
	SpeciesName string
	ImageData   string
	Explain     string
	Lat         float64
	Lng         float64
}
