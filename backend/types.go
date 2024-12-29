package main

type ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

type recipe struct {
	Title       string       `json:"title"`
	Ingredients []ingredient `json:"ingredients"`
	Steps       []string     `json:"steps"`
	ImageUrl    string       `json:"imageUrl"`
}

type requestBody struct {
	URL string `json:"url"`
}

type input struct {
	Url string `json:"url"`
}
