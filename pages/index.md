---
layout: default
title: Gonfig
description: Go configuration without the ceremony.
---

# Gonfig

## Go configuration without the ceremony

Gonfig is a small, dependency-free Go configuration library for populating structs from environment variables, command-line arguments, and default values.

[Get started](#quick-start) · [API reference](#api-reference) · [GitHub](https://github.com/vrianta/gonfig)

> Current release: **v1.1.0** · Requires **Go 1.24.3 or later**

## Why Gonfig?

Define configuration once in a Go struct and let Gonfig resolve values from familiar sources:

- Environment variables for deployment
- Command-line arguments for runtime overrides
- Default values for sensible fallbacks
- Required tags for final validation
- Nested structs and pointers
- Generated command-line help

## Installation

```bash
go get github.com/vrianta/gonfig/v1@v1.1.0
```

Import the v1 package with the recommended alias:

```go
import gonfig "github.com/vrianta/gonfig/v1"
```

## Quick start

```go
package main

import (
    "fmt"

    gonfig "github.com/vrianta/gonfig/v1"
)

type Config struct {
    Host   string `env:"APP_HOST" arg:"host" default:"localhost"`
    Port   int    `env:"APP_PORT" default:"8080"`
    Debug  bool   `arg:"debug" default:"false"`
    APIKey string `env:"APP_API_KEY" required:"true"`
}

func main() {
    cfg := gonfig.New[Config](true)
    fmt.Println(cfg.Host, cfg.Port, cfg.Debug)
}
```

The `true` argument makes a missing required field panic. Use `Parse` with `false` when your application needs to inspect errors explicitly.

## How it works

For each exported, non-struct field, Gonfig resolves values in this order:

1. `env`
2. `arg`
3. `default`
4. `required` validation

The first successful source wins. Nested structs are processed recursively. Pointer fields are allocated when a value is supplied and remain nil when no value is supplied.

## Struct tags

| Tag | Purpose | Example |
| --- | --- | --- |
| `env` | Read an environment variable. | `` `env:"APP_HOST"` `` |
| `arg` | Read a command-line option. | `` `arg:"host"` `` |
| `default` | Provide a fallback value. | `` `default:"localhost"` `` |
| `required` | Require a non-zero value after parsing. | `` `required:"true"` `` |
| `description` | Add text to generated help output. | `` `description:"Server host"` `` |

### Environment variables

```go
Host    string        `env:"APP_HOST"`
Port    int           `env:"APP_PORT"`
Timeout time.Duration `env:"APP_TIMEOUT"`
```

Empty or unset environment variables are skipped.

### Command-line arguments

The `arg` tag uses the option name without dashes. All of these forms are supported:

```text
--port=8080
-port=8080
--port 8080
-port 8080
--debug
```

A flag without an explicit value is treated as the boolean value `true`.

### Required fields

The presence of the `required` tag enables validation. Both forms are accepted:

```go
APIKey string `env:"APP_API_KEY" required:"true"`
Token  string `required:""`
```

A missing required value returns an error, or panics when `crashOnFail` is `true`.

### Help descriptions

The `description` tag adds text to the help table. It does not populate or change the field.

```go
Host string `arg:"host" description:"Server host address"`
```

## Supported types

All value sources support the same scalar types:

- `string`
- `bool`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `time.Duration`, using values such as `30s`, `5m`, or `2h`
- Pointers to the supported scalar types, such as `*string`, `*int`, and `*bool`

Nested structs are processed recursively. Arrays, slices, maps, interfaces, channels, functions, and complex numbers are not supported as tagged field values.

## API reference

### `New`

```go
func New[T any](crashOnFail bool) T
```

Creates and populates a configuration value. When `crashOnFail` is `true`, a missing required field causes a panic. `New` does not return parsing errors; use `Parse` when errors must be inspected.

```go
cfg := gonfig.New[Config](false)
```

### `Parse`

```go
func Parse[T any](ctx *T, crashOnFail bool) (*T, error)
```

Populates an existing non-nil pointer to a struct and returns that pointer. It returns errors for invalid input, unexported fields, and missing required fields.

```go
cfg := &Config{}
if _, err := gonfig.Parse(cfg, false); err != nil {
    log.Fatal(err)
}
```

### `ParseOSArgs`

```go
func ParseOSArgs()
```

Scans `os.Args` and caches command-line arguments for later parsing. Most applications do not need to call it directly because parsing initializes the cache on demand.

## CLI help

Pass `-help` or `--help` to print a table containing field names, flags, environment variables, defaults, required status, and descriptions:

```bash
./app --help
```

## Error handling

Use `Parse` for explicit startup errors:

```go
cfg := &Config{}
if _, err := gonfig.Parse(cfg, false); err != nil {
    log.Fatal(err)
}
```

Use `New` with `true` for fail-fast initialization:

```go
cfg := gonfig.New[Config](true)
```

## License and source

Gonfig is maintained at [github.com/vrianta/gonfig](https://github.com/vrianta/gonfig). See the [v1.1.0 release](https://github.com/vrianta/gonfig/releases/tag/v1.1.0) for the current release notes.
