package stat_test

import (
	"testing"

	"github.com/devlights/gomy/fileio/stat"
)

func TestUserStat(t *testing.T) {
	// r--
	{
		f, fi, err := tmpAndChmodAndStat(0400)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		var (
			s = stat.User(fi)
			r = s.CanRead()
			w = s.CanWrite()
			x = s.CanExecute()
		)

		if !(r && !w && !x) {
			t.Errorf("[want] true, false, false\t[got] %v, %v, %v\n", r, w, x)
		}
	}
	// -w-
	{
		f, fi, err := tmpAndChmodAndStat(0200)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		var (
			s = stat.User(fi)
			r = s.CanRead()
			w = s.CanWrite()
			x = s.CanExecute()
		)

		if !(!r && w && !x) {
			t.Errorf("[want] false, true, false\t[got] %v, %v, %v\n", r, w, x)
		}
	}

	// --x
	{
		f, fi, err := tmpAndChmodAndStat(0100)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		var (
			s = stat.User(fi)
			r = s.CanRead()
			w = s.CanWrite()
			x = s.CanExecute()
		)

		if !(!r && !w && x) {
			t.Errorf("[want] false, false, true\t[got] %v, %v, %v\n", r, w, x)
		}
	}
}

func TestGroupStat(t *testing.T) {
	// r--
	{
		f, fi, err := tmpAndChmodAndStat(0040)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		var (
			s = stat.Group(fi)
			r = s.CanRead()
			w = s.CanWrite()
			x = s.CanExecute()
		)

		if !(r && !w && !x) {
			t.Errorf("[want] true, false, false\t[got] %v, %v, %v\n", r, w, x)
		}
	}
	// -w-
	{
		f, fi, err := tmpAndChmodAndStat(0020)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		var (
			s = stat.Group(fi)
			r = s.CanRead()
			w = s.CanWrite()
			x = s.CanExecute()
		)

		if !(!r && w && !x) {
			t.Errorf("[want] false, true, false\t[got] %v, %v, %v\n", r, w, x)
		}
	}

	// --x
	{
		f, fi, err := tmpAndChmodAndStat(0010)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		var (
			s = stat.Group(fi)
			r = s.CanRead()
			w = s.CanWrite()
			x = s.CanExecute()
		)

		if !(!r && !w && x) {
			t.Errorf("[want] false, false, true\t[got] %v, %v, %v\n", r, w, x)
		}
	}
}

func TestOtherStat(t *testing.T) {
	// r--
	{
		f, fi, err := tmpAndChmodAndStat(0004)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		var (
			s = stat.Other(fi)
			r = s.CanRead()
			w = s.CanWrite()
			x = s.CanExecute()
		)

		if !(r && !w && !x) {
			t.Errorf("[want] true, false, false\t[got] %v, %v, %v\n", r, w, x)
		}
	}
	// -w-
	{
		f, fi, err := tmpAndChmodAndStat(0002)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		var (
			s = stat.Other(fi)
			r = s.CanRead()
			w = s.CanWrite()
			x = s.CanExecute()
		)

		if !(!r && w && !x) {
			t.Errorf("[want] false, true, false\t[got] %v, %v, %v\n", r, w, x)
		}
	}

	// --x
	{
		f, fi, err := tmpAndChmodAndStat(0001)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		var (
			s = stat.Other(fi)
			r = s.CanRead()
			w = s.CanWrite()
			x = s.CanExecute()
		)

		if !(!r && !w && x) {
			t.Errorf("[want] false, false, true\t[got] %v, %v, %v\n", r, w, x)
		}
	}
}
