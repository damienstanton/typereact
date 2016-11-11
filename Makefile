dev:
	npm i && npm link webpack
	webpack
	cd server && go build -o ./bin/server && ./bin/server

test:
	npm i && npm link webpack
	webpack
	cd server && go test -cover && go build -o ./bin/server
