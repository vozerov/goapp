package main

import "testing"

func TestSum(t *testing.T) {
    total := sum(5, 51)
    if total != 56 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 56)
    }
}

