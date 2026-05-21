# Tebot

基于 Go 的浏览器端到端（E2E）自动化工具。使用 ANTLR 解析 `.tbt` 脚本，通过 Selenium WebDriver 驱动 Firefox 执行操作，并支持截图、页面保存、剪贴板采集与 PostgreSQL 断言。

当前版本：**2.0**（见 `VERSION`）

## 特性

- 声明式脚本（`.tbt`），语法简洁，支持 `#` 行注释
- XPath 定位 DOM 元素
- 内置操作：`goto`、`click`、`input`、`sleep`、`capture`、`save`、`paste`、`move`、`scroll`、`html`、`assert` 等
- 可选手机视口（`-mobile`）
- 多脚本合并执行（`-scripts` 用分号分隔多个文件）
- 输出目录：`screenshot/`、`html/`、`content/`

## 项目结构

```
tebot/
├── gram/                 # ANTLR 语法 Tebot.g4
├── src/tebot/
│   ├── tebot.go          # CLI 入口
│   ├── parser/           # 词法/语法分析（由 gram 生成）
│   ├── runner/           # Selenium + robotgo 执行器
│   ├── model/            # 脚本 AST 模型
│   └── database/         # PostgreSQL 查询（assert）
├── testdata/             # 示例 .tbt 脚本
├── vendor/               # selenium-server JAR、各平台 geckodriver
├── screenshot/           # capture 截图输出（运行时创建）
├── html/                 # save / html 输出（运行时创建）
└── content/              # paste 剪贴板内容（运行时创建）
```

## 环境要求

| 依赖 | 说明 |
|------|------|
| Go | 1.25+（模块在 `src/tebot/go.mod`） |
| Java | 生成 parser；运行 Selenium 4 需 `java -jar` |
| Firefox | 通过 GeckoDriver 控制；Linux 建议使用 deb 版而非 Snap |
| GeckoDriver | 运行目录下 `./geckodriver`（`vendor/darwin` 或 `vendor/linux` 可复制） |
| Selenium Server | 默认 `vendor/selenium-server-4.43.0.jar` |

## 脚本语法（`.tbt`）

每条语句一行，形式为 `动作("参数")` 或 `动作("选择器" = "值")`。字符串使用双引号；`#` 开头为注释。

### 浏览器操作

| 动作 | 写法 | 说明 |
|------|------|------|
| `goto` | `goto("https://example.com")` | 打开 URL；之后按 `-mobile` 调整窗口大小 |
| `click` | `click("//button[@id='ok']")` | 点击 XPath 匹配元素 |
| `input` | `input("//input[@name='q']" = "关键词")` | 清空后输入；`<select>` 会选第一个 `option` |
| `sleep` | `sleep("3")` | 等待秒数（选择器位置传数字字符串） |
| `capture` | `capture("login")` | 全页截图 → `screenshot/{名}.{时间戳}.png` |
| `save` | `save("")` | 整页 HTML → `html/{时间戳}.html` |
| `html` | `html("//div[@id='main']" = "panel")` | 指定元素 outerHTML → `html/{名}.{时间戳}.html` |
| `paste` | `paste("")` | 系统剪贴板文本 → `content/{时间戳}.txt` |

### 桌面辅助（robotgo）

| 动作 | 写法 | 说明 |
|------|------|------|
| `move` | `move("100, 200")` 或 `move("锚点" = "x, y")` | 平滑移动鼠标到坐标 |
| `scroll` | `scroll("3")` | 向下滚动（数值为步数） |

### 数据库断言

```text
assert("postgres://user:pass@host:5432/db", "SELECT count(*) FROM t", "expected.json")
```

三个参数均为带引号的字符串：DSN、SQL、期望值文件路径（由 `database` 包执行查询）。

### 示例

`testdata/#001.tbt` — 本地站点导航、输入、鼠标移动与 assert：

```text
goto("http://localhost:8088")
capture("")

click("//span[@class='navbar-toggler-icon']")
input("//input[@class='form-control']" = "hello")
sleep("10")
move("//a[text()='基本资料']" = "100, 100")
assert("dsn", "sql", "file")
```

`testdata/#002.tbt` — 外部站点搜索、保存与剪贴板：

```text
goto("https://www.devv.ai")
input("//input[@aria-label='搜索']" = "hello")
click("//button[@aria-label='Search']")
save("")
paste("")
```

## 生成 Parser

修改 `gram/Tebot.g4` 后需重新生成 Go 代码：

```bash
cd gram
java -jar ~/export/opt/antlr4/antlr-4.13.0-complete.jar Tebot.g4 -Dlanguage=Go -o ../src/tebot/parser
cd ..
```

或使用 Makefile（路径需与本机 ANTLR JAR 一致）：

```bash
make
```

## 构建

```bash
cd src/tebot
go build
cd ../..
```

产物：`src/tebot/tebot`（已在 `.gitignore` 中忽略）。

## 依赖

在 `src/tebot` 目录下：

```bash
go mod tidy
```

主要模块：`github.com/antlr4-go/antlr/v4`、`github.com/tebeka/selenium`、`github.com/go-vgo/robotgo`、`golang.design/x/clipboard`、`github.com/jackc/pgx/v5`。

## 运行

### 1. 启动 Selenium（推荐 4.x）

```bash
java -jar vendor/selenium-server-4.43.0.jar standalone
```

或在另一终端由 tebot 自动拉起（传入 `-selenium` 指向 JAR）。

### 2. 准备运行环境

在**执行 tebot 的工作目录**下放置：

- `geckodriver`（可从 `vendor/darwin/geckodriver` 或 `vendor/linux/geckodriver` 复制）
- 可选：`firefox-profile/`（runner 中 Firefox 使用 `-profile ./firefox-profile`）

### 3. 执行脚本

```bash
cd src/tebot

# 单个脚本
./tebot -scripts=../../testdata/#001.tbt

# 多个脚本（按顺序拼接）
./tebot -scripts=../../testdata/#001.tbt;../../testdata/#002.tbt

# 指定 Selenium JAR
./tebot -selenium=../../vendor/selenium-server-4.43.0.jar -scripts=../../testdata/mobbin.tbt

# 手机视口（goto 后 375×812）
./tebot -mobile=true -scripts=../../testdata/mobbin.tbt
```

### 命令行参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `-scripts` | （必填） | `.tbt` 路径，多个用 `;` 分隔 |
| `-selenium` | `./selenium-server-4.43.0.jar` | Selenium Server JAR 路径 |
| `-mobile` | `false` | `goto` 后是否使用手机窗口尺寸 |

## 测试

仅解析、不启动浏览器：

```bash
cd src/tebot
go test -v -run TestParse
```

会读取 `testdata/#001.tbt` 并打印解析出的操作列表。

## Linux Firefox

Ubuntu 22.04+ 若需避免 Snap 版 Firefox，可参考：

[How to Install Firefox DEB on Ubuntu (Not Snap)](https://www.omgubuntu.co.uk/2022/04/how-to-install-firefox-deb-apt-ubuntu-22-04)

## 语法扩展

新动作或语法变更流程：

1. 编辑 `gram/Tebot.g4`
2. 重新生成 `src/tebot/parser/`
3. 在 `src/tebot/parser/parser.go`（listener）与 `src/tebot/model/` 中补充模型
4. 在 `src/tebot/runner/runner.go` 的 `switch op.Action` 中实现行为

## 许可与作者

（请按仓库实际情况补充 license 与联系方式。）
