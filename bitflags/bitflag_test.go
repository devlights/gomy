package bitflags_test

import (
	"testing"

	"github.com/devlights/gomy/bitflags"
)

type _Flag int

const (
	Bit1 _Flag = 0b0000_0001
	Bit2 _Flag = 0b0000_0010
	Bit4 _Flag = 0b0000_0100
	Bit8 _Flag = 0b0000_1000
)

func TestHas(t *testing.T) {
	// Arrange
	sut := Bit2

	// Act
	result1 := bitflags.Has(sut, Bit1)
	result2 := bitflags.Has(sut, Bit2)

	// Assert
	if result1 {
		t.Error("fail result1")
	}

	if !result2 {
		t.Error("fail result2")
	}
}

func TestSet(t *testing.T) {
	// Arrange
	sut := Bit1

	// Act
	bitflags.Set(&sut, Bit2, Bit4)

	// Assert
	if !bitflags.Has(sut, Bit2) {
		t.Error("sut does not have Bit2")
	}

	if !bitflags.Has(sut, Bit4) {
		t.Error("sut does not have Bit4")
	}
}

func TestSetEmpty(t *testing.T) {
	// Arrange
	sut := Bit1

	// Act
	bitflags.Set(&sut)

	// Assert
	if !bitflags.Has(sut, Bit1) {
		t.Error("sut does not have Bit2")
	}
}

func TestForce(t *testing.T) {
	// Arrange
	sut := Bit1

	// Act
	bitflags.Force(&sut, Bit2)

	// Assert
	if bitflags.Has(sut, Bit1) {
		t.Error("sut has Bit1")
	}

	if !bitflags.Has(sut, Bit2) {
		t.Error("sut does not have Bit2")
	}
}

func TestUnset(t *testing.T) {
	// Arrange
	sut := Bit2

	// Act
	bitflags.Unset(&sut, Bit2)

	// Assert
	if bitflags.Has(sut, Bit2) {
		t.Error("sut has Bit2")
	}
}
