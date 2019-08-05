package main

import (
	"errors"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

type OP int

const (
	InvalidFieldError    = errors.New("Invalid field")
	EQ                OP = iota
	NEQ
	GT
	GTE
	LT
	LTE
)

func StrToWhere(field string, op OP, value string) ([]qm.QueryMod, error) {
	switch field {
  {{range $i, $c := .Cases}}
  {{$c}}
  {{end}}
	}

	return nil, InvalidFieldError
}
