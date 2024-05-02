package shipping

type ShippingRequest struct {
	Couriers                   []string `json:"couriers"`
	IsCOD                      bool     `json:"is_cod"`
	OriginLat                  string   `json:"origin_lat"`
	OriginLong                 string   `json:"origin_long"`
	OriginCityCode             string   `json:"origin_city_code"`
	OriginCityName             string   `json:"origin_city_name"`
	OriginSubdistrictCode      string   `json:"origin_subdistrict_code"`
	OriginSubdistrictName      string   `json:"origin_subdistrict_name"`
	OriginPostalCode           string   `json:"origin_postal_code"`
	DestinationLat             string   `json:"destination_lat"`
	DestinationLong            string   `json:"destination_long"`
	DestinationCityCode        string   `json:"destination_city_code"`
	DestinationCityName        string   `json:"destination_city_name"`
	DestinationSubdistrictCode string   `json:"destination_subdistrict_code"`
	DestinationSubdistrictName string   `json:"destination_subdistrict_name"`
	DestinationPostalCode      string   `json:"destination_postal_code"`
	Items                      []Item   `json:"items"`
}

type Item struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Weight        int    `json:"weight"`
	WeightUOM     string `json:"weight_uom"`
	Qty           int    `json:"qty"`
	Value         int    `json:"value"`
	TotalValue    int    `json:"total_value"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	Length        int    `json:"length"`
	IsWoodPackage bool   `json:"is_wood_package"`
	DimensionUOM  string `json:"dimension_uom"`
}

type OngkirResponse struct {
	Errors []any `json:"errors"`
	Data   []struct {
		Courier          string  `json:"courier"`
		CourierCode      string  `json:"courierCode"`
		Duration         string  `json:"duration"`
		Price            int     `json:"price"`
		Service          string  `json:"service"`
		SupportCod       bool    `json:"supportCod"`
		Estimation       string  `json:"estimation"`
		InsuranceValue   int     `json:"insuranceValue"`
		WoodPackingValue int     `json:"woodPackingValue"`
		ReturnRate       float64 `json:"returnRate"`
		ReturnLevel      string  `json:"returnLevel"`
	} `json:"data"`
}
