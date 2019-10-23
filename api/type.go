package api

// Weather site data
type Weather struct {
	City     string
	Temp     string
	Weather  string
	Air      string
	Humidity string
	Wind     string
	Limit    string
	Note     string
}

// One site info
type One struct {
	Date     string
	ImgURL   string
	Sentence string
}

// English info
type English struct {
	ImgURL   string
	Sentence string
}

// Poem info
type Poem struct {
	Title   string   `json:"title"`
	Dynasty string   `json:"dynasty"`
	Author  string   `json:"author"`
	Content []string `json:"content"`
}

// PoemRes response data
type PoemRes struct {
	Status string `json:"status"`
	Data   struct {
		Origin Poem `json:"origin"`
	} `json:"data"`
}

// Wallpaper data
type Wallpaper struct {
	Title  string
	ImgURL string
}

// Trivia info
type Trivia struct {
	ImgURL      string
	Description string
}
