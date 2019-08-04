package main

import (
	"testing"

	"github.com/juntaki/pp"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	m, err := Parse("./measurements.go.test")
	require.NoError(t, err)
	pp.Println(m)
	t.Fatal(1)
}
