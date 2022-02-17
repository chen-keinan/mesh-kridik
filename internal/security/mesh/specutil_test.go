package mesh

import (
	"fmt"
	"testing"
)

func Test_meshUtil(t *testing.T) {
	is, err := LoadIstioSpecs()
	if err != nil {
		t.Error(err)
	}
	if len(is) != 24 {
		t.Error(fmt.Sprintf("Test_meshUtil Want %d Got %d", len(is), 32))
	}
}
