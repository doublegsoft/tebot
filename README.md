## Generates parser

```bash
cd gram

~/export/opt/jdk-21.0.1/Contents/Home/bin/java -jar ~/export/opt/antlr4/antlr-4.13.0-complete.jar Tebot.g4 -Dlanguage=Go -o ../src/tebot/parser

cd ..
```

## Build

```bash

cd src/tebot

go build

cd ../..

```

## Get Dependencies

```bash
go get github.com/antlr4-go/antlr/v4

go get github.com/tebeka/selenium

go get github.com/go-vgo/robotgo
```

## Run Tests

```
./tebot -selenium=../../vendor/selenium-server-standalone-3.141.59.jar -scripts=/Users/christian/export/local/projs/goodmanager/dev/e2e/gm.tebot
```