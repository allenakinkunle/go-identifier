package identifier

import (
	"github.com/gofrs/uuid"
	gonanoid "github.com/matoous/go-nanoid"
)

type IIdentifier interface {
	NewNanoID() string
	NewUUIDv5(name string) string
}

type Identifier struct {
	characters string
	namespace  string
}

func NewIdentifier(characters string, namespace string) *Identifier {
	return &Identifier{
		characters: characters,
		namespace:  namespace,
	}
}

func (i Identifier) NewNanoID() string {
	id, _ := gonanoid.Generate(i.characters, 21)
	return id
}

func (i Identifier) NewUUIDv5(name string) string {
	return uuid.NewV5(uuid.FromStringOrNil(i.namespace), name).String()
}
