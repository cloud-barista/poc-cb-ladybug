#!/bin/bash

HOST=${1-"localhost"}

resp=$(curl -sX GET http://${HOST}:4321/myfw/myObject/1); echo ${resp} | jq