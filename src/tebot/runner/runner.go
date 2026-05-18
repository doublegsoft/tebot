package runner

import (
  "fmt"
  "os"
  "strconv"
  "time"
  "io/ioutil"
  "strings"

  "github.com/tebeka/selenium"
  "github.com/go-vgo/robotgo"
  "golang.design/x/clipboard"

  "tebot/model"
  "tebot/database"
)

func Run(ops *model.TebotOperations, seleniumPath string, mobile bool) {
  const (
    port = 4444
  )

  opts := []selenium.ServiceOption{
    // selenium.StartFrameBuffer(),                     // Start an X frame buffer for the browser to run in.
    selenium.GeckoDriver("./geckodriver"),        // Specify the path to GeckoDriver in order to use Firefox.
    selenium.Output(os.Stderr),                         // Output debug information to STDERR.
  }
  selenium.SetDebug(false)

  /*!
  ** clipboard initialization
  */
  err := clipboard.Init()
  if err != nil {
    panic(err)
  }

  service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
  if err != nil {
    panic(err) // panic is used only as an example and is not otherwise recommended.
  }
  defer service.Stop()

  firefoxOpts := map[string]interface{}{
    "args": []string{
        "-profile",
        "./firefox-profile",
    },
  }

  caps := selenium.Capabilities {
    "browserName":          "firefox",
    "takeScreenshot":       false,
    "useTechnologyPreview": true,
    "moz:firefoxOptions":   firefoxOpts,
  }

  webdriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
  if err != nil {
    panic(err)
  }
  defer webdriver.Quit()

  // Navigate to the simple playground interface.
  count := len(ops.Operations)
  
  for i := 0; i < count; i++ {
    op := ops.Operations[i]
    switch op.Action {
    case "goto":
      if err := webdriver.Get(op.GetSelectorOrValue()); err != nil {
        fmt.Println("Error to execute: " + op.Script)
        panic(err)
      }
      time.Sleep(3 * time.Second)
      if mobile {
        webdriver.ResizeWindow("", 375, 812);
      } else {
        webdriver.MaximizeWindow("")
      }

      break
    case "click":
      elm := Locate(&webdriver, op)
      if elm == nil {
        break;
      }
      (*elm).Click()
      break
    case "input":
      elm := Locate(&webdriver, op)
      if elm == nil {
        break;
      }
      tagName, _ := (*elm).TagName()
      if tagName == "select" {
        (*elm).Click()
        option, _ := (*elm).FindElement(selenium.ByTagName, "option")
        option.Click()
        (*elm).Click()
      } else {
        (*elm).Clear()
        (*elm).SendKeys(op.GetValue())
      }
      break
    case "capture":
      bytes, _ := webdriver.Screenshot()
      now := time.Now()
      filename := now.Format("20060102150405")
      os.Mkdir("./screenshot", 0755)
      ioutil.WriteFile("./screenshot/" + filename + ".png", bytes, 0755)
      break
    case "save":
      /*!
      ** save content to local file.
      **
      ** @since 2.0
      */
      pageContent, _ := webdriver.PageSource()
      now := time.Now()
      filename := now.Format("20060102150405")
      os.Mkdir("./html", 0755)
      ioutil.WriteFile("./html/" + filename + ".html", []byte(pageContent), 0755)
      break
    case "paste":
      content := clipboard.Read(clipboard.FmtText)
      now := time.Now()
      filename := now.Format("20060102150405")
      os.Mkdir("./content", 0755)
      ioutil.WriteFile("./content/" + filename + ".txt", []byte(content), 0755)
      break
    case "move":
      /*!
      ** allow mouse move
      **
      ** @since 2.0
      */
      source := op.GetSelector()
      // target := op.GetValue()

      strs := strings.Split(source, ",")
      x, _ := strconv.ParseInt(strings.Trim(strs[0], " "), 10, 32)
      y, _ := strconv.ParseInt(strings.Trim(strs[1], " "), 10, 32)
      robotgo.MoveMouseSmooth(int(x), int(y), 1.0, 1.5)
      time.Sleep(1 * time.Second)
      robotgo.Click()
      // strs = strings.Split(target, ",")
      // x, _ = strconv.ParseInt(strings.Trim(strs[0], " "), 10, 32)
      // y, _ = strconv.ParseInt(strings.Trim(strs[1], " "), 10, 32)
      // robotgo.DragSmooth(int(x), int(y), 1.0, 3.0)
      break
    case "scroll":
      y, _ := strconv.ParseInt(op.GetSelector(), 10, 32)
      // robotgo.ScrollMouse(int(y), "down");
      for i := 0; i < int(y); i++ {
        // robotgo.AddEvent("down")
        fmt.Println("########################")
        //robotgo.KeyTap("down", "fn")
        //robotgo.ScrollMouse(100, "down");
        time.Sleep(1 * time.Second)
      }
      // robotgo.KeyTap("down", "fn")
      break
    case "assert":
      // elm := Locate(&webdriver, op)
      // txt, _ := (*elm).Text()
      // if strings.Index(txt, op.GetSelector()) == -1 {
      //   bytes, _ := webdriver.Screenshot()
      //   now := time.Now()
      //   fn := now.Format("20060102150405")
      //   ioutil.WriteFile("./"+ fn + ".assert.png", bytes, 0644)
      // }
      database.QueryPostgres(op.Assert.GetDsn(), op.Assert.GetSql(), op.Assert.GetExpected())
      break
    case "sleep":
      sec, _ := strconv.ParseInt(op.GetSelector(), 10, 32)
      time.Sleep(time.Duration(sec) * time.Second)
      break
    }

    time.Sleep(1 * time.Second)
  }
}

func Locate(webdriver *selenium.WebDriver, op *model.TebotOperation) *selenium.WebElement {
  ret, err := (*webdriver).FindElement(selenium.ByXPATH, op.GetSelector())
  if err != nil {
    fmt.Println("Error to execute: " + op.Script)
    return nil
  }
  return &ret
}
