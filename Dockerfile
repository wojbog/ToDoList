FROM golang:1.18.2
WORKDIR /app
COPY ./go.mod go.sum ./
RUN go mod download && go mod verify
COPY . /app/