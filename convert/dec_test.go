package convert

import (
	"testing"
)

func TestDec2Hex(t *testing.T) {
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
		{name: "not decimal", args: args{val: "abc", prefix: "", length: 0}, want: "", errShouldRaised: true},
		{name: "15 to F", args: args{val: "15", prefix: "", length: 0}, want: "F", errShouldRaised: false},
		{name: "15 to 0xF", args: args{val: "15", prefix: "0x", length: 0}, want: "0xF", errShouldRaised: false},
		{name: "777 to 0x309", args: args{val: "777", prefix: "0x", length: 0}, want: "0x309", errShouldRaised: false},
		{name: "777 to 0x0309", args: args{val: "777", prefix: "0x", length: 4}, want: "0x0309", errShouldRaised: false},
		{name: "777 to 0x00000309", args: args{val: "777", prefix: "0x", length: 8}, want: "0x00000309", errShouldRaised: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Dec2Hex(tt.args.val, tt.args.prefix, tt.args.length)

			if got != tt.want {
				t.Errorf("Dec2Hex() = %v, want %v", got, tt.want)
			}

			if tt.errShouldRaised {
				if err == nil {
					t.Errorf("Dec2Hex() should raise err but err is nil")
				}
			} else {
				if err != nil {
					t.Errorf("Dec2Hex() err = %v", err)
				}
			}
		})
	}
}

func TestDec2Bin(t *testing.T) {
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
		{name: "not decimal", args: args{val: "abc", prefix: "", length: 0}, want: "", errShouldRaised: true},
		{name: "15 to F", args: args{val: "15", prefix: "", length: 0}, want: "1111", errShouldRaised: false},
		{name: "15 to 0b1111", args: args{val: "15", prefix: "0b", length: 0}, want: "0b1111", errShouldRaised: false},
		{name: "112 to 0b1110000", args: args{val: "112", prefix: "0b", length: 0}, want: "0b1110000", errShouldRaised: false},
		{name: "112 to 0b01110000", args: args{val: "112", prefix: "0b", length: 8}, want: "0b01110000", errShouldRaised: false},
		{name: "112 to 0b000001110000", args: args{val: "112", prefix: "0b", length: 12}, want: "0b000001110000", errShouldRaised: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Dec2Bin(tt.args.val, tt.args.prefix, tt.args.length)

			if got != tt.want {
				t.Errorf("Dec2Bin() = %v, want %v", got, tt.want)
			}

			if tt.errShouldRaised {
				if err == nil {
					t.Errorf("Dec2Bin() should raise err but err is nil")
				}
			} else {
				if err != nil {
					t.Errorf("Dec2Bin() err = %v", err)
				}
			}
		})
	}
}
