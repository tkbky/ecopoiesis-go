package tf

import (
	"fmt"
	"testing"
)

func TestTf(t *testing.T) {
	cases := []struct {
		args []string
		want string
	}{
		{[]string{"kind", "name", "obj"}, "resource  \"kind\" \"name\" \"obj\""},
		{[]string{"foo", "bar", "baz"}, "resource  \"foo\" \"bar\" \"baz\""},
	}

	for _, tc := range cases {
		desc := fmt.Sprintf("Resource{Kind: \"%s\", Name: \"%s\", Obj: \"%s\"}", tc.args[0], tc.args[1], tc.args[2])
		rsrc := Resource{Kind: tc.args[0], Name: tc.args[1], Obj: tc.args[2]}
		got := rsrc.Tf()

		if got != tc.want {
			t.Errorf("%s=%q, want=%q", desc, got, tc.want)
		}
	}
}
