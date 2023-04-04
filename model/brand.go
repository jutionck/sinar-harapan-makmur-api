package model

type Brand struct {
	BaseModel
	Name     string    `json:"name"`
	Vehicles []Vehicle `json:"vehicles,omitempty"`
}

func (Brand) TableName() string {
	return "mst_brand"
}
