Всё подряд про Go

```sh
for task in task13{1..3}; do
  mkdir -p tasks/${task}
  echo "// echo abc | go run tasks/${task}/main.go" >> tasks/${task}/main.go
done
```