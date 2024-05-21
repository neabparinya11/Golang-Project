package payment

type (
	ItemServiceRequest struct {
		Items []*ItemServiceRequestDatum `json:"items" validate="required"`
	}

	ItemServiceRequestDatum struct {
		ItemId string  `json:"item_id" validate:"required,max=64"`
		Price  float64 `json:"price"`
	}

	PaymentTranferRequest struct {
		PlayerId string  `json:"player_id"`
		ItemId   string  `json:"item_id"`
		Amount   float64 `json:"amount"`
	}

	PaymentTranferResponse struct {
		InventoryId   string `json:"inventory_id"`
		TransactionId string `json:"transaction_id"`
		PlayerId      string `json:"player_id"`
		ItemId string `json:"item_id"`
		Amount int64 `json:"amount"`
		Error string `json:"error"`
	}
)
