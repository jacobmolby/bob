package language

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestGoImportsSort(t *testing.T) {
	t.Parallel()

	a1 := []string{
		`"fmt"`,
		`"errors"`,
	}
	a2 := []string{
		`_ "github.com/lib/pq"`,
		`_ "github.com/gorilla/n"`,
		`"github.com/gorilla/mux"`,
		`"github.com/gorilla/websocket"`,
	}

	a1Expected := []string{`"errors"`, `"fmt"`}
	a2Expected := []string{
		`"github.com/gorilla/mux"`,
		`_ "github.com/gorilla/n"`,
		`"github.com/gorilla/websocket"`,
		`_ "github.com/lib/pq"`,
	}

	slices.SortFunc(a1, goImportSorter)
	if !reflect.DeepEqual(a1, a1Expected) {
		t.Errorf("Expected a1 to match a1Expected, got: %v", a1)
	}

	for i, v := range a1 {
		if v != a1Expected[i] {
			t.Errorf("Expected a1[%d] to match a1Expected[%d]:\n%s\n%s\n", i, i, v, a1Expected[i])
		}
	}

	slices.SortFunc(a2, goImportSorter)
	if !reflect.DeepEqual(a2, a2Expected) {
		t.Errorf("Expected a2 to match a2expected, got: %v", a2)
	}

	for i, v := range a2 {
		if v != a2Expected[i] {
			t.Errorf("Expected a2[%d] to match a2Expected[%d]:\n%s\n%s\n", i, i, v, a1Expected[i])
		}
	}
}

func TestCombineStringSlices(t *testing.T) {
	t.Parallel()

	var a, b []string
	slice := combineStringSlices(a, b)
	if ln := len(slice); ln != 0 {
		t.Error("Len was wrong:", ln)
	}

	a = []string{"1", "2"}
	slice = combineStringSlices(a, b)
	if ln := len(slice); ln != 2 {
		t.Error("Len was wrong:", ln)
	} else if slice[0] != a[0] || slice[1] != a[1] {
		t.Errorf("Slice mismatch: %#v %#v", a, slice)
	}

	b = a
	a = nil
	slice = combineStringSlices(a, b)
	if ln := len(slice); ln != 2 {
		t.Error("Len was wrong:", ln)
	} else if slice[0] != b[0] || slice[1] != b[1] {
		t.Errorf("Slice mismatch: %#v %#v", b, slice)
	}

	a = b
	b = []string{"3", "4"}
	slice = combineStringSlices(a, b)
	if ln := len(slice); ln != 4 {
		t.Error("Len was wrong:", ln)
	} else if slice[0] != a[0] || slice[1] != a[1] || slice[2] != b[0] || slice[3] != b[1] {
		t.Errorf("Slice mismatch: %#v + %#v != #%v", a, b, slice)
	}
}

var testImportStringExpect = `import (
	"fmt"

	"github.com/friendsofgo/errors"
)`

func TestGoImportsFormat(t *testing.T) {
	t.Parallel()

	s := []string{
		`"github.com/friendsofgo/errors"`,
		`"fmt"`,
	}

	got := strings.TrimSpace(string(goOutputLanguage{}.formatGoImports(s, "", false)))
	if got != testImportStringExpect {
		t.Error("want:\n", testImportStringExpect, "\ngot:\n", got)
	}
}
