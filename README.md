<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/base" width="720"></p>

# base

[![ci](https://github.com/go-composites/base/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/base/actions/workflows/ci.yml)

The reflective base of [go-composites](https://github.com/go-composites). `Base`
gives every composite a common, self-describing root: it answers what kind of
interface a value is, whether it responds to a named method, and which methods
it exposes — all via reflection. Embed `Base.Interface` (or reuse its package
functions) so a composite can introspect itself without hand-written
boilerplate.

## Install

```sh
go get github.com/go-composites/base
```

## API

`Base.Interface` is:

| method | returns | notes |
| --- | --- | --- |
| `Kind()` | `string` | the interface kind, e.g. `Base.Interface` (package name + `.Interface`) |
| `RespondTo(name)` | `bool` | whether the value has an exported method named `name` |
| `Methods()` | `[]string` | the value's method names, via reflection |

`New()` returns a bare `Base.Interface`. Each method also has a package-level
twin that operates on any value:

- `Kind(instance interface{}) string`
- `RespondTo(instance interface{}, methodName string) bool`
- `Methods(instance interface{}) []string`

These are what an embedding type delegates to when it implements the interface.

## Usage

```go
package Vehicle

import (
	Base "github.com/go-composites/base/src"
)

type Interface interface {
	Base.Interface
	Start() bool
}

type data struct{}

func New() Interface { return &data{} }

func (d data) Kind() string             { return Base.Kind(d) }
func (d data) RespondTo(n string) bool  { return Base.RespondTo(d, n) }
func (d data) Methods() []string        { return Base.Methods(d) }
func (d data) Start() bool              { return true }
```

```go
v := Vehicle.New()
v.Kind()            // "Vehicle.Interface"
v.RespondTo("Start") // true
v.Methods()         // ["Kind" "Methods" "RespondTo" "Start"]
```

## License

BSD-3-Clause © the go-composites/base authors.
