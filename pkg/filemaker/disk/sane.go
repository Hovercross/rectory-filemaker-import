package disk

import (
	"fmt"
	"strconv"

	"github.com/hovercross/rectory-filemaker-import/pkg/filemaker/sane"
)

// Sane the entire file
func (f *File) Sane() (out *sane.Data, err error) {
	out = &sane.Data{}

	out.ErrorCode = f.ErrorCode

	if out.Product, err = f.Product.Sane(); err != nil {
		return
	}

	if out.Database, err = f.Database.Sane(); err != nil {
		return
	}

	if out.Metadata, err = f.Metadata.Sane(); err != nil {
		return
	}

	if out.ResultSet, err = f.ResultSet.Sane(out.Metadata.Fields); err != nil {
		return
	}

	return
}

// Sane the product
func (p Product) Sane() (out *sane.Product, err error) {
	out = &sane.Product{}

	out.Build = p.Build
	out.Name = p.Name
	out.Version = p.Version

	return
}

// Sane the database
func (d Database) Sane() (out *sane.Database, err error) {
	out = &sane.Database{}

	out.DateFormat = d.DateFormat
	out.Name = d.Name
	out.TimeFormat = d.TimeFormat

	if out.Records, err = strconv.ParseInt(d.Records, 10, 64); err != nil {
		out = nil
		return
	}

	return
}

// Sane the metadata
func (m Metadata) Sane() (out *sane.Metadata, err error) {
	out = &sane.Metadata{}

	out.Fields = make([]*sane.Field, len(m.Fileds))

	for i, field := range m.Fileds {
		if out.Fields[i], err = field.Sane(); err != nil {
			out = nil
			return
		}
	}

	return
}

// Sane the result set
func (rs ResultSet) Sane(fields []*sane.Field) (out *sane.ResultSet, err error) {
	out = &sane.ResultSet{}

	if out.Found, err = strconv.ParseInt(rs.Found, 10, 64); err != nil {
		out = nil
		return
	}

	out.Rows = make([]*sane.Row, len(rs.Rows))

	for i, row := range rs.Rows {
		out.Rows[i] = row.Sane(fields)
	}

	return
}

// Sane gives back the sane version of the field
func (f Field) Sane() (out *sane.Field, err error) {
	out = &sane.Field{}

	switch f.EmptyOK {
	case "YES":
		out.EmptyOK = true
	case "NO":
		out.EmptyOK = false
	default:
		out = nil
		err = fmt.Errorf("Could not understand EMPTYOK: %s", f.EmptyOK)
		return
	}

	if out.MaxRepeat, err = strconv.ParseInt(f.MaxRepeat, 10, 64); err != nil {
		err = fmt.Errorf("Could not parse MAXREPEAT as an int: %s; %v", f.MaxRepeat, err)
		return
	}

	out.Name = f.Name
	out.Type = f.Type

	return
}

// Sane the row
func (r Row) Sane(fields []*sane.Field) *sane.Row {
	out := &sane.Row{}

	out.ModID = r.ModID
	out.RecordID = r.RecordID

	out.Cols = make([]*sane.Col, len(r.Cols))

	for i, col := range r.Cols {
		field := fields[i]
		out.Cols[i] = col.Sane(field)
	}

	return out
}

// Sane the column
func (c Col) Sane(field *sane.Field) *sane.Col {
	out := &sane.Col{
		Field: field,
		Data:  make([]string, len(c.Data)),
	}

	copy(out.Data, c.Data)
	return out
}
