#!/bin/bash

if [ -z "$ENV" ]; then
    echo "ENV variable not found. Please set it and try again"
    exit 1
fi

if [ "$ENV" != "DEV" ] && [ "$ENV" != "PROD" ]; then
    echo "Valid values are: DEV or PROD. Got: $ENV"
    exit 1
fi

export GOTMPDIR="$JENKINS_HOME/$ENV"
mkdir -p $GOTMPDIR
echo "*** Building the app ..."
go build -o $GOTMPDIR/$ENV main.go
