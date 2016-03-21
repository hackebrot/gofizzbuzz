.PHONY: tests

help:
	@echo "tests - run the test suite"

tests:
	go test "github.com/hackebrot/gofizzbuzz/gofizzbuzz"
