package cli_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/fwmiller/kval/internal/cli"
	"github.com/fwmiller/kval/internal/kval"
)

func TestClient(t *testing.T) {
	testCases := []struct {
		name     string
		db       *fakeDB
		callback func(t *testing.T, client *cli.Client, db *fakeDB)
		expected []string
	}{
		{
			name: "keys",
			db: &fakeDB{
				db: map[string]string{"a": "1", "b": "2"},
			},
			callback: func(t *testing.T, client *cli.Client, db *fakeDB) {
				client.Keys()
			},
			expected: []string{"a", "b"},
		},
		{
			name: "get",
			db:   &fakeDB{db: map[string]string{"c": "3"}},
			callback: func(t *testing.T, client *cli.Client, db *fakeDB) {
				client.Get("c")
			},
			expected: []string{"3"},
		},
	}

	for _, tc := range testCases {
		old := os.Stdout // keep backup of the real stdout
		f, _ := ioutil.TempFile("", "")
		os.Stdout = f

		t.Run(tc.name, func(subt *testing.T) {
			client := cli.New(tc.db)

			client.Create("testdb")
			tc.callback(subt, client, tc.db)

			os.Stdout = old
			f.Seek(0, 0)

			data, err := ioutil.ReadAll(f)
			if err != nil {
				subt.Fatalf("unexpected error reading tempfile: %s", err)
			}

			actual := strings.Split(strings.TrimSpace(string(data)), "\n")
			if err := equalStrings(actual, tc.expected); err != nil {
				subt.Error(err)
			}

			f.Close()
		})

	}
}

func equalStrings(a, b []string) error {
	if len(a) != len(b) {
		return fmt.Errorf("lengths of %v and %v don't match, %d, %d", a, b, len(a), len(b))
	}

	for i, val := range a {
		if val != b[i] {
			return fmt.Errorf("%v does not equal %v", a, b)
		}
	}

	return nil
}

type fakeDB struct {
	kval.DB

	db map[string]string
}

// since fakeDB embeds kval.DB you only need to implement
// the methods that you actually test.

func (f *fakeDB) Create(dbName string) error {
	return nil
}

func (f *fakeDB) Keys(dbname string) ([]string, error) {
	out := make([]string, len(f.db))
	var i int
	for k := range f.db {
		out[i] = k
		i++
	}

	return out, nil
}

func (f *fakeDB) Get(dbname, key string) (string, error) {
	return f.db[key], nil
}
