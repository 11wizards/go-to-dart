package prefix

import (
	"github.com/11wizards/go-to-dart/examples/prefix/shared"
)

type KeyValuePair[TKey, TValue any] struct {
	Key   TKey   `json:"key"`
	Value TValue `json:"value"`
}

type Map[TKey, TValue any] struct {
	Items []KeyValuePair[TKey, TValue] `json:"items"`
}

type Instance struct {
	M Map[string, int] `json:"m"`
}

type UserRepository struct {
	Users []shared.User `json:"users"`
}
