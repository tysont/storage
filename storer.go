package storage

import (
	"fmt"
)

type Storer interface {
	
	Get(class string, key string) fmt.Stringer
	Put(fmt.Stringer)
}