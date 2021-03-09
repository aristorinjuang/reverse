package reverse

import (
	"testing"
)

func TestString(t *testing.T) {
	message, err := String([]byte("Hello World!"))

	if string(message) != "!dlroW olleH" || err != nil {
		t.Fatalf(`String([]byte("Hello World!")) = %q, %v, want "!dlroW olleH", error`, message, err)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := String([]byte("Hello World!"))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestEmptyString(t *testing.T) {
	message, err := String([]byte(""))

	if string(message) != "" || err == nil {
		t.Fatalf(`String([]byte("")) = %q, %v, want "", error`, message, err)
	}
}

func TestLinkedList(t *testing.T) {
	list, _ := Init([]byte("Hello World!"))
	message, err := LinkedList(list)

	if string(Print(message)) != "!dlroW olleH" || err != nil {
		t.Fatalf(`LinkedList(Init([]byte("Hello World!"))) = %q, %v, want "!dlroW olleH", error`, Print(message), err)
	}
}

func BenchmarkLinkedList(b *testing.B) {
	list, _ := Init([]byte("Hello World!"))

	for i := 0; i < b.N; i++ {
		_, err := LinkedList(list)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestEmptyLinkedList(t *testing.T) {
	list, err := Init([]byte(""))

	if list != nil || err == nil {
		t.Fatalf(`Init([]byte("")) = %v, %v, want "<nil>", error`, list, err)
	}
}
