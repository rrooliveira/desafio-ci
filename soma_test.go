package main

import "testing"

func TestSoma(t *testing.T) {
    total := soma(5, 5)
    if total != 10 {
       t.Errorf("O resultado está incorreto, o obtido foi: %d, e o esperado é: %d.", total, 10)
    }
}