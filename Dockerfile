FROM golang:1.22 as gobuild

WORKDIR /home

# COPY go.mod go.sum .
COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags "-s -w" -a -trimpath  -o stoik-api .

FROM scratch

COPY --from=gobuild /home/stoik-api /

USER 65532:65532

ENTRYPOINT [ "/stoik-api" ]
