#!/bin/bash
godep restore
go install
cp ./pipemon_database.yml $GOPATH/bin/pipemon_database.yml
