package main

import "testing"

func TestHello(t *testing.T) {

    assertCorrectMessgae := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    }

    t.Run("Saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"
        assertCorrectMessgae(t, got, want)
    })

    t.Run("empty string defaults to 'World'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, world"
        assertCorrectMessgae(t, got, want)
    })
}
