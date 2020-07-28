FROM golang:latest as kube-fpifetch
COPY . /app/
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -v -o kube-fpifetch

FROM alpine
RUN mkdir -p /app/
WORKDIR /app/
COPY --from=kube-fpifetch /app/cmd/kube-fpifetch .
ENTRYPOINT [ "/app/kube-fpifetch" ]