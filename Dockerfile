FROM golang:1.11

WORKDIR ${GOPATH}/src/github.com/hohihohi/k8s-client-go-sample

# Install some middleware for development
RUN apt update && \
    apt upgrade -y && \
    apt install -y git

# Add Gopkg.toml
ADD ./Gopkg.toml ${GOPATH}/src/github.com/hohihohi/k8s-client-go-sample/Gopkg.toml
ADD ./Gopkg.lock ${GOPATH}/src/github.com/hohihohi/k8s-client-go-sample/Gopkg.lock

# Install external packages for development
RUN go get -u github.com/golang/dep/cmd/dep && \
    go get -u github.com/alecthomas/gometalinter && \
    gometalinter --install --update