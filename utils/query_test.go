package utils

import (
	"testing"
)

func TestWithOrderBy(t *testing.T) {
	type args struct {
		mainQuery string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "adding query args",
			args: args{
				mainQuery: "SELECT * from chapters c WHERE c.manga_id = ?",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WithOrderBy(tt.args.mainQuery)
			s := got("c.id", "desc")
			tt.args.mainQuery = s

			f := WithLimit(tt.args.mainQuery)
			s2 := f("1", "2")
			tt.args.mainQuery = s2

			t.Log(tt.args.mainQuery)

		})
	}
}
