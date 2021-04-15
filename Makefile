default:
	go build -o _out/pbld main.go

run:
	go run main.go

clean:
	rm -rf _out/
