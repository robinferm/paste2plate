package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	http.HandleFunc("/", handleHome)
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
