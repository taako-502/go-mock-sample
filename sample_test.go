package sample

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	tests := []struct {
		name    string
		want    *user
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sample()
			if (err != nil) != tt.wantErr {
				t.Errorf("Sample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sample() = %v, want %v", got, tt.want)
			}
		})
	}
}
