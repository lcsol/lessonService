FROM golang:latest

LABEL maintainer="lcsol"

# defines the working directory
WORKDIR /app

COPY go.mod .

COPY go.sum .
# fetch dependancies
RUN go mod download
# copy the source code
COPY . .
# container port
ENV PORT 8000

RUN go build
# run app
CMD ["./lesson-service-draft"]