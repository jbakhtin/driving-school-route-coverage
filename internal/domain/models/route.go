package models

import (
	"github.com/twpayne/go-geom/encoding/geojson"
	"github.com/twpayne/go-geom/encoding/wkb"
)

// Geometry TODO: вынести в типы
type Geometry []byte

func (ls Geometry) MarshalJSON() ([]byte, error) {
	test, _ := wkb.Unmarshal(ls)
	json, _ := geojson.Marshal(test)
	return json, nil
}

type Route struct {
	ID         int64    `json:"id,omitempty"`
	UserID     int64    `json:"user_id,omitempty"`
	Name       string   `json:"name,omitempty"`
	LineString Geometry `json:"line_string,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty"`
	UpdatedAt  *string  `json:"updated_at,omitempty"`
}
