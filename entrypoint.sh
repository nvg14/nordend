#!/bin/bash

if [ "$1" = "server" ]; then
    ./server --port $PORT
else
    echo "Unknown Module"
    exit 1
fi