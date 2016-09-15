package main

import (
	"fmt"
)

type Feet float64
type Meters float64

func FeetToMeters(f Feet) Meters { return Meters(f/3.28084) }
func MetersToFeet(m Meters) Feet { return Feet(m*3.28084)}

func (f Feet) String() string    { return fmt.Sprintf("%g ft", f) }
func (m Meters) String() string  { return fmt.Sprintf("%g m", m) }
