package parser

import (
    "testing"
    "io/ioutil"
    "fmt"

)


func TestParse(t *testing.T) {
    buf, _ := ioutil.ReadFile("../../../testdata/#001.tbt")
    fmt.Println(string(buf))
    ops := Parse(string(buf))

    //if len(ops.Operations) != 2 {
    //   t.Fatalf("the operation count is " + string(len(ops.Operations)))
    //}

    op := ops.Operations[0]

    fmt.Println("action: " + op.Action)
    fmt.Println("selector: " + op.GetSelector())
    fmt.Println("value: " + op.Value)

    op = ops.Operations[1]

    fmt.Println("action: " + op.Action)
    fmt.Println("selector: " + op.GetSelector())
    fmt.Println("value: " + op.Value)
}