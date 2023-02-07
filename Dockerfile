FROM golang

WORKDIR /app

COPY . . 

RUN go mod download

RUN go build -o live

ENTRYPOINT ./live
