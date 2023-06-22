package entity

type Asset struct {
	ID          string
	Name        string
	MarketValue int
}

func NewAsset(id string, name string, marketValue int) *Asset {
	return &Asset{
		ID:          id,
		Name:        name,
		MarketValue: marketValue,
	}
}
