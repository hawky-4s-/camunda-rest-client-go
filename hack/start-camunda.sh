#!/usr/bin/env bash -xe

command -v docker >/dev/null 2>&1 || { echo >&2 "I require docker but it's not installed.  Aborting."; exit 1; }
docker run -d -it -p 8080:8080 --name camunda-external-task-test camunda/camunda-bpm-platform:SNAPSHOT
