package main

import (
	"fmt"
	"net/http"

	"github.com/shipdeo/shipdeo-go/pkg/api"
	core "github.com/shipdeo/shipdeo-go/pkg/core/auth"
)

func main() {
	httpClient := &http.Client{}
	apiClient := &api.HttpShipdeoAuthClient{
		HttpClient:   httpClient,
		BaseURL:      "https://auth-api-development.shipdeo.com",
		ClientId:     "X9NO1Hd0rRaS3VTa",
		ClientSecret: "l6NexU9Ud8r2lC75",
	}

	authService := core.NewShipdeoAuthService(apiClient)

	token, err := authService.ShipdeoAuthenticate()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Access Token:", token.AccessToken)
}
