package main

import "testing"

func TestSum(t *testing.T) {
    total := sum(5, 6)
    if total != 11 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

