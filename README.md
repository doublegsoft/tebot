## Generates parser

```bash
cd gram

java -jar ~/export/opt/antlr4/antlr-4.13.0-complete.jar Tebot.g4 -Dlanguage=Go -o ../src/tebot/parser
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

## Run Test


## Run Script

```
./tebot -selenium=../../vendor/selenium-server-standalone-3.141.59.jar -scripts=/Users/christian/export/local/projs/goodmanager/dev/e2e/gm.tebot

java -jar selenium-server-4.43.0.jar standalone

./tebot -scripts=your.tbt
```

## Linux Firefox

[How to Install Firefox DEB on Ubuntu (Not Snap)](https://www.omgubuntu.co.uk/2022/04/how-to-install-firefox-deb-apt-ubuntu-22-04)