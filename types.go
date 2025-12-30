package main

type Matrix struct {
	Rows int
	Cols int
	Data []float64
}

func NewMatrix(rows int, cols int) *Matrix {
	return &Matrix{
		Rows: rows,
		Cols: cols,
		Data: make([]float64, rows*cols),
	}
}

func (m *Matrix) At(i int, j int) float64 {
	return m.Data[i*m.Cols+j]
}

func (m *Matrix) Set(i int, j int, v float64) {
	m.Data[i*m.Cols+j] = v
}

func (m *Matrix) GetRows(i int, j int) *Matrix {
	if i < 0 || j >= m.Rows || i > j {
		panic("invalid row indices")
	}
	numRows := j - i + 1
	numCols := m.Cols
	data := make([]float64, numRows*numCols)

	for row := range numRows {
		copy(data[row*numCols:(row+1)*numCols], m.Data[(i+row)*numCols:(i+row+1)*numCols])
	}

	return &Matrix{
		Rows: numRows,
		Cols: numCols,
		Data: data,
	}
}
