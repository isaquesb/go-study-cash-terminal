#!/bin/sh

go test -v -coverprofile=./profile.cov
test_result=$?
go tool cover -func ./profile.cov
exit $test_result
