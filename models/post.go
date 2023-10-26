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
	PostId      int
	SpeciesName string
	ImageData   string
	Explain     string
	Lat         float64
	Lng         float64
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
