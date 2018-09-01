BIN="$(shell go env GOPATH)/bin"

DEP="$(BIN)/dep"

.PHONY:
all: $(BIN)/aqrs

# Build tools.
$(DEP):
	[ -f $@ ] || go get -u github.com/golang/dep/cmd/dep

# Build tasks.
$(BIN)/aqrs:
	go build -o $@ github.com/liftM/aquarius/cmd/aqrs

# Development tasks.
.PHONY: vendor
vendor: $(DEP)
	$< ensure -v

.PHONY: clean
clean:
	rm -f $(BIN)/aqrs
