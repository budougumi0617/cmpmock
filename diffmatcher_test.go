package cmpmock

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_DiffEq(t *testing.T) {
	type Foo struct{ CreateAt time.Time }
	t1 := time.Now()
	t2 := t1.Add(100 * time.Millisecond)
	defaultString := "diff(-got +want) is %s"
	type args struct {
		want interface{}
		opts cmp.Options
	}
	tests := []struct {
		name      string
		args      args
		x         Foo
		wantMatch bool
		wantDiff  string
	}{
		{
			name: "diff less than a second with default option",
			args: args{
				want: Foo{CreateAt: t1},
				opts: nil,
			},
			x:         Foo{CreateAt: t2},
			wantMatch: true,
			wantDiff:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := DiffEq(tt.args.want, tt.args.opts...)
			x := Foo{CreateAt: t2}
			if got := sut.Matches(x); got != tt.wantMatch {
				t.Errorf("Matches() = %v, want %v", got, tt.wantMatch)
			}
			if got := sut.String(); got != fmt.Sprintf(defaultString, tt.wantDiff) {
				t.Errorf("String() = %q, want %q", got, tt.wantDiff)
			}
		})
	}
}
