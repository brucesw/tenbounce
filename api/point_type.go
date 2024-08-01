package api

import "tenbounce/model"

// FUTURE STORE METHODS

// TODO(bruce): use store
func GetPointTypesFromDB() ([]model.PointType, error) {
	var pointTypes = []model.PointType{
		{
			ID:   "4e4b2b1c-5063-425a-a409-71b431068f78",
			Name: "Compulsory Routine",
		},
		{
			ID:   "0d1b30ef-00d4-41d6-8581-b8d554752816",
			Name: "Optional Routine",
		},
		{
			ID:   "dade4383-d869-4562-a680-88cb38f9972a",
			Name: "Tenbounce",
		},
		{
			ID:   "8640f8e9-0cf6-4be4-b182-d40c21a44067",
			Name: "Ten Doubles",
		},
	}

	return pointTypes, nil
}
