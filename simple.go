package gomobile_lib

import "fmt"

func Multiply(value int32) (int32, error) {
    fmt.Println(value)

    if value < 0 || value > 10 {
        return 0, fmt.Errorf("value out of range: must be within the range of 0 to 10")
    }

    return value * 2, nil
}