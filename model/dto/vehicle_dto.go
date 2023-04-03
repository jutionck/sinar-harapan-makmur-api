package dto

import "github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"

type VehiclePayloadDto struct {
	Id             string `json:"id"`
	BrandID        string `json:"brandId"`
	Model          string `json:"model"`
	ProductionYear int    `json:"productionYear"`
	Color          string `json:"color"`
	IsAutomatic    bool   `json:"isAutomatic"`
	Stock          int    `json:"stock"`
	SalePrice      int    `json:"salePrice"`
	Status         string `json:"status"`
}

type VehicleDto struct {
	Id             string      `json:"id"`
	Brand          model.Brand `json:"brand"`
	Model          string      `json:"model"`
	ProductionYear int         `json:"productionYear"`
	Color          string      `json:"color"`
	IsAutomatic    bool        `json:"isAutomatic"`
	Stock          int         `json:"stock"`
	SalePrice      int         `json:"salePrice"`
	Status         string      `json:"status"`
}
