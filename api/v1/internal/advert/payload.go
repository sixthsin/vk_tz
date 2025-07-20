package advert

type AdvertCreateRequest struct {
	Title       string `json:"title" binding:"required,min=5,max=100"`
	Description string `json:"description" binding:"omitempty,min=10,max=1000"`
	ImageURL    string `json:"image_url" binding:"omitempty,url"`
	Price       uint64 `json:"price" binding:"required,gt=0,lte=1000000000"`
}

type AdvertResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Price       uint64 `json:"price"`
	Author      string `json:"author"`
}

type AllAdvertsResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	IsMine    bool   `json:"is_mine"`
	AdvertResponse
}
