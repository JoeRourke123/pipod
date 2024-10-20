package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"net/http"
	"orchestrator/service/db"
	"orchestrator/service/spotify"
	"orchestrator/ui/model"
	"orchestrator/util/api"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Get(api.LoginRedirect(), redirectToSpotifyLogin)
	app.Get(api.CompleteAuth(), handleAuthComplete)
	app.Get(api.IsAuth(), handleIsAuth)
}

func redirectToSpotifyLogin(ctx *fiber.Ctx) error {
	return ctx.Redirect(spotify.GetAuthUrl())
}

func handleAuthComplete(ctx *fiber.Ctx) error {
	request, _ := adaptor.ConvertRequest(ctx, false)
	token, err := spotify.Auth.Token(ctx.Context(), spotify.AuthState, request)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Sorry could not build Spotify Token")
	}

	db.SetSpotifyToken(token)

	return ctx.SendString("Thanks!")
	//return ctx.Redirect("http://pipod.local:3000/login/success")
}

func handleIsAuth(ctx *fiber.Ctx) error {
	token := spotify.GenerateAccessToken(ctx.Context())

	authResponse := model.AuthResponse{
		HasToken: token != nil,
		AuthUrl:  api.Full(api.LoginRedirect()),
	}

	if authResponse.HasToken {
		authResponse.AccessToken = token.AccessToken
	}

	responseJson, _ := json.Marshal(authResponse)

	return ctx.SendString(string(responseJson))
}
