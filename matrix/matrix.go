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
			go func(coll []float64, row []float64, r int, c int, cha chan float64){
				result := 0.0
				for k:=0; k< len(coll); k++{
					result += coll[k] * row[k]
				}
				results[r*cols+c] = result
				cha<-1
			}(col, b.Matrix[j*cols:j*cols+cols], i, j, ch)
                }
        }

	<-ch

        return Matrix2d{rows, cols, results}
}

func (m Matrix2d) getColumn(col int) []float64{
        result := make([]float64, m.Rows)
        cols := m.Columns
        for i :=0; i< m.Rows; i++{
                result[i] = m.Matrix[i*cols + col]
        }
        return result
}
