FROM --platform=${BUILDPLATFORM} golang:latest AS builder
WORKDIR /todolist
ENV CGO_ENABLED=0
COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main .

RUN ls

FROM scratch
COPY --from=builder /todolist/main .
COPY --from=builder /todolist/config_remote.yaml .
ENV CONFIG_FILE="config_remote.yaml"
EXPOSE 8080

ENTRYPOINT ["/main"]
