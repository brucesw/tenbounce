package api

import (
	"tenbounce/model"
	"tenbounce/util"
	"testing"
)

func Test_ValidPointTypeID(t *testing.T) {
	var battery = []struct {
		point           model.Point
		validPointTypes []model.PointType
		valid           bool
	}{
		{
			point:           model.Point{PointTypeID: model.PointTypeID("ID1")},
			validPointTypes: []model.PointType{{ID: model.PointTypeID("ID1")}},
			valid:           true,
		},
		{
			point:           model.Point{PointTypeID: model.PointTypeID("ID1")},
			validPointTypes: []model.PointType{{ID: model.PointTypeID("ID1")}, {ID: model.PointTypeID("ID2")}},
			valid:           true,
		},
		{
			point:           model.Point{PointTypeID: model.PointTypeID("ID1")},
			validPointTypes: []model.PointType{{ID: model.PointTypeID("ID2")}},
			valid:           false,
		},
		{
			point:           model.Point{PointTypeID: model.PointTypeID("ID1")},
			validPointTypes: []model.PointType{{ID: model.PointTypeID("ID2")}, {ID: model.PointTypeID("ID3")}},
			valid:           false,
		},
	}

	for _, tcase := range battery {
		if err := validPointTypeID(tcase.point, tcase.validPointTypes); (err == nil) != tcase.valid {
			t.Errorf("invalid point type: got %s, expected one of %v", tcase.point.PointTypeID, util.Map(tcase.validPointTypes, func(pt model.PointType) model.PointTypeID { return pt.ID }))
		}
	}
}
