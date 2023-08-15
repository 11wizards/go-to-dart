package multipackage

import "github.com/11wizards/go-to-dart/examples/multipackage/shared"

type Outer struct {
	ID   *shared.DomainID `firestore:"id"`
	Name string           `firestore:"name"`
}
