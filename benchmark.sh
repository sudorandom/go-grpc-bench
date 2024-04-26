#!/usr/bin/env bash

set -e

BACKEND=$1
SCENARIO=$2
WITH_VT=$3

mkdir -p ./results/

go run cmd/${BACKEND}/main.go $WITH_VT &
BACKEND_PID=$!
clean_up () {
    ARG=$?
    echo "> clean up ${BACKEND_PID}"
    pkill -P $BACKEND_PID
    exit $ARG
} 
trap clean_up EXIT

sleep 2

if ps -p $BACKEND_PID > /dev/null
then
   echo "> backend is running"
else
    echo "> backend failed"
    exit 1
fi

echo "> running benchmark"
ghz --config "./scenarios/${SCENARIO}.json" -o "results/$BACKEND-$SCENARIO$WITH_VT.txt" -O summary
