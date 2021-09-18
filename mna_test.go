package mna

import (
	"reflect"
	"testing"
)

func TestFormat(t *testing.T) {
	type args struct {
		phoneNumber string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test with plus",
			args: args{
				phoneNumber: "+255765992153",
			},
			want:    "255765992153",
			wantErr: false,
		},
		{
			name: "test with just 255",
			args: args{
				phoneNumber: "255765992153",
			},
			want:    "255765992153",
			wantErr: false,
		},
		{
			name: "test with zero",
			args: args{
				phoneNumber: "0765992153",
			},
			want:    "255765992153",
			wantErr: false,
		},
		{
			name: "test without zero",
			args: args{
				phoneNumber: "765992153",
			},
			want:    "255765992153",
			wantErr: false,
		},
		{
			name: "test with spaces",
			args: args{
				phoneNumber: "0765 992 153",
			},
			want:    "255765992153",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Format(tt.args.phoneNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetails(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name    string
		args    args
		want    Data
		wantErr bool
	}{
		{
			name:    "test ridiculous number",
			args:    args{
				phone: "+255 765 992 153",
			},
			want:    Data{
				OperatorName: "Vodacom Tanzania PLC",
				CommonName:   Vodacom,
				Status:       statusOperational,
				Prefixes:     vodaPrefixes,
			},
			wantErr: false,
		},
		{
			name:    "test ridiculous number",
			args:    args{
				phone: "765 992 153",
			},
			want:    Data{
				OperatorName: "Vodacom Tanzania PLC",
				CommonName:   Vodacom,
				Status:       statusOperational,
				Prefixes:     vodaPrefixes,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Details(tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("Details() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Details() got = %v, want %v", got, tt.want)
			}
		})
	}
}