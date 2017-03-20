package main

import(
        "github.com/Heisler0/matrixmult/matrix"
        "math/rand"
)

func main(){

        rows, cols := 10, 10
        values := make([]float64, rows * cols)

        for i := 0; i < rows; i++{
                for j := 0; j < cols; j++{
                        values[i*cols + j] = rand.Float64()
                }
        }

        m := matrix.Matrix2d{rows, cols, values}

        _ = m.Mult(m)
}

