lint:
	golangci-lint run ./...

fmt: gofmt

gofmt:
	gofumpt -l -w .

install_gofumpt:
	which gofumpt || go install mvdan.cc/gofumpt@latest

test:
	go clean -testcache
	go test -race -v -count 1 -failfast ./... | \
		sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | \
		sed ''/SKIP/s//$$(printf "\033[33mSKIP\033[0m")/'' | \
		sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
