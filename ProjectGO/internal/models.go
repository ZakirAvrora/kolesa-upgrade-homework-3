package internal

const (
	Api        = "https://api.thecatapi.com/v1/images/search"
	DetailsApi = "https://api.thecatapi.com/v1/images/"
	BreedApi   = "https://api.thecatapi.com/v1/breeds"
	// https://api.thecatapi.com/v1/images/search?limit=10&breed_ids=beng
)

type CatBreed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CatPicture struct {
	Id    string `json:"id"`
	Image string `json:"url"`
}

type CatDetail struct {
	Id    string    `json:"id"`
	Image string    `json:"url"`
	Info  []Details `json:"breeds" `
}

type Details struct {
	Weight      Weight `json:"weight"`
	Name        string `json:"name"`
	Origin      string `json:"origin"`
	Temperament string `json:"temperament"`
	Description string `json:"description"`
	LifeSpan    string `json:"life_span"`
}

type Weight struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}
