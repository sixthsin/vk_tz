package advert

import (
	"marketplace-api/config"
	"marketplace-api/pkg/middleware"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AdvertHandlerDeps struct {
	*AdvertRepository
	*config.Config
}

type AdvertHandler struct {
	*AdvertRepository
	*config.Config
}

func NewAdvertHandler(r *gin.RouterGroup, deps AdvertHandlerDeps) {
	handler := &AdvertHandler{
		Config:           deps.Config,
		AdvertRepository: deps.AdvertRepository,
	}

	adverts := r.Group("/adverts")
	{
		adverts.GET("/", middleware.OptionalAuthMiddleware(deps.Config), handler.GetAll)
		protected := adverts.Use(middleware.AuthMiddleware(deps.Config))
		{
			protected.POST("/", handler.Create)
		}
	}
}

func (h *AdvertHandler) Create(c *gin.Context) {
	var req AdvertCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	if req.ImageURL != "" {
		if !isValidImageURL(req.ImageURL) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid image URL format"})
			return
		}
	}

	username, _ := c.Get("username")

	createdAdvert, err := h.AdvertRepository.Create(&Advert{
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		Price:       req.Price,
		Author:      username.(string),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"advert": &AdvertResponse{
			Title:       createdAdvert.Title,
			Description: createdAdvert.Description,
			ImageURL:    createdAdvert.ImageURL,
			Price:       createdAdvert.Price,
			Author:      createdAdvert.Author,
		},
	})
}

func (h *AdvertHandler) GetAll(c *gin.Context) {
	var filter AdvertFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if filter.Limit == 0 {
		filter.Limit = 10
	}
	if filter.SortBy == "" {
		filter.SortBy = "created_at"
		filter.SortOrder = "desc"
	}

	adverts := h.AdvertRepository.GetAdverts(filter)

	username, exists := c.Get("username")
	if !exists {
		username = ""
	}

	for i := range adverts {
		adverts[i].IsMine = adverts[i].Author == username
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   adverts,
		"meta": gin.H{
			"offset": filter.Offset,
			"limit":  filter.Limit,
			"total":  len(adverts),
		},
	})
}
func isValidImageURL(imageURL string) bool {
	validExtensions := []string{".jpg", ".jpeg", ".png"}
	for _, extension := range validExtensions {
		if strings.HasSuffix(strings.ToLower(imageURL), extension) {
			return true
		}
	}
	return false
}
