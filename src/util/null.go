package util

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (x *NullString) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.String)
}

type NullInt64 struct {
	sql.NullInt64
}

func (x *NullInt64) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Int64)
}
