#!/bin/sh

# Force Server Operator
if ! { [ -z ${SERVER_OP+x} ] || [ -z "$SERVER_OP" ]; }
then
  echo "$SERVER_OP"
  grep -sqxF "$SERVER_OP" ops.txt || echo "$SERVER_OP" >> ops.txt
fi

# Run Server
java $JAVA_ARGS -jar Tekkit.jar nogui
