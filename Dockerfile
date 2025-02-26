FROM golang:1.22 as gobuild

WORKDIR /home

# COPY go.mod go.sum .
COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -a -trimpath  -o slugifyer .

FROM alpine

COPY --from=gobuild /home/slugifyer /

USER 65532:65532

ENTRYPOINT [ "/slugifyer" ]
