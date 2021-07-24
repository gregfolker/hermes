default:
	go build -o _out/hermes main.go

run:
	go run main.go

clean:
	rm -rf _out/
