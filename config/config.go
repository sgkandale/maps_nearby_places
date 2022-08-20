package config

type OutputFile struct {
	Name string
	Type string
}

type config struct {
	APIKey         string
	RadiusInMeters uint
	Keyword        string
	MinPrice       int
	MaxPrice       int
	Name           string
	OpenNow        bool
	RankBy         string
	PlaceType      string
	Latitude       float64
	Longitude      float64
	MaxRequests    int
	OutputFile     OutputFile
}
