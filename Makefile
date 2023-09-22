BINARY_NAME=birthdaymessenger.out
 
build:
	go build -o ${BINARY_NAME} cmd/birthdaymessenger/main.go
 
run:
	go build -o ${BINARY_NAME} cmd/birthdaymessenger/main.go
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}