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
	length     uint
	namespace  string
}

func NewIdentifier(characters string, length uint, namespace string) (*Identifier, error) {
	return &Identifier{
		characters: characters,
		length:     length,
		namespace:  namespace,
	}, nil
}

func (i Identifier) NewNanoID() string {
	id, _ := gonanoid.Generate(i.characters, int(i.length))
	return id
}

func (i Identifier) NewUUIDv5(name string) string {
	return uuid.NewV5(uuid.FromStringOrNil(i.namespace), name).String()
}
