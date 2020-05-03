FROM anapsix/alpine-java as builder

ENV JAVA_ARGS="-Xmx3G -Xms2G"
ENV SERVER_OP=""

WORKDIR /minecraft
RUN apk add --no-cache unzip wget

RUN wget -O /tmp/tekkit.zip http://servers.technicpack.net/Technic/servers/tekkit/Tekkit_Server_3.1.2.zip
RUN unzip /tmp/tekkit.zip -d /minecraft/
RUN rm launch.bat

FROM anapsix/alpine-java
WORKDIR /minecraft

RUN apk add --no-cache tmux
COPY --from=builder /minecraft /minecraft
COPY ./launch.sh /minecraft/launch.sh
RUN chmod +x /minecraft/launch.sh

VOLUME /minecraft
EXPOSE 25565

ENTRYPOINT ["/minecraft/launch.sh"]
