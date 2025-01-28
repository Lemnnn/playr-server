.PHONY: run-server

run-server:
\tdocker run -p 8080:8080 --rm -v ${PWD}:/app:delegated --name playr-server-air playr-server