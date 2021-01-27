package reverse

import (
	"testing"
)

func TestString(t *testing.T) {
	hello := "Hello World!"
	message, err := String([]byte(hello))

	if string(message) != "!dlroW olleH" || err != nil {
		t.Fatalf(`String("Hello World!") = %q, %v, want "!dlroW olleH", error`, message, err)
	}
}

func TestEmptyString(t *testing.T) {
	message, err := String([]byte(""))

	if string(message) != "" || err == nil {
		t.Fatalf(`String("") = %q, %v, want "", error`, message, err)
	}
}

func TestLinkedList(t *testing.T) {
	list, err := Init([]byte(string("Hello World!")))
	message, err := LinkedList(list)

	if string(Print(message)) != "!dlroW olleH" || err != nil {
		t.Fatalf(`LinkedList(Init([]byte(string("Hello World!")))) = %q, %v, want "!dlroW olleH", error`, Print(message), err)
	}
}

func TestEmptyLinkedList(t *testing.T) {
	list, err := Init([]byte{})

	if list != nil || err == nil {
		t.Fatalf(`Init("") = %v, %v, want "<nil>", error`, list, err)
	}
}
