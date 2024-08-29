package main

import (
    "bytes"
    "fmt"
    "reflect"
)



func testSlice() {
    // slice
    a := make([]int, 32)
    b := a[1:16]
    fmt.Printf("&a:%p, &b:%p\n", &a, &b)
    a = append(a, 1)
    a[2] = 42
    fmt.Printf("&a:%p, &b:%p\n", &a, &b)

    path := []byte("AAAA/BBBBBBBBB")
    sepIndex := bytes.IndexByte(path,'/')
    dir1 := path[:sepIndex]
    //dir1 := path[:sepIndex:sepIndex] // full slice expression
    dir2 := path[sepIndex+1:]
    fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
    fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB
    dir1 = append(dir1,"suffix"...)
    fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
    fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB
}



func deepCmp() {
    //v1 := data{}
    //v2 := data{}
    //fmt.Println("v1 == v2:",reflect.DeepEqual(v1,v2))    //prints: v1 == v2: true

    m1 := map[string]string{"one": "a","two": "b"}
    m2 := map[string]string{"two": "b", "one": "a"}
    fmt.Println("m1 == m2:",reflect.DeepEqual(m1, m2))    //prints: m1 == m2: true
    s1 := []int{1, 2, 3}
    s2 := []int{1, 2, 3}
    fmt.Println("s1 == s2:",reflect.DeepEqual(s1, s2))    //prints: s1 == s2: true
}



// interface
type Printable interface {
    ToString()
}

type Country struct {
    Name string
}

type City struct {
    Name string
}

func (c Country) ToString() {
    fmt.Println(c.Name)
}

func (c City) ToString() {
    fmt.Println(c.Name)
}

func PrintName(printable Printable) {
    printable.ToString()
}

var _ Printable = (*Country)(nil)
var _ Printable = (*City)(nil)

func testInterface() {
    d1 := Country{"USA"}
    d2 := City{"Beijing"}
    PrintName(d1)
    PrintName(d2)
}

func main() {
    //testSlice()
    //deepCmp()
    testInterface()
}
