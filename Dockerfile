


#stage 1

FROM golang:alpine AS builder
LABEL authors="ArminEyvazi"

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64

# working directory
WORKDIR /build

COPY go.mod .
COPY go.sum .


RUN go mod download

#copy our application code into the container
COPY . .

RUN go build -o main .

#stage 2
FROM scratch

#arguments
ARG MY_APP_PORT
ARG DB_HOST
ARG DB_PORT
ARG JWT_TOKEN_SECRET

# enviroment varibales for the application
ENV MY_APP_PORT=${MY_APP_PORT}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV JWT_TOKEN_SECRET=${JWT_TOKEN_SECRET}

#copy frome stage1 image
COPY --from=builder /build/main /

#expose the port app
EXPOSE 8080

#command to run
ENTRYPOINT ["/main"]
