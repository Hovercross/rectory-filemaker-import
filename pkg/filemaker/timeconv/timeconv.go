package timeconv

import "time"

// A TimeFunc gives back a time.Time object
type TimeFunc func(string) (time.Time, error)

// A StringFunc gives back a string object, such as for JSON
type StringFunc func(string) (string, error)

// A Parser manages the various format types
type Parser struct {
	registeredFormats map[string]dateTimeFormater
}

// RegisterFormat will register a new format for parsing
func (p *Parser) RegisterFormat(filemaker, reference string, withDate bool, withTime bool) {
	if p.registeredFormats == nil {
		p.registeredFormats = map[string]dateTimeFormater{}
	}

	dtf := dateTimeFormater{
		referenceTime: reference,
	}

	if withDate && withTime {
		dtf.formatTime = time.RFC3339
	} else if withDate {
		dtf.formatTime = "2006-01-02"
	} else if withTime {
		dtf.formatTime = "15:04:05"
	}

	p.registeredFormats[filemaker] = dtf
}

// Defaults will pop in some default parsers
func (p *Parser) Defaults() {
	p.RegisterFormat("M/d/yyyy", "1/2/2006", true, false)
}

type dateTimeFormater struct {
	referenceTime string
	formatTime    string
}

// Gives back a function that gives back a string for marshaling
func (dtf dateTimeFormater) String() StringFunc {
	timeParser := dtf.Time()

	return func(s string) (string, error) {
		t, err := timeParser(s)

		if err != nil {
			return "", err
		}

		return t.Format(dtf.formatTime), nil
	}
}

// Gives back a function that gives back a time
func (dtf dateTimeFormater) Time() TimeFunc {
	return func(s string) (time.Time, error) {
		return time.Parse(dtf.referenceTime, s)
	}
}

// String gives back the strong parser for a given format
func (p *Parser) String(format string) (StringFunc, bool) {
	f, found := p.registeredFormats[format]

	if found {
		return f.String(), true
	}

	return nil, false
}

// Time gives back the time parser for a given format
func (p *Parser) Time(format string) (TimeFunc, bool) {
	f, found := p.registeredFormats[format]

	if found {
		return f.Time(), true
	}

	return nil, false
}
