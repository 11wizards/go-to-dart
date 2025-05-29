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
	Assert    string    `json:"assert"`
	Break     string    `json:"break"`
	Case      string    `json:"case"`
	Catch     string    `json:"catch"`
	Class     string    `json:"class"`
	Const     string    `json:"const"`
	Continue  string    `json:"continue"`
	Default   string    `json:"default"`
	Do        string    `json:"do"`
	Else      string    `json:"else"`
	Enum      string    `json:"enum"`
	Extends   string    `json:"extends"`
	False     string    `json:"false"`
	Final     string    `json:"final"`
	Finally   string    `json:"finally"`
	For       string    `json:"for"`
	If        string    `json:"if"`
	In        string    `json:"in"`
	Is        string    `json:"is"`
	New       string    `json:"new"`
	Null      string    `json:"null"`
	Rethrow   string    `json:"rethrow"`
	Return    string    `json:"return"`
	Super     string    `json:"super"`
	Switch    string    `json:"switch"`
	This      string    `json:"this"`
	Throw     string    `json:"throw"`
	True      string    `json:"true"`
	Try       string    `json:"try"`
	Var       string    `json:"var"`
	Void      string    `json:"void"`
	While     string    `json:"while"`
	With      string    `json:"with"`
	Yield     string    `json:"yield"`
	CreatedAt time.Time `json:"createdAt"` // Not a keyword
}

// Another class with a keyword name
type Interface struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}
