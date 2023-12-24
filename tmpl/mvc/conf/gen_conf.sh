#! /usr/local/bin env bash

protoc -I. --go_out=paths=source_relative:. conf.proto