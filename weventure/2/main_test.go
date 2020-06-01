package main

import "testing"

//func TestHello(t *testing.T){
//	got := Hello("Chris")
//	want := "Hello, " + name
//
//	if got != want {
//		t.Errorf("got %q want %q", got,want)
//	}
//}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Toan")
		want := "Hello, Toan"
		if got != want {
			t.Errorf("got %q want %q", got,want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}