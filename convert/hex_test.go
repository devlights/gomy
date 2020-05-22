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
