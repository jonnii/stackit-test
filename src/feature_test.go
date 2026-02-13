package src

import "testing"

func TestFeature(t *testing.T) {
    if got := Feature(); got != "hello from feature" {
        t.Errorf("got %q, want %q", got, "hello from feature")
    }
}
