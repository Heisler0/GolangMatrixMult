package main

import(
        "github.com/Heisler0/matrixmult/matrix"
        "math/rand"
	"time"
	"fmt"
	"runtime"
)

func main(){

	runtime.GOMAXPROCS(4)

        rows, cols := 3000, 3000
        values := make([]float64, rows * cols)

        for i := 0; i < rows; i++{
                for j := 0; j < cols; j++{
                        values[i*cols + j] = rand.Float64()
                }
        }

        m := matrix.Matrix2d{rows, cols, values}

	runs := 100

	totalTime := 0.0

	for i := 0; i < runs; i++{

		start := time.Now()

		_ = m.Mult(m)

		totalTime += time.Since(start).Seconds()
	}

	fmt.Println(totalTime/float64(runs))
}
