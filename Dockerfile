# On Golang Will build the projects 
FROM golang:1.19-alpine as build

WORKDIR /app 

COPY . ./

RUN go mod tidy 

RUN GOOS=linux GOARCH=amd64 go build -o manga cmd/main.go

# Stage Deployer
# OS Alpine Will Running The app with copy all data from build stage
FROM alpine as main

WORKDIR /app

COPY --from=build /app/manga /app 

