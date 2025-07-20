package advert

type AdvertFilter struct {
	Offset    int     `form:"offset" binding:"min=0"`
	Limit     int     `form:"limit" binding:"min=0,max=100"`
	SortBy    string  `form:"sort_by" binding:"omitempty,oneof=created_at price"`
	SortOrder string  `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	MinPrice  float64 `form:"min_price" binding:"min=0"`
	MaxPrice  float64 `form:"max_price" binding:"min=0"`
}
