#!/bin/sh
echo $SERVER_OP > /minecraft/ops.txt
java $JAVA_ARGS -jar Tekkit.jar nogui
