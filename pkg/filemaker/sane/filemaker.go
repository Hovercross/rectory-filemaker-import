package sane

import (
	"fmt"
	"reflect"

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

// ImposeOnto will push the vaues onto a struct by tag
func (c *Col) ImposeOnto(v interface{}) error {
	t := reflect.TypeOf(v)

	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("value must be a pointer")
	}

	elem := reflect.ValueOf(v).Elem()
	elemType := elem.Type()

	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		fmTag := field.Tag.Get("filemaker")

		if fmTag != "" && fmTag == c.Field.Name {
			c.impose(elem, i)
		}
	}

	return nil
}

func (c *Col) impose(elem reflect.Value, i int) error {
	// Now, we want to set it

	field := elem.Type().Field(i)

	if field.Type.Kind() == reflect.String && c.Field.Type == "TEXT" && c.Field.MaxRepeat == 1 {
		elem.Field(i).SetString(c.Data[0])
		return nil
	}

	return fmt.Errorf("Didn't know what to do with it")
}
