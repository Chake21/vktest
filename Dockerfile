FROM golang:1.19

RUN mkdir /app
ADD . /app/
COPY Makefile /app/cmd
ARG PORT
ENV SERVICE_PORT $PORT
WORKDIR /app/cmd
EXPOSE $SERVICE_PORT
RUN go build -o main .
CMD ["/app/cmd/main"]
WORKDIR /app
