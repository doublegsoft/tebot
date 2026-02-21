package main

import (
  "testing"
  "io/ioutil"

  "./parser"
  "./runner"
)

func TestParse(t *testing.T) {
  buf, _ := ioutil.ReadFile("../../testdata/#001.tbt")
  ops := parser.Parse(string(buf))

  if len(ops.Operations) != 1 {
    t.Fatalf("the operation count is" + string(len(ops.Operations)))
  }
  
  runner.Run(ops)
}
