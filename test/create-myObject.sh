#!/bin/bash

HOST=${1-"localhost"}

resp=$(curl -sX POST http://${HOST}:4321/myfw/myObject); echo ${resp} | jq