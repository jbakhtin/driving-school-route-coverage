package models

import (
	"github.com/twpayne/go-geom/encoding/geojson"
	"github.com/twpayne/go-geom/encoding/wkb"
)

type LineString []byte

func (ls LineString) MarshalJSON() ([]byte, error) {
	test, _ := wkb.Unmarshal(ls)
	json, _ := geojson.Marshal(test)
	return json, nil
}

type Route struct {
	ID int64
	LineString LineString
	CreatedAt string
	UpdatedAt *string
}
