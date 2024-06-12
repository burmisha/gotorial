Всё подряд про Go

* https://www.youtube.com/watch?v=1MXIGYrMk80
* https://stepik.org/course/96832
* https://go.dev/ref/spec
* https://go.dev/wiki/CodeReviewComments


```sh
for task in task15{1..2} task16{1..2} task17{1..3}; do
  if ! [ -f tasks/${task}/main.go ]; then
    mkdir -p tasks/${task}
    echo "// echo 'abc' | go run tasks/${task}/main.go" > tasks/${task}/main.go
  else
    echo "${task} exists"
  fi
done
```