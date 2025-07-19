package advert

type AdvertCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"image_url" binding:"omitempty,url"`
	Price       uint64 `json:"price" binding:"required"`
}

type AdvertCreateResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Price       uint64 `json:"price"`
	Author      string `json:"author"`
}

type AdvertResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	IsMine    bool   `json:"is_mine"`
	AdvertCreateResponse
}
