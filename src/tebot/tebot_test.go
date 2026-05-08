package main

import (
  "testing"
  "io/ioutil"
  "fmt"
  "strconv"

  "tebot/parser"
  // "tebot/runner"
)

func TestParse(t *testing.T) {
  buf, _ := ioutil.ReadFile("../../testdata/#001.tbt")
  ops := parser.Parse(string(buf))

  fmt.Println("the operation count is " + strconv.Itoa(len(ops.Operations)))

  for _, op := range ops.Operations {
    fmt.Println("action: " + op.Action)  
    fmt.Println("selector: " + op.GetSelector())
    fmt.Println("value: " + op.Value)
    if op.Assert != nil {
      fmt.Println("dsn: " + op.Assert.GetDsn())
      fmt.Println("sql: " + op.Assert.GetSql())
      fmt.Println("expected: " + op.Assert.GetExpected())
    }
    fmt.Println("--------------------------------")
  }
  
  // runner.Run(ops)
}
