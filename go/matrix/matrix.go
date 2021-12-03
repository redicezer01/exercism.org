// package matrix helps you work with matrix
package matrix

import "strings"
import "strconv"
import _ "fmt"
import "errors"

// Matrix type store matrix and implements functions to work with cols and rows
type Matrix struct {
	rows [][]int
}

// New parse string to Matrix type
func New(s string) (*Matrix, error) {
	table, err := parseMatrix(s)
	if err != nil {
		return nil, err
	}
	matrix := &Matrix{
		rows: table,
	}

	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.

// Cols returns columns of the matrix
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, 0, 10)
	for i, v := range m.rows {
		for colN, item := range v {
			if i == 0 {
				cols = append(cols, make([]int, len(m.rows)))
			}
			cols[colN][i] = item
		}
	}
	return cols
}

// Rows return rows
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, len(m.rows))
	for i, v := range m.rows {
		rows[i] = make([]int, len(v))
		copy(rows[i], v)
	}
	return rows
}

// Set cells value
func (m *Matrix) Set(row, col, val int) bool {
	rowCnt := len(m.rows)
	if rowCnt-1 < row || row < 0 {
		return false
	}

	colCnt := len(m.rows[row])
	if colCnt-1 < col || col < 0 {
		return false
	}

	m.rows[row][col] = val
	return true
}

// parse string to slice of rows
func parseMatrix(s string) ([][]int, error) {
	//s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	table := make([][]int, 0)
	colCount := 0
	rowCount := 0

	for _, line := range lines {
		if line == "" {
			return nil, errors.New("incorrect input.")
		}

		line = strings.TrimSpace(line)
		row := make([]int, 0, len(line))
		splitedLine := strings.Split(line, " ")

		for _, v := range splitedLine {
			number, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				return nil, err
			}
			row = append(row, number)
		}

		if rowCount == 0 || colCount == 0 {
			colCount = len(row)
		} else if len(row) != colCount {
			return nil, errors.New("not a matrix.")
		}

		table = append(table, row)
		rowCount++
	}
	return table, nil
}
