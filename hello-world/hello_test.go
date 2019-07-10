package helloworld

import "testing"

func TestHello(t *testing.T) {

  assertCorrectMessage := func(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
      t.Errorf("got '%s' want '%s'", got, want)
    }
  }

  t.Run("saying hello to people", func(t *testing.T) {
    name := "Gede"
    got := Hello(name, "english")
    want := "Hello, " + name
    assertCorrectMessage(t, got, want)
  })

  t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
    got := Hello("", "")
    want := "Hello, World"
    assertCorrectMessage(t, got, want)
  })

  t.Run("say in Spanish", func(t *testing.T) {
    got := Hello("Gede", "spanish")
    want := "Hola, Gede"
    assertCorrectMessage(t, got, want)
  })

  t.Run("say in French", func(t *testing.T) {
    got := Hello("Gede", "french")
    want := "Bonjour, Gede"
    assertCorrectMessage(t, got, want)
  })
}
