package main

import "fmt"

type Logistic struct {
	inputs  *Matrix
	weights *Matrix
	bias    []float64
}

func NewLogisticRegressor(rows int, cols int, classes int) *Logistic {
	return &Logistic{
		inputs:  NewMatrix(rows, cols),
		weights: NewMatrix(cols, classes),
		bias:    make([]float64, classes),
	}
}

func (l *Logistic) predict(X *Matrix) *Matrix {
	n := X.Rows
	m := X.Cols
	c := len(l.bias)
	res := NewMatrix(n, c)

	for i := range n {
		for j := range c {
			t := l.bias[j]
			for k := range m {
				t += X.At(i, k) * l.weights.At(k, j)
			}
			res.Set(i, j, t)
		}
		start := i * c
		end := (i + 1) * c
		copy(res.Data[start:end], softmax(res.Data[start:end]))
	}

	return res
}

func (l *Logistic) fit(X *Matrix, Y []int, lr float64, batchSize int, epochs int) {
	l.inputs = X
	dataSize := l.inputs.Rows
	numClasses := len(l.bias)
	features := X.Cols
	for epoch := 1; epoch <= epochs; epoch++ {
		var epochLoss float64
		var numBatches int

		for start := 0; start < dataSize; start += batchSize {
			end := min(start+batchSize, dataSize)
			batchX := X.GetRows(start, end) // [bs, m]
			batchY := Y[start:end]          // [bs]
			currentBatchSize := end - start

			p := l.predict(batchX)
			zGrads := NewMatrix(currentBatchSize, numClasses)

			for i := range currentBatchSize {
				for j := range numClasses {
					zgrad := p.At(i, j)
					if j == batchY[i] {
						zgrad -= 1
					}
					zGrads.Set(i, j, zgrad)
				}
			}

			for i := range features{
				for j := range numClasses{
					var grad float64
					for k := range currentBatchSize{
						grad += zGrads.At(k, j) * batchX.At(k, i)
					}
					grad /= float64(currentBatchSize)
					w := l.weights.At(i, j)
					l.weights.Set( i, j, w - grad * lr )
				}
			}

			for i := range numClasses{
				var grad float64
				for k := range currentBatchSize{
					grad += zGrads.At(k, i)
				}
				grad /= float64(currentBatchSize)
				l.bias[i] -= grad * lr
			}

			epochLoss += SCCEL(p, batchY)
			numBatches++
		}

		epochLoss /= float64(numBatches)
		if epoch%100 == 0 {
			fmt.Printf("Epoch %v: Loss %.4f\n", epoch, epochLoss)
		}
	}
}
