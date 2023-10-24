package models

type Post struct {
	SpeciesName string
	ImagePath   string
	Explain     string
	Lat         float64
	Lng         float64
}

type NewPost struct {
	SpeciesName string
	ImageData   []byte
	Explain     string
	Lat         float64
	Lng         float64
}

type SearchRequest struct {
	Name string
}
