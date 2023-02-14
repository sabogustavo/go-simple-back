FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

#RUN go install
RUN go build -o /godocker

EXPOSE 2000

#CMD ["go", "run", "main.go"]
CMD [ "/godocker" ]