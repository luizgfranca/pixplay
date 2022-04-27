from golang:1.15

WORKDIR /go/src
ENV path="/go/bin:${path}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update
RUN apt-get install build-essential protobuf-compiler librdkafka-dev -y
RUN go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go
RUN go get github.com/spf13/cobra
RUN wget https://github.com/ktr0731/evans/releases/download/0.9.1/evans_linux_amd64.tar.gz
RUN ls -la
RUN tar -xzvf evans_linux_amd64.tar.gz
RUN mv evans ../bin
RUN rm -f evans_linux_amd64.tar.gz

CMD ["tail", "-f", "/dev/null"]