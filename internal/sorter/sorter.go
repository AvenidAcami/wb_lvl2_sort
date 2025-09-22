package sorter

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Options struct {
	SortColumn uint32
	Nuneric    bool
	Reverse    bool
	Unique     bool
}

func ReadLines() ([]string, error) {
	var lines []string
	var line string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		lines = append(lines, strings.TrimRight(line, "\r"))
	}

	return lines, scanner.Err()
}

func Sort(lines []string, options Options) []string {
	if options.Unique {
		lines = getUniqueLines(lines)
	}

	// Логика сортировки
	sort.SliceStable(lines, func(i, j int) bool {
		var less bool

		columnI := getColumn(lines[i], options.SortColumn)
		columnJ := getColumn(lines[j], options.SortColumn)

		if options.Nuneric {
			ni, errI := strconv.ParseFloat(columnI, 64)
			nj, errJ := strconv.ParseFloat(columnJ, 64)

			if (errI == nil) && (errJ == nil) {
				less = ni < nj
			} else {
				less = columnI < columnJ
			}
		} else {
			less = columnI < columnJ
		}

		if options.Reverse {
			return !less
		}

		return less
	})

	return lines
}

func getUniqueLines(lines []string) []string {
	checked := make(map[string]struct{})
	res := make([]string, 0)

	for _, val := range lines {
		if _, ok := checked[val]; !ok {
			checked[val] = struct{}{}
			res = append(res, val)
		}
	}

	return res
}

func getColumn(line string, ind uint32) string {
	columns := strings.Split(line, "\t")
	if ind == 0 {
		return line
	}
	if ind > uint32(len(columns)) {
		return ""
	}
	return columns[ind-1]
}
