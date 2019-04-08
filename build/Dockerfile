FROM alpine:3.9

ENV OPERATOR=/usr/local/bin/goclient \
    USER_UID=1001 \
    USER_NAME=goclient

# install operator binary
COPY build/_output/bin/goclient ${OPERATOR}

COPY build/bin /usr/local/bin

RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
