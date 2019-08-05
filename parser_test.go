package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	_, err := Parse("./measurements.go.test")
	require.NoError(t, err)
	// pp.Println(m)
	t.Fatal(1)
}
