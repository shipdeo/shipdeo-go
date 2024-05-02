package auth

// ShipdeoAuthClient defines the primary port used by the application to interact with external API services.
type ShipdeoAuthClient interface {
	ShipdeoAuthenticate() (*TokenResponse, error)
}

// TokenResponse is the output data from the API authentication.
type TokenResponse struct {
	AccessToken          string
	AccessTokenExpiresAt string
}
