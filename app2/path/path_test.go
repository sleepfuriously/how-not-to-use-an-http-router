package path_test

import (
	"fmt"
	"testing"

	"github.com/sleepfuriously/how-not-to-use-an-http-router/app2/path"
)

func TestPath(t *testing.T) {
	check(t, "")
	check(t, "/")
	check(t, "//")
	check(t, "user", "user")
	check(t, "/user", "user")
	check(t, "/user/", "user")
	check(t, "//user", "user")
	check(t, "user//", "user")
	check(t, "user/1", "user", "1")
	check(t, "user/1/", "user", "1")
	check(t, "/user/1/", "user", "1")
	check(t, "user/1//2", "user", "1", "2")
	check(t, "user/1//2/", "user", "1", "2")
}

func check(t *testing.T, pathString string, expect ...string) {
	iter := path.NewIterator(pathString)
	var got []string
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		got = append(got, next)
	}
	if len(got) != len(expect) {
		t.Fatal(fmt.Sprintf("pathString: %s, got: %v, expect: %v", pathString, got, expect))
	}
	for i := range expect {
		if got[i] != expect[i] {
			t.Fatal(fmt.Sprintf("pathString: %s, got: %v, expect: %v", pathString, got, expect))
		}
	}
}
