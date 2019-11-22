 FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY . .
#COPY front/dist/ ./public/
#COPY go.sum .

# RUN go mod download

#COPY . .

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN CGO_ENABLED=0 GOOS=linux go build
RUN ip link show

FROM scratch
COPY --from=builder /app/server /app/
COPY --from=builder /app/public /app/
EXPOSE 80
ENTRYPOINT ["/app/server"]