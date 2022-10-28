package dto

type (
	GetCustomers struct {
		Name   string `json:"name"`
		UserId string `json:"user_id"`
		PaginateFilter
	}

	GetOrders struct {
		Item   string `json:"item"`
		Price  string `json:"price"`
		UserId string `json:"user_id"`
		PaginateFilter
	}

	PaginateFilter struct {
		SortField string `json:"sort_field"`
		SortOrder string `json:"sort_order"`
		First     int    `json:"first"`
		Rows      int    `json:"rows"`
	}
)
