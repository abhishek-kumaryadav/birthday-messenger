FROM golang:1.21.1-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./ ./

RUN ls -lct
# RUN cat ./go.mods
# RUN cat ./go.mod
# RUN cat ./cmd/birthdaymessenger/main.go
RUN go build -o birthdaymessenger.out ./cmd/birthdaymessenger/main.go

# EXPOSE 8080

CMD [ "./birthdaymessenger.out" ]
