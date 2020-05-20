FROM openjdk:8u212-jre-alpine as builder
ARG TEKKIT_VERSION=3.1.2

WORKDIR /minecraft
RUN apk add unzip wget

RUN wget --quiet -O /tmp/tekkit.zip http://servers.technicpack.net/Technic/servers/tekkit/Tekkit_Server_${TEKKIT_VERSION}.zip
RUN unzip /tmp/tekkit.zip -d /minecraft/
RUN rm launch.bat

FROM golang:alpine as tool-builder
WORKDIR /tool
RUN apk add git

FROM tool-builder as rcon-cli
ARG RCON_CLI_VER=1.4.7

RUN git clone --branch ${RCON_CLI_VER} https://github.com/itzg/rcon-cli.git .
RUN CGO_ENABLED=0 go build

FROM tool-builder as trapper

COPY ./tools/trapper .
RUN CGO_ENABLED=0 go build

FROM tool-builder as propenv

COPY ./tools/propenv .
RUN CGO_ENABLED=0 go build

FROM openjdk:8u212-jre-alpine
WORKDIR /minecraft

ENV JAVA_ARGS="-Xmx3G -Xms2G" \
  SERVER_OP="" \
  ENABLE_RCON="true" \
  RCON_PORT="25575" \
  RCON_PASSWORD="minecraft" \
  MOTD="Tekkit Classic powered by Docker"

COPY --from=rcon-cli /tool/rcon-cli /bin/.
COPY --from=trapper /tool/trapper /bin/.
COPY --from=propenv /tool/propenv /bin/.
COPY --from=builder /minecraft /minecraft

COPY ./launch.sh /minecraft/launch.sh
RUN chmod +x /minecraft/launch.sh

VOLUME /minecraft
EXPOSE 25565

ENTRYPOINT ["/minecraft/launch.sh"]
