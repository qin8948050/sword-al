package analysis

import "fmt"

func PrintMatrixInCircle(matrix [][]int, rows, columns int) {
	if rows == 0 && columns == 0 {
		return
	}
	start := 0
	for rows > start*2 && columns > start*2 {
		printMatrixInCircle(matrix,rows,columns,start)
		start++
	}
}

func printMatrixInCircle(matrix [][]int,rows,columns,start int) {
	endX:=columns-start-1
	endY:=rows-start-1
	for i := start; i <= endX; i++ {
		fmt.Println(matrix[start][i])
	}
	if start<endY{
		for i := start+1; i <=endY; i++ {
			fmt.Println(matrix[i][endX])
		}
	}
	if start<endX && start<endY{
		for i := endX-1; i >=start; i-- {
			fmt.Println(matrix[endY][i])
		}
	}
	if start<endX && start<endY-1{
		for i := endY-1; i >=start+1; i-- {
			fmt.Println(matrix[i][start])
		}
	}
}