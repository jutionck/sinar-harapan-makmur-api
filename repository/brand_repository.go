package repository

import (
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"gorm.io/gorm"
)

type BrandRepository interface {
	BaseRepository[model.Brand]
	ListBrandWithVehicle() ([]model.Brand, error)
	GetBrandWithVehicle(brandId string) (*model.Brand, error)
	CountByName(name string, id string) (int64, error)
}

type brandRepository struct {
	db *gorm.DB
}

func (b *brandRepository) Delete(id string) error {
	return b.db.Delete(&model.Brand{}, "id=?", id).Error
}

func (b *brandRepository) Get(id string) (*model.Brand, error) {
	var brand model.Brand
	result := b.db.First(&brand, "id=?", id).Error
	if result != nil {
		return nil, result
	}
	return &brand, nil
}

func (b *brandRepository) List() ([]model.Brand, error) {
	var brands []model.Brand
	result := b.db.Find(&brands).Error
	if result != nil {
		return nil, result
	}
	return brands, nil
}

func (b *brandRepository) Save(payload *model.Brand) error {
	return b.db.Save(payload).Error
}

func (b *brandRepository) Search(by map[string]interface{}) ([]model.Brand, error) {
	var brands []model.Brand
	result := b.db.Where(by).Find(&brands).Error
	if result != nil {
		return nil, result
	}
	return brands, nil
}

func (b *brandRepository) GetBrandWithVehicle(brandId string) (*model.Brand, error) {
	var brand model.Brand
	result := b.db.Preload("Brand").First(&brand, "id=?", brandId).Error
	if result != nil {
		return nil, result
	}
	return &brand, nil
}

func (b *brandRepository) ListBrandWithVehicle() ([]model.Brand, error) {
	var brands []model.Brand
	result := b.db.Preload("Brand").Find(&brands).Error
	if result != nil {
		return nil, result
	}
	return brands, nil
}

func (b *brandRepository) CountByName(name string, id string) (int64, error) {
	var count int64
	var result *gorm.DB
	if id != "" {
		result = b.db.Model(&model.Brand{}).Where("name ILIKE ? AND id <> ?", "%"+name+"%", id).Count(&count)
	} else {
		result = b.db.Model(&model.Brand{}).Where("name ILIKE ?", "%"+name+"%").Count(&count)
	}
	if err := result.Error; err != nil {
		return 0, err
	}
	return count, nil
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepository{db: db}
}
