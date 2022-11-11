FROM golang:1.18.3

WORKDIR /app

# Install the Delve debugger
RUN go install github.com/go-delve/delve/cmd/dlv@v1.8.3
# Install AIR for live reloading
RUN go install github.com/cosmtrek/air@v1.40.2

CMD ["air"]
