.PHONY: all build list run

all: build list run

build:
	dotnet publish -c Release -r linux-x64 /p:PublishSingleFile=true /p:PublishTrimmed=True

list:
	ls -la bin/Release/net5.0/linux-x64/publish/

run:
	./bin/Release/net5.0/linux-x64/publish/hello-console

