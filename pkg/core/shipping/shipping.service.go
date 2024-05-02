package shipping

type ShipdeoShippingService struct {
	Client ShipdeoShippingClient
}

// NewShipdeoShippingService create a new instance of ShipdeoAuthService
func NewShipdeoShippingService(client ShipdeoShippingClient) *ShipdeoShippingService {
	return &ShipdeoShippingService{
		Client: client,
	}
}

// GetOngkir uses the ShipdeoShippingClient to get ongkir shipdeo.
func (s *ShipdeoShippingService) GetOngkir(paylod ShippingRequest) (*OngkirResponse, error) {
	return s.Client.GetOngkir(paylod)
}
