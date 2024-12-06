package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

type recipe struct {
	Title       string
	Ingredients []string
	Steps       []string
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

	recipeUrl := string(body)

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
		recipe.Ingredients = append(recipe.Ingredients, e.Text)
	})

	c.OnHTML("div.cooking-steps-card", func(e *colly.HTMLElement) {
		recipe.Steps = append(recipe.Steps, e.Text)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return recipe
}
