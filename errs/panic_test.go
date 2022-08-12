package errs

import (
	"errors"
	"testing"
)

func TestPanic(t *testing.T) {
	//
	// Arrange
	//
	noNil := func() (bool, error) {
		return true, nil
	}
	notNil := func() (bool, error) {
		return false, errors.New("this is test")
	}

	//
	// Act
	//
	func() {
		defer func() {
			//
			// Assert
			//
			err := recover()
			if err != nil {
				t.Error("[want] not panic\t[got] panic")
			}
		}()
		Panic(noNil())
	}()

	func() {
		defer func() {
			//
			// Assert
			//
			err := recover()
			if err == nil {
				t.Error("[want] panic\t[got] not panic")
			}
		}()
		Panic(notNil())
	}()
}
