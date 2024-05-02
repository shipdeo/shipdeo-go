package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shipdeo/shipdeo-go/pkg/api"
	"github.com/shipdeo/shipdeo-go/pkg/core/shipping"
)

func main() {
	httpClient := &http.Client{}

	appClient := &api.HttpShipdeoCoreClient{
		HttpClient: httpClient,
		BaseURL:    "https://main-api-development.shipdeo.com",
		Token:      "937540dac86c15046204cde3b52d5a5113da6bf2",
	}

	service := shipping.NewShipdeoShippingService(appClient)

	jsonStr := `
    {
        "couriers": ["ninja", "sicepat"],
        "is_cod": false,
        "origin_subdistrict_code": "31.71.01",
        "origin_subdistrict_name": "GAMBIR",
        "origin_city_code": "31.71",
        "origin_city_name": "KOTA ADM. JAKARTA PUSAT",
        "origin_lat": "-6.1737996",
        "origin_long": "106.8266873",
        "destination_subdistrict_code": "31.73.07",
        "destination_subdistrict_name": "PAL MERAH",
        "destination_city_code": "31.73",
        "destination_city_name": "KOTA ADM. JAKARTA BARAT",
        "destination_lat": null,
        "destination_long": null,
        "destination_note": "",
        "items": [
            {
                "value": 100000,
                "height": 11,
                "width": 11,
                "length": 11,
                "weight": 1,
                "qty": 1,
                "weight_uom": "kg",
                "dimension_uom": "cm",
                "name": "Baju Kerja",
                "description": "ini baju kerja",
                "is_wood_package": false
            }
        ]
    }`

	var request shipping.ShippingRequest
	err := json.Unmarshal([]byte(jsonStr), &request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response, err := service.GetOngkir(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Ongkir Result:", response.Data)

}
