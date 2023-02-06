package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPasswordHash(t *testing.T) {

	pass := "awiroot123"

	s, err := HashPassword(pass)
	assert.NoError(t, err)

	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should match",
			args: args{
				password: "awiroot123",
				hash:     s,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
