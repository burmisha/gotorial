#!/usr/bin/env bash

set -ue -o posix -o pipefail

for task in task15{1..2} task16{1..2} task17{1..3} task22{1,3,4}; do
  if ! [ -f tasks/${task}/main.go ]; then
    mkdir -p tasks/${task}
    echo "// echo 'abc' | go run tasks/${task}/main.go" > tasks/${task}/main.go
  else
    echo "${task} already exists"
  fi
done
