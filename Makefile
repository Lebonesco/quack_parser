GOCMD=go 
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get 
GENERATOR=quack.bnf
GENERATE=../../../../bin/gocc
BINARY_NAME=scanner


all: test build run
build: 
	$(GENERATE) $(GENERATOR) 
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GENERATE) $(GENERATOR)
	$(GOTEST) -v 
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME) util token lexer parser errors
	rm -f LR1_conflicts.txt LR1_sets.txt first.txt lexer_sets.txt terminals.txt
run:
	 $(GENERATE) $(GENERATOR) 
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME) $(file)
deps:
	$(GOGET) github.com/goccmack/gocc