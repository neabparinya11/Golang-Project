package payment

type(
	ItemServiceRequest struct{
		Items []*ItemServiceRequestDatum `json:"items" validate="required"`
	}

	ItemServiceRequestDatum struct{
		ItemId string `json:"item_id" validate:"required,max=64"`
		Price float64 `json:"price"`
	}
)