FROM golang:1.11-alpine


RUN apk update && apk add --no-cache git curl build-base bash openssh-client shadow

RUN go get github.com/githubnemo/CompileDaemon && mv /go/bin/CompileDaemon /usr/bin
RUN mkdir -p /go/src/aproject/service0

ENV GO111MODULE=on

# mirror dev user
ARG UID
ARG USER
ARG SSH_PUB
ARG SSH_PRI

RUN if test "$USER" != 'root'; then groupadd --gid $UID $USER \
  && useradd --uid $UID --gid $USER --shell /bin/bash --create-home $USER; fi

RUN chown -R $UID /go

USER $USER

# use ssh for git operations (specifically `go get` in our case)
# also disable StrictHostKeyChecking to get around the "authenticity of host ... (yes/no)"
# RUN printf "[url \"git@bitbucket.org:\"]\n\tinsteadOf = https://bitbucket.org/\n" >> \
#     /home/$USER/.gitconfig
# RUN mkdir -p /home/$USER/.ssh
# RUN echo "StrictHostKeyChecking no " > /home/$USER/.ssh/config
# RUN echo $SSH_PUB > /home/$USER/.ssh/id_rsa.pub
# RUN echo $SSH_PRI > /home/$USER/.ssh/id_rsa


EXPOSE 8080

WORKDIR /go/src/aproject/service0
CMD CompileDaemon -build="go build -o /tmp/service0d /go/src/aproject/service0/cmd/service0d/main.go" \
    -command="/tmp/service0d" \
    -directory="/go/src/aproject/service0"