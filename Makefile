test:
	go test

test-cov:
	mkdir -p coverage
	go test -cover -coverprofile=coverage/coverage.out
	go tool cover -html=coverage/coverage.out -o coverage/index.html
	go tool cover -func=coverage/coverage.out
