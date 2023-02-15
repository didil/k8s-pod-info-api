MYGOBIN = "$(PWD)/bin"

install-tools:
	@echo MYGOBIN: $(MYGOBIN)
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | GOBIN=$(MYGOBIN) xargs -tI % go install %

test:
	go test ./...

lint:
	$(MYGOBIN)/golangci-lint run

run-dev: 
	go run main.go
