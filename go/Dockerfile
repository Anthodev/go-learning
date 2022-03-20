FROM golang:1.18-alpine

ARG USER_ID
ARG GROUP_ID

ENV UID=${USER_ID:-1000}
ENV GID=${GROUP_ID:-1000}

RUN addgroup -S app_group --gid ${GID}

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup app_group \
    --no-create-home \
    --uid "$UID" \
    app_user

WORKDIR /usr/src/app
RUN chown app_user:app_group -R /usr/src/app

USER app_user
