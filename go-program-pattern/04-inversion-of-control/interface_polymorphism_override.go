package main

import "fmt"

type Widget struct {
    X, Y int
}

type Label struct {
    Widget // Embedding (delegation)
    Text string // Aggregation
}

type Button struct {
    Label // Embedding (delegation)
}

type ListBox struct {
    Widget // Embedding (delegation)
    Texts []string // Aggregation
    Index int // Aggregation
}

type Painter interface {
    Paint()
}

type Clicker interface {
    Click()
}

func (label *Label) Paint() {
    fmt.Printf("%p Label.Paint(%q)\n", &label, label.Text)
}

//因为这个接口可以通过 Label 的嵌入带到新的结构体，
//所以，可以在 Button 中重载这个接口方法 override
func (button *Button) Paint() {
    fmt.Printf("%p Button.Paint(%q)\n", &button, button.Text)
}

func (button *Button) Click() {
    fmt.Printf("%p Button.Click(%q)\n", &button, button.Text)
}


func (listBox *ListBox) Paint() {
    fmt.Printf("%p ListBox.Paint(%q)\n", &listBox, listBox.Texts)
}

func (listBox *ListBox) Click() {
    fmt.Printf("%p ListBox.Click(%q)\n", &listBox, listBox.Texts)
}

func main() {
    label := Label{
        Widget{10, 10},
        "State:",
    }
    label.X = 11
    label.Y = 12

    ok := Button{Label{Widget{10, 70}, "OK"}}
    cancel := Button{Label{Widget{50, 70}, "Cancel"}}
    listBox := ListBox{
        Widget{10, 40},
        []string{"AL", "AK", "AZ", "AR"},
        0,
    }

    for _, painter := range []Painter{&label, &ok, &cancel, &listBox} {
        painter.Paint()
    }

    for _, widget := range []interface{} {&label, &ok, &cancel, &listBox} {
        widget.(Painter).Paint()
        if clicker, ok := widget.(Clicker); ok {
            clicker.Click()
        } else {
            fmt.Printf("%+v is not a clicker\n", widget)
        }
    }
}
