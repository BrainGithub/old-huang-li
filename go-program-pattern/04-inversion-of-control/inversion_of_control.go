package main

import (
    "errors"
    "fmt"
    "reflect"
    "runtime"
)

type Undo []func()

func (undo *Undo) Add(f func()) {
    *undo = append(*undo, f)
}

func (undo *Undo)Undo() error {
    functions := *undo
    fmt.Printf("undo:%p, functions:%p\n", undo, &functions)
    index := len(functions) - 1
    if index < 0 {
        return errors.New("no undo functions to do")
    }
    if f := functions[index]; f != nil {
        f()
        functions[index] = nil
    }
    *undo = functions[:index]
    fmt.Printf("undo:%p, functions:%p\n", undo, &functions)
    return nil
}

type IntSet struct {
    data map[int]bool
    undo Undo
}

func NewIntSet() IntSet {
    return IntSet{data: make(map[int]bool)}
}

func (set *IntSet) ToString() {
    fmt.Printf("\n\nToString start\n")
    fmt.Printf("\tdata:%+v\n\tfunctions:%+v\n", set.data, set.undo)
    for _, f := range set.undo {
        //fmt.Printf("\t\t%v\n", reflect.ValueOf(f).NumMethod())
        fmt.Printf("\t\t%v\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
    }
    fmt.Printf("ToString end\n")

}

func (set *IntSet) Undo() error {
    return set.undo.Undo()
}

func (set *IntSet) Contains(x int) bool {
    return set.data[x]
}

func (set *IntSet) Add(x int) {
    if !set.Contains(x) {
        set.data[x] = true
        set.undo.Add(func() {set.Delete(x) })
    } else {
        set.undo.Add(nil)
    }
}

func (set *IntSet) Delete(x int) {
    if set.Contains(x) {
        delete(set.data, x)
        set.undo.Add(func() {set.Add(x) })
    } else {
        set.undo.Add(nil)
    }
}

func main() {
    s := NewIntSet()
    s.ToString()

    s.Add(1)
    s.ToString()

    s.Add(2)
    s.ToString()

    s.Undo()
    s.ToString()

}
