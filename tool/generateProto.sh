#!/bin/bash

protoc  --proto_path=../proto/define --go_out=../proto ../proto/define/*.proto