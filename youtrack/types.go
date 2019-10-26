package youtrack

import "time"

type enumBundle struct {
	Values []EnumType `json:"values"`
	_      string     `json:"$type"`
}

type EnumType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Ordinal     int    `json:"ordinal"`
	Description string `json:"description"`
	_           string `json:"$type"`
}

type versionBundle struct {
	Values []VersionType `json:"values"`
	_      string        `json:"$type"`
}

type VersionType struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Ordinal            int    `json:"ordinal"`
	Description        string `json:"description"`
	Released           bool   `json:"released"`
	ReleaseDateNumeric int64  `json:"releaseDate"`
	ReleaseDate        time.Time
	_                  string `json:"$type"`
}

type Issue struct {
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	ID          string   `json:"id"`
	IsResolved  bool     `json:"is_resolved"`
	Priority    string   `json:"priority"`
	Stage       string   `json:"stage"`
	FixVersion  string   `json:"fix_version"`
	Type        string   `json:"type"`
	Subsystems  []string `json:"subsystems"`
}

type projectCustomFieldDefinition struct {
	_    string `json:"$type"`
	Name string `json:"name"`
}

type projectCustomField struct {
	_     string                       `json:"$type"`
	Field projectCustomFieldDefinition `json:"field"`
}

type customField struct {
	_                  string             `json:"$type"`
	_                  string             `json:"id"`
	ProjectCustomField projectCustomField `json:"projectCustomField"`
	Value              interface{}        `json:"value"`
}
