package main

import (
  // "encoding/json"
  "fmt"
  "flag"
  "io/ioutil"
  "strconv"
  "strings"
  "tebot/parser"
  "tebot/runner"
  // "tebot/database"
)

func main() {
  // out, _ := database.QueryPostgres("postgres://hk202401:hk202401@123456@192.168.20.28:5432/hktest", "select 1", "1");
  // data, _ := json.MarshalIndent(out, "", "  ")
  // fmt.Println(string(data))

  ptrScripts := flag.String("scripts", "", "tebot测试脚本")
  ptrSelenium := flag.String("selenium", "", "selenium的架包路径")
  ptrMobile := flag.String("mobile", "false", "手机模式")
  flag.Parse()

  if *ptrScripts == "" {
    Usage()
    return
  }

  if *ptrSelenium == "" {
    *ptrSelenium = "./selenium-server-4.43.0.jar"
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
