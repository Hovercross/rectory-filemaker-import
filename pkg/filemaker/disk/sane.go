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

	if out.Product, err = f.Product.Sane(out); err != nil {
		return
	}

	if out.Database, err = f.Database.Sane(out); err != nil {
		return
	}

	if out.Metadata, err = f.Metadata.Sane(out); err != nil {
		return
	}

	if out.ResultSet, err = f.ResultSet.Sane(out); err != nil {
		return
	}

	return
}

// Sane the product
func (p Product) Sane(parent *sane.Data) (out *sane.Product, err error) {
	out = &sane.Product{}

	out.Build = p.Build
	out.Name = p.Name
	out.Version = p.Version

	return
}

// Sane the database
func (d Database) Sane(parent *sane.Data) (out *sane.Database, err error) {
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
func (m Metadata) Sane(parent *sane.Data) (out *sane.Metadata, err error) {
	out = &sane.Metadata{}

	out.Fields = make([]*sane.Field, len(m.Fileds))

	for i, field := range m.Fileds {
		if out.Fields[i], err = field.Sane(out); err != nil {
			out = nil
			return
		}
	}

	return
}

// Sane the result set
func (rs ResultSet) Sane(parent *sane.Data) (out *sane.ResultSet, err error) {
	return
}

// Sane gives back the sane version of the field
func (f Field) Sane(parent *sane.Metadata) (out *sane.Field, err error) {
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
