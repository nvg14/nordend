#!/bin/bash

if [ -d "/etc/secrets/" ]; then
    echo "loading secrets.."
    for file in $(ls /etc/secrets/*);
    do
        env_var="$(basename -- $file)"
        while read -rd $'' line
        do
            export "$line"
        done < <(jq -r <<<"$(cat $file)" 'to_entries|map("\(.key)=\(.value)\u0000")[]')
    done
else
    echo "Could not find /etc/secrets directory."
fi

if [ "$1" = "server" ]; then
    ./server --port $PORT
else
    echo "Unknown Module"
    exit 1
fi