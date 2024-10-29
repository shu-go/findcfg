Package findcfg finds a config file.

[![](https://godoc.org/github.com/shu-go/findcfg?status.svg)](https://godoc.org/github.com/shu-go/findcfg)
[![Go Report Card](https://goreportcard.com/badge/github.com/shu-go/findcfg)](https://goreportcard.com/report/github.com/shu-go/findcfg)
![MIT License](https://img.shields.io/badge/License-MIT-blue)

# An example

```go
// finder for a YAML file in os.Executable() and os.UserConfigDir()+myapp
finder := findcfg.New( // config.yaml
    findcfg.Name("config"),
    findcfg.YAML(),
    findcfg.ExecutableDir(),
    findcfg.UserConfigDir("myapp"),
)

if found := finder.Find(); found != nil {
	return found.Path
}

return finder.FallbackPath()
```

# go get

```
go get -u github.com/shu-go/findcfg
```

----

Copyright 2024 Shuhei Kubota

<!--  vim: set et ft=markdown sts=4 sw=4 ts=4 tw=0 : -->
