package sane

import (
	"encoding/json"
	"fmt"

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

// ToInterface will parse the thing into a generic version, acceptabe for processing
func (c *Col) ToInterface() (interface{}, error) {
	// Non-repeated dates become the ISO8601 date, since Go doesn't have a plain date type
	if c.Field.MaxRepeat == 1 && c.Field.Type == "DATE" {
		s, err := c.parent.DateStringParser(c.Data[0])

		if err != nil {
			return nil, err
		}

		return s, nil
	}

	// Repeated dates become an array of ISO8601 dates
	if c.Field.Type == "DATE" {
		out := make([]string, len(c.Data))

		for i, rawValue := range c.Data {
			s, err := c.parent.DateStringParser(rawValue)

			if err != nil {
				return nil, err
			}

			out[i] = s
		}

		return out, nil
	}

	// Non-repeated text becomes a straight string
	if c.Field.MaxRepeat == 1 && c.Field.Type == "TEXT" {
		return c.Data[0], nil
	}

	// Repeated text is as is
	if c.Field.Type == "TEXT" {
		return c.Data, nil
	}

	return nil, fmt.Errorf("Unknown field type: %s", c.Field.Type)
}

// ToMap converts a row into a map[string]interface{} for generic handling
func (r *Row) ToMap() (map[string]interface{}, error) {
	out := map[string]interface{}{}

	for _, c := range r.Cols {
		val, err := c.ToInterface()

		if err != nil {
			return nil, err
		}

		out[c.Field.Name] = val
	}

	out["_recordID"] = r.RecordID
	out["_modID"] = r.ModID

	return out, nil
}
