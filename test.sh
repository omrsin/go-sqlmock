#!/bin/bash

RET=0
until [ ${RET} -eq 1 ]
do
    go test -count=1 ./...
    RET=$?
done





