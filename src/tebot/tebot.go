package main

import (
  "fmt"
  "flag"
  "io/ioutil"
  "strconv"
  "strings"
  "tebot/parser"
  "tebot/runner"
)

func main() {
  ptrScripts := flag.String("scripts", "", "tebot测试脚本")
  ptrSelenium := flag.String("selenium", "", "selenium的架包路径")
  ptrMobile := flag.String("mobile", "false", "手机模式")
  flag.Parse()

  if *ptrScripts == "" {
    Usage()
    return
  }

  if *ptrSelenium == "" {
    *ptrSelenium = "./selenium-server-standalone-3.141.59.jar"
  }

  scriptPaths := strings.Split(*ptrScripts, ";")

  scripts := ""

  for i := 0; i < len(scriptPaths); i++ {
    scriptPaths[i] = strings.TrimSpace(scriptPaths[i])
    if scriptPaths[i] == "" {
      continue
    }
    buf, _ := ioutil.ReadFile(scriptPaths[i])
    scripts += string(buf)
  }

  ops := parser.Parse(scripts)

  mobile, _ := strconv.ParseBool(*ptrMobile)
  runner.Run(ops, "selenium-server-standalone-3.141.59.jar", mobile)
}

func Usage() {
  fmt.Println("正确用法：")
  fmt.Println("\ttebot -scripts=tebot测试脚本")
}
