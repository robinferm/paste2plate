package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/recipe", handleRecipe)
	http.HandleFunc("/capture", handleCapture)
	http.ListenAndServe(":5000", nil)
}

func handleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "home\n")
}

type Input struct {
	Url string `json:"url"`
}

func handleCapture(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	input := Input{}
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = captureScreenshot(input.Url)
	if err != nil {
		http.Error(w, "Error capturing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func captureScreenshot(url string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// chromedp.Sleep(2*time.Second),
		chromedp.FullScreenshot(&buf, 90),
	)
	if err != nil {
		return err
	}

	err = os.WriteFile("filename.png", buf, 0o644)
	if err != nil {
		return err
	}

	fmt.Printf("Screenshot saved to %s\n", "fileName.png")
	return nil
}

type ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

type recipe struct {
	Title       string       `json:"title"`
	Ingredients []ingredient `json:"ingredients"`
	Steps       []string     `json:"steps"`
	ImageUrl    string       `json:"imageurl"`
}

type RequestBody struct {
	URL string `json:"url"` // The field name should match the key in the JSON body
}

func handleRecipe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	recipeUrl := string(requestBody.URL)

	recipe := scrape(recipeUrl)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(recipe)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func scrape(url string) recipe {
	c := colly.NewCollector()

	recipe := recipe{}

	c.OnHTML("h1.recipe-header__title", func(e *colly.HTMLElement) {
		recipe.Title = e.Text
	})

	c.OnHTML("div.ingredients-list-group__card", func(e *colly.HTMLElement) {
		amount := e.ChildText(".ingredients-list-group__card__qty")
		name := e.Text
		name = strings.Replace(name, amount, "", 1)
		name = strings.TrimSpace(name)

		ingredient := ingredient{
			Name:   name,
			Amount: amount,
		}
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	})

	c.OnHTML("div.cooking-steps-card", func(e *colly.HTMLElement) {
		recipe.Steps = append(recipe.Steps, e.Text)
	})

	c.OnHTML("img.recipe-header__image", func(e *colly.HTMLElement) {
		recipe.ImageUrl = e.Attr("src")
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return recipe
}
