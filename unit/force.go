// Code generated by "go generate gonum.org/v1/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

// Force represents a force in Newtons.
type Force float64

const (
	Yottanewton Force = 1e24
	Zettanewton Force = 1e21
	Exanewton   Force = 1e18
	Petanewton  Force = 1e15
	Teranewton  Force = 1e12
	Giganewton  Force = 1e9
	Meganewton  Force = 1e6
	Kilonewton  Force = 1e3
	Hectonewton Force = 1e2
	Decanewton  Force = 1e1
	Newton      Force = 1.0
	Decinewton  Force = 1e-1
	Centinewton Force = 1e-2
	Millinewton Force = 1e-3
	Micronewton Force = 1e-6
	Nanonewton  Force = 1e-9
	Piconewton  Force = 1e-12
	Femtonewton Force = 1e-15
	Attonewton  Force = 1e-18
	Zeptonewton Force = 1e-21
	Yoctonewton Force = 1e-24
)

// Unit converts the Force to a *Unit
func (f Force) Unit() *Unit {
	return New(float64(f), Dimensions{
		LengthDim: 1,
		MassDim:   1,
		TimeDim:   -2,
	})
}

// Force allows Force to implement a Forcer interface
func (f Force) Force() Force {
	return f
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (f *Force) From(u Uniter) error {
	if !DimensionsMatch(u, Newton) {
		*f = Force(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*f = Force(u.Unit().Value())
	return nil
}

func (f Force) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", f, float64(f))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " N"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(f))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(f))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(f))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(f))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g N)", c, f, float64(f))
	}
}
