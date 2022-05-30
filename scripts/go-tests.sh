#!/usr/bin/env bash
set -e

cd ./Hello-World
go test -v

cd ../Integers
go test -v

cd ../Iteration
go test -v
go test -bench=.

cd ..
