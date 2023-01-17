.PHONY: run
run:
	present

.PHONY: deps
deps:
	go install golang.org/x/tools/cmd/present@latest
