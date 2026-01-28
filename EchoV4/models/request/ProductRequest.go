package request

type ProductRequest struct {
	Name        string  `json:"name" form:"name" query:"name" validate:"required"`
	Description string  `json:"description" form:"description" query:"description"`
	Price       float64 `json:"price" form:"price" query:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" form:"stock" query:"stock" validate:"gte=0"`
	Category    string  `json:"category" form:"category" query:"category"`
}

type SearchProductRequest struct {
	Query    string   `query:"q"`
	Category string   `query:"category"`
	MinPrice float64  `query:"min_price"`
	MaxPrice float64  `query:"max_price"`
	Tags     []string `query:"tags"`
	Page     int      `query:"page"`
	Limit    int      `query:"limit"`
}
