package advert

import (
	"marketplace-api/pkg/db"
)

type AdvertRepository struct {
	Db *db.Db
}

func NewAdvertRepository(db *db.Db) *AdvertRepository {
	return &AdvertRepository{
		Db: db,
	}
}

func (repo *AdvertRepository) Create(advert *Advert) (*Advert, error) {
	result := repo.Db.DB.Create(advert)
	if result.Error != nil {
		return nil, result.Error
	}
	return advert, nil
}

func (repo *AdvertRepository) GetAdverts(filter AdvertFilter) []*AllAdvertsResponse {
	var adverts []*AllAdvertsResponse

	query := repo.Db.
		Table("adverts").
		Where("deleted_at is null")

	if filter.SortBy != "" {
		order := filter.SortBy
		if filter.SortOrder != "" {
			order += " " + filter.SortOrder
		}
		query = query.Order(order)
	} else {
		query = query.Order("created_at desc")
	}

	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	if filter.MinPrice > 0 {
		query = query.Where("price >= ?", filter.MinPrice)
	}
	if filter.MaxPrice > 0 {
		query = query.Where("price <= ?", filter.MaxPrice)
	}

	query.Scan(&adverts)
	return adverts
}
