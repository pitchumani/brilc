all: main.go
	go build -o brilc main.go

brilc: all

run: brilc
	@echo Running brilc with sample nop.json
	./brilc nop.json

clean:
	rm -f brilc
