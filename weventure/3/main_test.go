package main

import "testing"

//
//func TestHello(t *testing.T) {
//	t.Run("saying hello to people", func(t *testing.T){
//		got := Hello("Toan")
//		want := "Hello, Toan"
//		if got != want {
//			t.Errorf("got %q want %q", got,want)
//		}
//	})
//	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
//		got := Hello("")
//		want := "Hello, world"
//
//		if got != want {
//			t.Errorf("got %q want %q", got, want)
//		}
//	})
//}

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Toan", "")
		want := "Hello, Toan"
		assertCorrectMessage(t,got,want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t,got,want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Truong", "Spanish")
		want := "Hola, Truong"
		assertCorrectMessage(t,got,want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Ha", "French")
		want := "Bonjour, Ha"
		assertCorrectMessage(t,got,want)
	})
}