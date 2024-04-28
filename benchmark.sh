#!/usr/bin/env bash

set -e

SCENARIO=$1
PROFILE=$2
DATE=$(date '+%Y-%m-%d')

OUTPUTFILE="results/${DATE}/${SCENARIO}/${PROFILE}.json"
PROFILEFILE="pprof/${DATE}/${SCENARIO}/${PROFILE}.pb.gz"
mkdir -p "$(dirname ${OUTPUTFILE})"
mkdir -p "$(dirname ${PROFILEFILE})"

# Start backend
. "./profiles/${PROFILE}.sh"
BACKEND_PID=$!
clean_up () {
    ARG=$?
    pkill -P $BACKEND_PID
    exit $ARG
} 
trap clean_up EXIT

sleep 2

# Make sure the backend is still alive
ps -p $BACKEND_PID > /dev/null || (echo "> backend failed"; exit 1)

# Capture profiling data with pprof
go tool pprof -proto -output "${PROFILEFILE}" "http://localhost:6669/debug/pprof/profile?seconds=30" &

echo "> running benchmark (${SCENARIO}) (${PROFILE})"
ghz --config "./scenarios/${SCENARIO}.json" -o "${OUTPUTFILE}" -O json

mv "${OUTPUTFILE}" "${OUTPUTFILE}.full"
jq 'del(.details)' "${OUTPUTFILE}.full" > "$OUTPUTFILE"
rm "${OUTPUTFILE}.full"

jq -r '"rps: \(.rps), average: \(.average/100000), slowest: \(.slowest/100000), fastest: \(.fastest/100000)"' "$OUTPUTFILE"
echo "> finished benchmark (${SCENARIO}) (${PROFILE})"
