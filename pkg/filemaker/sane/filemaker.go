package sane

import (
	"encoding/json"

	"github.com/hovercross/rectory-filemaker-import/pkg/filemaker/timeconv"
)

// Data represents the full set of on disk data
type Data struct {
	ErrorCode int64
	Product   *Product
	Database  *Database
	Metadata  *Metadata
	ResultSet *ResultSet

	DateStringParser timeconv.StringFunc
	TimeStringParser timeconv.StringFunc
}

// Product entity
type Product struct {
	Build   string
	Name    string
	Version string
}

// Database entity
type Database struct {
	DateFormat string
	Name       string
	TimeFormat string
	Records    int64
	// Layout     string
}

// Metadata entity
type Metadata struct {
	Fields []*Field
}

// Field entity
type Field struct {
	EmptyOK   bool
	MaxRepeat int64
	Name      string
	Type      string
}

// ResultSet entity
type ResultSet struct {
	Rows  []*Row
	Found int64
}

// Row entity
type Row struct {
	ModID    string
	RecordID string
	Cols     []*Col
}

// Col entity
type Col struct {
	Field *Field
	Data  []string

	parent *Data
}

// RegisterParent lets marshal magic happen
func (c *Col) RegisterParent(d *Data) {
	c.parent = d
}

// MarshalJSON will handle dates and times at least
func (c *Col) MarshalJSON() ([]byte, error) {
	if c.Field.MaxRepeat == 1 && c.Field.Type == "DATE" {
		s, err := c.parent.DateStringParser(c.Data[0])

		if err != nil {
			return nil, err
		}

		return []byte(s), nil
	}

	if c.Field.MaxRepeat == 1 && c.Field.Type == "TIME" {
		s, err := c.parent.TimeStringParser(c.Data[0])

		if err != nil {
			return nil, err
		}

		return []byte(s), nil
	}

	if c.Field.MaxRepeat == 1 {
		return json.Marshal(c.Data[0])
	}

	return json.Marshal(c.Data)
}
