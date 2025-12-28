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
