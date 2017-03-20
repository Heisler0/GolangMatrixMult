package matrix

type Matrix2d struct{
        Rows int
        Columns int
        Matrix []float64
}

func (a Matrix2d) Mult(b Matrix2d) Matrix2d{

        rows := a.Rows
        cols := b.Columns

        results := make([]float64, rows * cols)

        ch := make(chan float64, rows * cols)

        length := a.Columns
        var col []float64
        for i := 0; i < length; i++{
                col = a.getColumn(i)
                for j := 0; j < length; j++{
                        go innerprod(col, b.Matrix[j*cols:j*cols+cols], ch)
                }
        }
        for i := 0; i < length; i++{
                for j := 0; j < length; j++{
                        results[i*cols + j] = <-ch
                }
        }

        return Matrix2d{rows, cols, results}
}

func innerprod(c []float64, r []float64, ch chan float64){
        result := 0.0
        for i := 0; i < len(c); i++{
                result += c[i] * r[i]
        }
        ch <- result
}

func (m Matrix2d) getColumn(col int) []float64{
        result := make([]float64, m.Rows)
        cols := m.Columns
        for i :=0; i< m.Rows; i++{
                result[i] = m.Matrix[i*cols + col]
        }
        return result
}
