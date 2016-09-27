package storage_test

import (

	"testing"
	"github.com/tysont/storage"
)

func TestPutGet(t *testing.T) {

	object := new(TestObject)
	object.Payload = "foo"

	storer := new(storage.MemoryStorer)
	storer.Put(object)

	otherObject := storer.Get("testobject", object.String())
	if otherObject == nil || otherObject.String() != object.String() {

		t.Errorf("Stored object couldn't be retrieved, or didn't match expected object.")
	}
}

type TestObject struct { 

	Payload string
}

func (this *TestObject) String() string {

	return this.Payload
}