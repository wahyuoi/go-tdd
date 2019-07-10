package integers

import "testing"

func TestAdder(t *testing.T) {

  assertCorrectMessage := func(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
      t.Errorf("got '%d' want '%d'", got, want)
    }
  }

  sum := Add(2, 2)
  expected := 4

  assertCorrectMessage(t, sum, expected)

}
