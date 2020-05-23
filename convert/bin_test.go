package convert

import (
	"testing"
)

func TestBin2Dec(t *testing.T) {
	type args struct {
		val    string
		prefix string
		length int
	}
	tests := []struct {
		name            string
		args            args
		want            string
		errShouldRaised bool
	}{
		{name: "empty", args: args{val: "", prefix: "", length: 0}, want: "", errShouldRaised: false},
		{name: "not decimal", args: args{val: "xyz", prefix: "", length: 0}, want: "", errShouldRaised: true},
		{name: "1000 to 8", args: args{val: "1000", prefix: "", length: 0}, want: "8", errShouldRaised: false},
		{name: "1000 to dec8", args: args{val: "1000", prefix: "dec", length: 0}, want: "dec8", errShouldRaised: false},
		{name: "1110000 to 112", args: args{val: "1110000", prefix: "", length: 0}, want: "112", errShouldRaised: false},
		{name: "1110000 to 0112", args: args{val: "1110000", prefix: "", length: 4}, want: "0112", errShouldRaised: false},
		{name: "1110000 to 00000112", args: args{val: "1110000", prefix: "", length: 8}, want: "00000112", errShouldRaised: false},
		{name: "0b1110000 to 00000112", args: args{val: "0b1110000", prefix: "", length: 8}, want: "00000112", errShouldRaised: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Bin2Dec(tt.args.val, tt.args.prefix, tt.args.length)

			if got != tt.want {
				t.Errorf("Bin2Dec() = %v, want %v", got, tt.want)
			}

			if tt.errShouldRaised {
				if err == nil {
					t.Errorf("Bin2Dec() should raise err but err is nil")
				}
			} else {
				if err != nil {
					t.Errorf("Bin2Dec() err = %v", err)
				}
			}
		})
	}
}

func TestBin2Hex(t *testing.T) {
	type args struct {
		val    string
		prefix string
		length int
	}
	tests := []struct {
		name            string
		args            args
		want            string
		errShouldRaised bool
	}{
		{name: "empty", args: args{val: "", prefix: "", length: 0}, want: "", errShouldRaised: false},
		{name: "not decimal", args: args{val: "xyz", prefix: "", length: 0}, want: "", errShouldRaised: true},
		{name: "1000 to 8", args: args{val: "1000", prefix: "", length: 0}, want: "8", errShouldRaised: false},
		{name: "1000 to 0x8", args: args{val: "1000", prefix: "0x", length: 0}, want: "0x8", errShouldRaised: false},
		{name: "1110000 to 0x70", args: args{val: "1110000", prefix: "0x", length: 0}, want: "0x70", errShouldRaised: false},
		{name: "1110000 to 0x0070", args: args{val: "1110000", prefix: "0x", length: 4}, want: "0x0070", errShouldRaised: false},
		{name: "1110000 to 0x00000070", args: args{val: "1110000", prefix: "0x", length: 8}, want: "0x00000070", errShouldRaised: false},
		{name: "0b1110000 to 0x00000070", args: args{val: "0b1110000", prefix: "0x", length: 8}, want: "0x00000070", errShouldRaised: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Bin2Hex(tt.args.val, tt.args.prefix, tt.args.length)

			if got != tt.want {
				t.Errorf("Bin2Hex() = %v, want %v", got, tt.want)
			}

			if tt.errShouldRaised {
				if err == nil {
					t.Errorf("Bin2Hex() should raise err but err is nil")
				}
			} else {
				if err != nil {
					t.Errorf("Bin2Hex() err = %v", err)
				}
			}
		})
	}
}
