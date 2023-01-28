FROM golang:1.18-alpine 

RUN mkdir app 

COPY . /app 

WORKDIR /app 

# Install soda
RUN go install github.com/gobuffalo/pop/v6/soda@latest

RUN CGO_ENABLED=0 go build -o notesApp ./cmd/api

CMD [ "/app/notesApp" ]
