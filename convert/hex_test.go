package convert

import (
	"testing"
)

func TestHex2Dec(t *testing.T) {
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
		{name: "F to 15", args: args{val: "F", prefix: "", length: 0}, want: "15", errShouldRaised: false},
		{name: "F to dec15", args: args{val: "F", prefix: "dec", length: 0}, want: "dec15", errShouldRaised: false},
		{name: "0x309 to 777", args: args{val: "0x309", prefix: "", length: 0}, want: "777", errShouldRaised: false},
		{name: "0x309 to 0777", args: args{val: "0x309", prefix: "", length: 4}, want: "0777", errShouldRaised: false},
		{name: "0x309 to 00000777", args: args{val: "0x309", prefix: "", length: 8}, want: "00000777", errShouldRaised: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hex2Dec(tt.args.val, tt.args.prefix, tt.args.length)

			if got != tt.want {
				t.Errorf("Hex2Dec() = %v, want %v", got, tt.want)
			}

			if tt.errShouldRaised {
				if err == nil {
					t.Errorf("Hex2Dec() should raise err but err is nil")
				}
			} else {
				if err != nil {
					t.Errorf("Hex2Dec() err = %v", err)
				}
			}
		})
	}
}

func TestHex2Bin(t *testing.T) {
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
		{name: "8 to 1000", args: args{val: "8", prefix: "", length: 0}, want: "1000", errShouldRaised: false},
		{name: "8 to 0b1000", args: args{val: "8", prefix: "0b", length: 0}, want: "0b1000", errShouldRaised: false},
		{name: "0x70 to 0b1110000", args: args{val: "0x70", prefix: "0b", length: 0}, want: "0b1110000", errShouldRaised: false},
		{name: "0x70 to 0b01110000", args: args{val: "0x70", prefix: "0b", length: 8}, want: "0b01110000", errShouldRaised: false},
		{name: "0x70 to 0b000001110000", args: args{val: "0x70", prefix: "0b", length: 12}, want: "0b000001110000", errShouldRaised: false},
		{name: "length==-1", args: args{val: "0x0F", prefix: "", length: -1}, want: "00001111", errShouldRaised: false},
		{name: "int32.max", args: args{val: "0x7FFFFFFF", prefix: "", length: -1}, want: "01111111111111111111111111111111", errShouldRaised: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hex2Bin(tt.args.val, tt.args.prefix, tt.args.length)

			if got != tt.want {
				t.Errorf("Hex2Bin() = %v, want %v", got, tt.want)
			}

			if tt.errShouldRaised {
				if err == nil {
					t.Errorf("Hex2Bin() should raise err but err is nil")
				}
			} else {
				if err != nil {
					t.Errorf("Hex2Bin() err = %v", err)
				}
			}
		})
	}
}
