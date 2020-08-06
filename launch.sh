#!/bin/sh

# Force Server Operator
if ! { [ -z ${SERVER_OP+x} ] || [ -z "$SERVER_OP" ]; }
then
  grep -sqxF "$SERVER_OP" ops.txt || echo "$SERVER_OP" >> ops.txt
fi

# Run Server
propenv
exec trapper java $JAVA_ARGS -jar Tekkit.jar nogui
