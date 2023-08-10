/* package lineutil surfaces some utilites for dealing with line ranges
and reading ranges of lines from files

Copyright © 2023 Charles <Asciifaceman> Corbett
*/

package lineutil

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ParseLineRanges produces a slice of int slice representing ranges
// of lines to read from a list of range strings like "12-15"
// [0] = start, [1] = end
func ParseLineRanges(lines []string) ([][]int, error) {
	var pairs [][]int
	for _, line := range lines {
		var set []int

		pair := strings.Split(line, "-")
		if len(pair) < 2 {
			pair = append(pair, pair[0])
		}

		start, err := strconv.Atoi(pair[0])
		if err != nil {
			return nil, err
		}

		end, err := strconv.Atoi(pair[1])
		if err != nil {
			return nil, err
		}

		if start > end {
			return nil, fmt.Errorf("invalid range, end comes before start: [%d-%d]", start, end)
		}

		set = append(set, start)
		set = append(set, end)

		pairs = append(pairs, set)

	}
	return pairs, nil
}

// ReadLineFromFile returns a specific line from the given file
//
// currently unused bet kept for the moment
func ReadLineFromFile(filename string, lineNum int) (string, error) {
	r, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	sc := bufio.NewScanner(r)
	increment := 0
	for sc.Scan() {
		increment++
		if increment == lineNum {
			return sc.Text(), nil
		}
	}
	return "", io.EOF
}

// ReadLineRangeFromFile returns lines from a given []int range
func ReadLineRangeFromFile(filename string, lineRange []int) ([]string, error) {

	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	if len(lineRange) > 2 {
		return nil, fmt.Errorf("provided range defines more than start and endpoint, expected length of 2: %v", lineRange)
	}
	var lines []string

	sc := bufio.NewScanner(r)
	increment := 0
	for sc.Scan() {
		increment++

		if increment >= lineRange[0] && increment <= lineRange[1] {
			lines = append(lines, sc.Text())
		}

		if increment >= lineRange[1] {
			return lines, nil
		}

	}
	if len(lines) > 0 {
		fmt.Fprintln(os.Stderr, "[WARN] EOF reached but did capture lines")
		return lines, nil
	}

	return nil, io.EOF
}
