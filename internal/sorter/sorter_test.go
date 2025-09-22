package sorter

import (
	"os"
	"strings"
	"testing"
)

type testCase struct {
	name      string
	inputFile string
	wantFile  string
	opts      Options
}

func TestSortFromFiles(t *testing.T) {
	tests := []testCase{
		{
			name:      "simple",
			inputFile: "test_cases/simple_test_input.txt",
			wantFile:  "test_cases/simple_test_output.txt",
			opts:      Options{SortColumn: 1},
		},
		{
			name:      "bu column",
			inputFile: "test_cases/by_column_input.txt",
			wantFile:  "test_cases/by_column_output.txt",
			opts:      Options{SortColumn: 2},
		},
		{
			name:      "numeric",
			inputFile: "test_cases/numeric_input.txt",
			wantFile:  "test_cases/numeric_output.txt",
			opts:      Options{SortColumn: 2, Nuneric: true},
		},
		{
			name:      "reverse",
			inputFile: "test_cases/reverse_input.txt",
			wantFile:  "test_cases/reverse_output.txt",
			opts:      Options{Reverse: true},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inBytes, err := os.ReadFile(tc.inputFile)
			if err != nil {
				t.Fatalf("can't read input: %v", err)
			}
			lines := strings.Split(strings.TrimSpace(string(inBytes)), "\n")

			got := Sort(lines, tc.opts)

			wantBytes, err := os.ReadFile(tc.wantFile)
			if err != nil {
				t.Fatalf("can't read want: %v", err)
			}
			wantLines := strings.Split(strings.TrimSpace(string(wantBytes)), "\n")

			if len(got) != len(wantLines) {
				t.Fatalf("lines count mismatch: got %d, want %d", len(got), len(wantLines))
			}
			for i := range wantLines {
				if got[i] != wantLines[i] {
					t.Errorf("[%d] got %q, want %q", i, got[i], wantLines[i])
				}
			}
		})
	}
}
