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
		callback func(t *testing.T, client *cli.Client)
		expected []string
	}{
		{
			name: "keys",
			db: &fakeDB{
				keys: func(n string) ([]string, error) {
					return []string{"a", "b"}, nil
				},
			},
			callback: func(t *testing.T, client *cli.Client) {
				client.Keys()
			},
			expected: []string{"a", "b"},
		},
	}

	for _, tc := range testCases {
		old := os.Stdout // keep backup of the real stdout
		f, _ := ioutil.TempFile("", "")
		os.Stdout = f

		t.Run(tc.name, func(subt *testing.T) {
			client := cli.New(tc.db)

			client.Create("testdb")
			tc.callback(subt, client)

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
		return fmt.Errorf("lengths don't match, %d, %d", len(a), len(b))
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

	isDB   func(dbName string) (string, error)
	create func(dbname string) error
	remove func(dbname string) error
	keys   func(dbname string) ([]string, error)
	set    func(dbname string, key string, value string) error
	get    func(dbname string, key string) (string, error)
	del    func(dbname string, key string) error
	list   func() ([]string, error)
	time   func() string
}

// since fakeDB embeds kval.DB you only need to implement
// the methods that you actually test.

func (f *fakeDB) Create(dbName string) error {
	return nil
}

func (f *fakeDB) Keys(dbname string) ([]string, error) {
	return f.keys(dbname)
}
