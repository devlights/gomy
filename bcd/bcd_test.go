package bcd

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestToBcd(t *testing.T) {
	type (
		value struct {
			num   uint64
			count int
		}
		testin struct {
			data value
		}
		testout struct {
			result []byte
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{data: value{
				num:   2019,
				count: 2,
			}},
			out: testout{result: []byte{
				0x20, 0x19,
			}},
		},
	}

	for _, c := range cases {
		bcd := ToBcd(c.in.data.num, c.in.data.count)

		t.Logf("[bcd] %v", hex.Dump(bcd))

		if !bytes.Equal(bcd, c.out.result) {
			t.Errorf("[want] %v\t[got] %v", c.out.result, bcd)
		}
	}
}

func TestToUInt64(t *testing.T) {
	type (
		testin struct {
			data []byte
		}
		testout struct {
			result uint64
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{data: []byte{
				0x20,
				0x19,
			}},
			out: testout{result: 2019},
		},
	}

	for _, c := range cases {
		num := ToUInt64(c.in.data)

		t.Logf("[uint64] %v", num)

		if num != c.out.result {
			t.Errorf("[want] %v\t[got] %v", c.out.result, num)
		}
	}
}
