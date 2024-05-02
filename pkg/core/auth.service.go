package core

// ShipdeoAuthService is the service responsible for handling authentication logic.
type ShipdeoAuthService struct {
	Client ShipdeoAuthClient
}

// NewShipdeoAuthService creates a new instance of ShipdeoAuthService.
func NewShipdeoAuthService(client ShipdeoAuthClient) *ShipdeoAuthService {
	return &ShipdeoAuthService{
		Client: client,
	}
}

// Authenticate uses the APIClient to authenticate and retrieve a token.
func (s *ShipdeoAuthService) Authenticate() (*TokenResponse, error) {
	return s.Client.Authenticate()
}
