package main

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want *BPMClient
	}{
		{
			name: "isabelle_wu",
			args: args{
				username: "isabelle_wu",
			},
			want: nil,
		},
		{
			name: "Unknown User",
			args: args{
				username: "unknown_user",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.username); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
