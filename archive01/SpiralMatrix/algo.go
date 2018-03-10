package algo

// https://leetcode.com/problems/spiral-matrix/description/
func spiralOrder(matrix [][]int) []int {
	m := NewMatrix(matrix)
	result := make([]int, 0, m.Size())

	for {
		if v, ok := m.Get(); ok {
			result = append(result, v)
		} else {
			break
		}
	}

	return result
}

type Matrix struct {
	matrix [][]int
	flags  [][]bool
	dir    byte

	m, n int
	i, j int
}

const (
	dirRight = iota
	dirDown
	dirLeft
	dirUp
)

func NewMatrix(matrix [][]int) *Matrix {
	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}

	flags := make([][]bool, m)
	for i := range flags {
		flags[i] = make([]bool, n)
	}
	mm := Matrix{
		matrix: matrix,
		flags:  flags,
		dir:    dirRight,
		m:      m,
		n:      n,

		i: 0,
		j: -1,
	}
	return &mm
}

func (m *Matrix) Get() (int, bool) {
	switch m.dir {
	case dirRight:
		m.j++
		if m.j >= m.n || m.flags[m.i][m.j] {
			m.dir = dirDown
			m.j--
			m.i++
		}
	case dirDown:
		m.i++
		if m.i >= m.m || m.flags[m.i][m.j] {
			m.dir = dirLeft
			m.i--
			m.j--
		}
	case dirLeft:
		m.j--
		if m.j < 0 || m.flags[m.i][m.j] {
			m.dir = dirUp
			m.j++
			m.i--
		}
	case dirUp:
		m.i--
		if m.i < 0 || m.flags[m.i][m.j] {
			m.dir = dirRight
			m.i++
			m.j++
		}
	}

	if m.i < 0 || m.i >= m.m || m.j < 0 || m.j >= m.n || m.flags[m.i][m.j] {
		return 0, false
	}

	m.flags[m.i][m.j] = true
	return m.matrix[m.i][m.j], true
}

func (m *Matrix) Size() int {
	return m.m * m.n
}
