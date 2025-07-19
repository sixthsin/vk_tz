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

func (repo *AdvertRepository) GetAdverts(limit, offset int) []*AdvertResponse {
	var adverts []*AdvertResponse
	repo.Db.
		Table("adverts").
		Where("deleted_at is null").
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Scan(&adverts)
	return adverts
}
