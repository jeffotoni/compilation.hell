.PHONY: all build list run

all: build list run

build:
	dotnet publish -c Release -r linux-x64 /p:PublishSingleFile=true /p:PublishTrimmed=True

list:
	ls -la bin/Release/net5.0/linux-x64/publish/

run:
	ASPNETCORE_ENVIRONMENT=Production ./bin/Release/net5.0/linux-x64/publish/api

run_fast:
	dotnet run

run_watch:
	dotnet watch run