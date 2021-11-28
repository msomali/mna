package mna_test

import (
	"github.com/techcraftlabs/mna"
	"github.com/techcraftlabs/mna/rand"
	"strings"
	"testing"
)

func FilterByPrefix(prefix string) mna.FilterPhoneFunc {
	return func(phone string) bool {
		return strings.HasPrefix(phone, prefix)
	}
}

func FilterBySubstring(substr string) mna.FilterPhoneFunc {
	return func(phone string) bool {
		return strings.Contains(phone, substr)
	}
}

func FilterBySuffix(suffix string) mna.FilterPhoneFunc {
	return func(phone string) bool {
		return strings.HasSuffix(phone, suffix)
	}
}

func OperatorsListFilter(ops ...mna.Operator) mna.FilterOperatorFunc {
	return func(op mna.Operator) bool {
		for _, operator := range ops {
			if op == operator {
				return true
			}
		}

		return false
	}
}


func TestGet(t *testing.T) {
	type args struct {
		phoneNumber string
	}
	tests := []struct {
		name    string
		args    args
		want    mna.Operator
		wantErr bool
	}{
		{
			name:    "test vodacom number",
			args:    args{
				phoneNumber: "0765999999",
			},
			want:    mna.Vodacom,
			wantErr: false,
		},
		{
			name:    "test tigo number",
			args:    args{
				phoneNumber: "0712999999",
			},
			want:    mna.Tigo,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mna.Get(tt.args.phoneNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetThenFilter(t *testing.T) {
	type args struct {
		phoneNumber string
		f2           mna.FilterOperatorFunc
		f1          mna.FilterPhoneFunc
	}
	tests := []struct {
		name    string
		args    args
		want    mna.Operator
		wantErr bool
	}{
		{
			name:    "test filter with suffix and pass tigo and vodacom numbers only",
			args:    args{
				phoneNumber: "0712915799",
				f2:           OperatorsListFilter(mna.Tigo, mna.Vodacom),
				f1:          FilterBySuffix("799"),
			},
			want:    mna.Tigo,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mna.GetAndFilter(tt.args.phoneNumber, tt.args.f1, tt.args.f2)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAndFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAndFilter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomAndFilters(t *testing.T) {

	type args struct {
		len int
		suffix string
		ops  []mna.Operator
	}

	f1 := func (suffix string)mna.FilterPhoneFunc{
		return func(phone string) bool {
            return strings.HasSuffix(phone, suffix)
        }
    }

	f2 := func(ops...mna.Operator)mna.FilterOperatorFunc{
        return func(op mna.Operator) bool {
            for _, operator := range ops {
                if op == operator {
                    return true
                }
            }

            return false
        }
    }

	//numbers := rand.GenerateNWithFilters(1000,f1("99"),f2(mna.Tigo,mna.Vodacom))

    tests := []struct {
        name    string
        args    args

    }{
        {
            name:    "randomly generate 100 numbers that ends with 799 and are either Tigo/Voda/Airtel",
            args:    args{
				len: 100,
                suffix: "799",
                ops: []mna.Operator{mna.Tigo, mna.Vodacom,mna.Airtel},
            },
        },
		{
			name: "generate 10 halotel number that has 300 as suffix",
			args: args{
				len:    10,
				suffix: "300",
				ops:    []mna.Operator{mna.Halotel},
			},
		},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
           got := rand.GenerateNWithFilters(tt.args.len,f1(tt.args.suffix),f2(tt.args.ops...))

		   if len(got) != tt.args.len {
			   t.Errorf("GenerateNWithFilters() got = %v, want %v", len(got), tt.args.len)
           }

			for i := 0; i < len(got); i++ {
				t.Logf("%d: %s\n", i, got[i])
			}
        })

    }
}