#!/bin/bash

NAME=$1
MOOD=$2
TIMESTAMP=$3

curl -X POST http://192.168.49.2:30080/mood \
  -H "Content-Type: application/json" \
  -H "X-Timestamp: $TIMESTAMP" \
  -d "{
        \"name\": \"$NAME\",
        \"mood\": \"$MOOD\"
      }"
