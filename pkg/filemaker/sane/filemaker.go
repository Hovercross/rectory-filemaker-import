package sane

// Data represents the full set of on disk data
type Data struct {
	ErrorCode int64
	Product   *Product
	Database  *Database
	Metadata  *Metadata
	ResultSet *ResultSet
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
	ModID    int64
	RecordID string
	Cols     []*Col
}

// Col entity
type Col struct {
	Field *Field
	Data  []string
}
