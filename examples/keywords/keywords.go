package keywords

import "time"

// Class name that is a Dart keyword
type Class struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Struct with field names that are Dart keywords
type TestKeywords struct {
	ID        int       `json:"id"` // Not a keyword
	Class     string    `json:"class"`
	Static    string    `json:"static"`
	Final     string    `json:"final"`
	Const     string    `json:"const"`
	Var       string    `json:"var"`
	If        string    `json:"if"`
	Else      string    `json:"else"`
	For       string    `json:"for"`
	While     string    `json:"while"`
	Switch    string    `json:"switch"`
	Case      string    `json:"case"`
	Default   string    `json:"default"`
	Break     string    `json:"break"`
	Continue  string    `json:"continue"`
	Return    string    `json:"return"`
	Try       string    `json:"try"`
	Catch     string    `json:"catch"`
	Finally   string    `json:"finally"`
	Throw     string    `json:"throw"`
	New       string    `json:"new"`
	This      string    `json:"this"`
	Super     string    `json:"super"`
	Null      string    `json:"null"`
	True      string    `json:"true"`
	False     string    `json:"false"`
	Async     string    `json:"async"`
	Await     string    `json:"await"`
	Yield     string    `json:"yield"`
	Abstract  string    `json:"abstract"`
	Extends   string    `json:"extends"`
	With      string    `json:"with"`
	Mixin     string    `json:"mixin"`
	Enum      string    `json:"enum"`
	CreatedAt time.Time `json:"createdAt"` // Not a keyword
}

// Another class with a keyword name
type Interface struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}
