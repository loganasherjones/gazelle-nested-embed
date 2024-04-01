# Broken Nested Embed

This project demonstrates the broken behavior of gazelle in reference to
embed directives that reference nested data files in golang.

## Steps to Reproduce

```
bazel build //...
```

This will error, while the following succeeds:

```
go build main.go
```