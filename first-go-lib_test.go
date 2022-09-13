/**
 * Author: Alex Aloshchenko
 * File: first-go-lib_test.go
 */

package lib

import (
	"math"
	"testing"
)

func Test_Add(t *testing.T) {
	tests := map[string]struct {
		a      int
		b      int
		exp    int
		expErr error
	}{
		"success":      {a: 3, b: 5, exp: 8, expErr: nil},
		"int overflow": {a: math.MaxInt, b: 5, exp: 0, expErr: ErrOverflow},
	}
	for name, data := range tests {
		got, err := Add(data.a, data.b)
		if err != nil && err != data.expErr {
			t.Errorf("[%s] unexpected error. expected:\"%s\", got:\"%s\"", name, data.expErr, err)
		}
		if got != data.exp || err != data.expErr {
			t.Errorf("[%s] expected: %v, %v; got %v, %v", name, data.exp, data.expErr, got, err)
		}
	}
}

func Test_Subtract(t *testing.T) {
	tests := map[string]struct {
		a      int
		b      int
		exp    int
		expErr error
	}{
		"success":      {a: 5, b: 3, exp: 2, expErr: nil},
		"int overflow": {a: -5, b: math.MaxInt, exp: 0, expErr: ErrOverflow},
	}
	for name, data := range tests {
		got, err := Subtract(data.a, data.b)
		if err != nil && err != data.expErr {
			t.Errorf("[%s] unexpected error. expected:\"%s\", got:\"%s\"", name, data.expErr, err)
		}
		if got != data.exp || err != data.expErr {
			t.Errorf("[%s] expected: %v, %v; got %v, %v", name, data.exp, data.expErr, got, err)
		}
	}
}

func Test_Multiply(t *testing.T) {
	tests := map[string]struct {
		a      int
		b      int
		exp    int
		expErr error
	}{
		"success":      {a: 5, b: 3, exp: 15, expErr: nil},
		"int overflow": {a: 2, b: math.MaxInt, exp: 0, expErr: ErrOverflow},
	}
	for name, data := range tests {
		got, err := Multiply(data.a, data.b)
		if err != nil && err != data.expErr {
			t.Errorf("[%s] unexpected error. expected:\"%s\", got:\"%s\"", name, data.expErr, err)
		}
		if got != data.exp || err != data.expErr {
			t.Errorf("[%s] expected: %v, %v; got %v, %v", name, data.exp, data.expErr, got, err)
		}
	}
}

func Test_Division(t *testing.T) {
	a := 10
	b := 2
	expected := a / b

	if got := Division(a, b); got != expected {
		t.Errorf("Add(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}
