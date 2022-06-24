package idgen

import (
	"github.com/google/uuid"
)

func New(count int) string {
	id := uuid.New()
	var idx string = id.String()
	idx = idx[:count]
	return idx
}
