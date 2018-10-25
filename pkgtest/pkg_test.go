package pkgtest

import (
	"log"
	"os"
	"testing"
)

var tests = []struct {
	name    string
	key     int
	want    int
	wantErr bool
}{
	{
		name:    "case 1",
		key:     10,
		want:    20,
		wantErr: false,
	},
	{
		name:    "case 2",
		key:     5,
		want:    10,
		wantErr: false,
	},
}

func multiple2(n int, exp int) func(*testing.T) {
	return func(t *testing.T) {
		res := n * 2
		if res != exp {
			t.Errorf("Error got = %v, want = %v", res, exp)
		}
	}
}

func TestCase1(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, multiple2(test.key, test.want))
	}
}

func TestCase2(t *testing.T) {
	t.Errorf("Case2 Fail")
}

func TestMain(m *testing.M) {
	log.Print("Pre1")

	exitCode := m.Run()

	log.Print("After1")
	os.Exit(exitCode)
}
