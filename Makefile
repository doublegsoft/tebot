
all:
	cd gram && java -jar ~/export/opt/antlr4/antlr-4.13.0-complete.jar Tebot.g4 -Dlanguage=Go -o ../src/tebot/parser
	cd src/tebot && go build && cd ../..
	cp src/tebot/tebot rel/darwin
