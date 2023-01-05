#!/usr/bin/env bash

# TODO:
# - [ ] Support language templates

set -e

YEAR=${1:-$(date +%Y)}
DAY=${2:-$(date +%-d)}
SESSION_COOKIE=${AOC_SESSION_COOKIE:-}

fatal_error() {
    echo "$1"
    exit 1
}

[[ -z $SESSION_COOKIE ]] && fatal_error "Please provide AOC_SESSION_COOKIE environment variable"

CURRENT_DIR=$(pwd)
DIRECTORY="${CURRENT_DIR}/${YEAR}/${DAY}"
OUT_FILE="${DIRECTORY}/input.txt"
URL="https://adventofcode.com/${YEAR}/day/${DAY}/input"

mkdir -p "${DIRECTORY}"
curl -o "${OUT_FILE}" --cookie "session=${SESSION_COOKIE}" "${URL}"
