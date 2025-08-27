package api

import (
	"tenbounce/model"
	"tenbounce/util"
	"testing"
)

func pointTypes(typeIDs ...string) []model.PointType {
	return util.Map(
		typeIDs,
		func(typeID string) model.PointType {
			return model.PointType{ID: model.PointTypeID(typeID)}
		},
	)
}

func Test_ValidPointTypeID(t *testing.T) {
	var battery = []struct {
		point           model.Point
		validPointTypes []model.PointType
		valid           bool
	}{
		{
			point:           model.Point{PointTypeID: model.PointTypeID("ID1")},
			validPointTypes: pointTypes("ID1"),
			valid:           true,
		},
		{
			point:           model.Point{PointTypeID: model.PointTypeID("ID1")},
			validPointTypes: pointTypes("ID1", "ID2"),
			valid:           true,
		},
		{
			point:           model.Point{PointTypeID: model.PointTypeID("ID1")},
			validPointTypes: pointTypes("ID2"),
			valid:           false,
		},
		{
			point:           model.Point{PointTypeID: model.PointTypeID("ID1")},
			validPointTypes: pointTypes("ID2", "ID3"),
			valid:           false,
		},
	}

	for _, tcase := range battery {
		if err := validPointTypeID(tcase.point, tcase.validPointTypes); (err == nil) != tcase.valid {
			t.Errorf("invalid point type: got %s, expected one of %v", tcase.point.PointTypeID, util.Map(tcase.validPointTypes, func(pt model.PointType) model.PointTypeID { return pt.ID }))
		}
	}
}
