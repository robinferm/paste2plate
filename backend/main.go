package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()
	app.OnRecordAuthWithOAuth2Request().BindFunc(func(e *core.RecordAuthWithOAuth2RequestEvent) error {
		e.Response.Header().Set("Cross-Origin-Opener-Policy", "same-origin-allow-popups")
		return e.Next()
	})
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/api/recipe", handleRecipe).Bind(apis.RequireAuth())
		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
