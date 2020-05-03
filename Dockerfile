FROM openjdk:8u212-jre-alpine as builder

ENV JAVA_ARGS="-Xmx3G -Xms2G"
ENV SERVER_OP=""

WORKDIR /minecraft
RUN apk add unzip wget

RUN wget -O /tmp/tekkit.zip http://servers.technicpack.net/Technic/servers/tekkit/Tekkit_Server_3.1.2.zip
RUN unzip /tmp/tekkit.zip -d /minecraft/
RUN rm launch.bat

FROM golang:alpine as tool-builder
WORKDIR /tool
RUN apk add git

FROM tool-builder as rcon-cli
ARG RCON_CLI_VER=1.4.7

RUN git clone --branch ${RCON_CLI_VER} https://github.com/itzg/rcon-cli.git .
RUN CGO_ENABLED=0 go build

FROM openjdk:8u212-jre-alpine
WORKDIR /minecraft

ENV RCON_PORT=25575 RCON_PASSWORD=minecraft

COPY --from=rcon-cli /tool/rcon-cli .
COPY --from=builder /minecraft /minecraft

COPY ./launch.sh /minecraft/launch.sh
COPY ./server.properties /minecraft/server.properties
RUN chmod +x /minecraft/launch.sh

VOLUME /minecraft
EXPOSE 25565

ENTRYPOINT ["/minecraft/launch.sh"]
