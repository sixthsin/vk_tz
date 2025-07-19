package advert

import (
	"marketplace-api/config"
	"marketplace-api/pkg/middleware"
	"net/http"
	"strconv"

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
		adverts.GET("/", handler.GetAll)
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
		"advert": &AdvertCreateResponse{
			Title:       createdAdvert.Title,
			Description: createdAdvert.Description,
			ImageURL:    createdAdvert.ImageURL,
			Price:       createdAdvert.Price,
			Author:      createdAdvert.Author,
		},
	})
}

func (h *AdvertHandler) GetAll(c *gin.Context) {
	limitString := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitString)
	if err != nil || limit < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid limit parameter",
		})
		return
	}

	offsetString := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetString)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid offset parameter",
		})
		return
	}

	username, _ := c.Get("username")

	adverts := h.AdvertRepository.GetAdverts(limit, offset)
	for i := range adverts {
		if adverts[i] == username {
			adverts[i].IsMine = true
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"adverts": adverts,
	})
}
