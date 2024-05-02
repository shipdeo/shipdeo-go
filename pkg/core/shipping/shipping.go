package shipping

type ShipdeoShippingClient interface {
	GetOngkir(data ShippingRequest) (*OngkirResponse, error)
}
