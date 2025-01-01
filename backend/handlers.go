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
	"github.com/pocketbase/pocketbase/core"
)

func handleCapture(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	input := input{}
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

func handleRecipe(e *core.RequestEvent) error {
	body, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Unable to read request body")
	}
	defer e.Request.Body.Close()

	var requestBody requestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid JSON format")
	}

	recipeUrl := requestBody.URL

	recipe := recipe{}
	switch {
	case strings.Contains(recipeUrl, "ica.se"):
		recipe = scrape(recipeUrl)
	case strings.Contains(recipeUrl, "arla.se"):
		recipe = scrapeArla(recipeUrl)
	}

	return e.JSON(http.StatusOK, recipe)
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
		if recipe.ImageUrl == "" {
			recipe.ImageUrl = e.Attr("src")
		}
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return recipe
}

func scrapeArla(url string) recipe {
	c := colly.NewCollector()

	recipe := recipe{}

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		recipe.Title = e.Text
	})
	c.OnHTML(".c-recipe__ingredients-group > table > tbody > tr", func(e *colly.HTMLElement) {
		amount := e.ChildText("td")
		name := e.ChildText("th > div")

		ingredient := ingredient{
			Name:   name,
			Amount: amount,
		}
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	})

	c.OnHTML("li.c-recipe__instructions-step", func(e *colly.HTMLElement) {
		recipe.Steps = append(recipe.Steps, e.Text)
	})

	c.OnHTML(".c-recipe__image-ratio-holder > img:nth-child(4)", func(e *colly.HTMLElement) {
		if recipe.ImageUrl == "" {
			recipe.ImageUrl = e.Attr("src")
		}
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(recipe)
	return recipe
}
