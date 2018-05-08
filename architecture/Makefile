example-tcp-cli:
	go run client/main.go tcp ./client/main.go
example-udp-cli:
	go run client/main.go udp ./client/main.go
example-server:
	go run server/main.go

example-server-with-output:
	go run server/main.go >> log_server.txt
example-tcp-with-output:
	go run client/main.go tcp ./README.md >> log_client.txt
example-udp-with-output:
	go run client/main.go udp ./README.md >> log_client.txt

clean:
	rm ./*.txt
