package main

func MapStrToStr(arr []string, fn func(s string) string) []string {
    var newArray []string
    for _, it := range arr {
        newArray = append(newArray, fn(it))
    }
    return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
    var newArray []int
    for _, it := range arr {
        newArray = append(newArray, fn(it))
    }
    return newArray
}


func Reduce(arr []string, fn func(s string)int) int {
    sum := 0
    for _, it := range arr {
        sum += fn(it)
    }
    return sum
}

func Filter(arr []int, fn func(n int) bool) []int {
    var newArray []int
    for _, i := range arr {
        if fn(i) {
            newArray = append(newArray, i)
        }
    }

    return newArray
}



func main() {

}
