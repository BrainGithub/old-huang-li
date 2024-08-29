package main

import "fmt"

type Info3 struct {
	Namespace string
	Name string
	OtherThings string
}

func (info *Info3) Visit(fn VisitorFunc3) error {
	return fn(info, nil)
}

type VisitorFunc3 func(*Info3, error) error


type Visitor3 interface {
	Visit(VisitorFunc3) error
}

type DecoratedVisitor struct {
	visitor Visitor3
	decorators []VisitorFunc3
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc3) error {
	return v.visitor.Visit(func(info *Info3, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func NewDecoratedVisitor(v Visitor3, fn ...VisitorFunc3) Visitor3 {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

func LoadFile(info *Info3, err error) error {
	info.Name = "Hao Chen"
	info.Namespace = "MegaEase"
	info.OtherThings = "We are running as remote team."
	return nil
}

func NameVisitorFunc(info *Info3, err error) error {
	fmt.Printf("NameVisitorFunc\n")
	return nil
}

func OtherVisitorFunc(info *Info3, err error) error {
	fmt.Printf("OtherVisitorFunc\n")
	return nil
}


func main() {
	info := Info3{}
	var v Visitor3 = &info
	v = NewDecoratedVisitor(v, NameVisitorFunc, OtherVisitorFunc)
	v.Visit(LoadFile)
}
