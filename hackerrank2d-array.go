// https://www.hackerrank.com/challenges/2d-array/problem?isFullScreen=true
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

//also my code
func hourglassCalc(a int, b int, arr [][]int32) int32{
    // return [[a, b], [a+1, b], [a+2, b], [a+1, b+1], [a, b+2], [a+1, b+2], [a+2, b+2]]
    // return [[a, b], [a, b+1], [a, b+2], [a+1, b+1], [a+2, b], [a+2, b+1], [a+2, b+2]]
    hourglassSize := 3
    var sum int32
    sum = 0
    for i := 0; i < hourglassSize; i++ {
        for j:= 0; j < hourglassSize; j++ {
            if i == 1 && j != 1 {
                continue
            } else {
                sum += arr[a + i][b + j]
                
            }
        }
    }
    return sum    
}
// https://stackoverflow.com/questions/50510871/how-to-check-if-int-variable-has-been-set
// https://gobyexample.com/structs
type maxint32 struct {
    isSet bool
    val int32
}
func hourglassSum(arr [][]int32) int32 {
    // Write your code here
    // (0,0) : (0,0), (0,1), (0,2), (1,1), (2,0), (2,1), (2,2)
    
    var max maxint32
    max.isSet = false
    nrow := len(arr)/2 + 1
    // https://stackoverflow.com/questions/13449773/go-range-and-len-of-multidimensional-array
    
    for i := 0; i < nrow ; i++ { // 4 = 6/2 + 1
        for j := 0; j < nrow; j++ {
            // i,j is upper left most point of the hourglass in the array
            var sum int32
            sum = hourglassCalc(i, j, arr)
            fmt.Println(sum)
            if max.isSet {
                if max.val < sum {
                    max.val = sum
                }
            } else {
                max.val = sum
                max.isSet = true
            }

            
        }
    
    }
    return max.val
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != 6 {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := hourglassSum(arr)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
