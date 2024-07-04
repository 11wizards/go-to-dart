package generics

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
