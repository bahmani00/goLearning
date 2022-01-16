package main

import "testing"

func TestHello(t * testing.T)  {
  t.Run("Saying hello to people", func (t * testing.T)  {
    got := Hello("Chris", "")
    want := "Hello, Chris"

    if got != want {
      t.Errorf("got %q wanted %q", got, want)
    }
  })

  t.Run("Saying hello to nobody", func (t * testing.T)  {
    got := Hello("", "")
    want := "Hello, World"

    if got != want {
      t.Errorf("got %q wanted %q", got, want)
    }
  })

  t.Run("Saying hello to people in Spanish", func (t * testing.T)  {
    got := Hello("Chris", "sp")
    want := "Hola, Chris"

    if got != want {
      t.Errorf("got %q wanted %q", got, want)
    }
  })  
}