package models

// Import all your models here

// RegisterModels returns a slice of all models to be migrated
func RegisterModels() []interface{} {
	return []interface{}{
		&List{},
	}
}
