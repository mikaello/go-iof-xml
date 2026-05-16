# go-iof-xml

[![Go CI](https://github.com/mikaello/go-iof-xml/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/mikaello/go-iof-xml/actions/workflows/build-and-test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/mikaello/go-iof-xml.svg)](https://pkg.go.dev/github.com/mikaello/go-iof-xml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikaello/go-iof-xml)](https://goreportcard.com/report/github.com/mikaello/go-iof-xml)

Go bindings for the [IOF Interface Standard v3](https://github.com/international-orienteering-federation/datastandard-v3) XML format used by orienteering software.

- `pkg/iof_v3` — Go structs generated from the official v3 XSD.
- `pkg/marshallers` — small helpers for decoding and encoding v3 documents.

## Install

```sh
go get github.com/mikaello/go-iof-xml
```

## Usage

Decode a document when you know its type:

```go
package main

import (
    "fmt"
    "os"

    "github.com/mikaello/go-iof-xml/pkg/iof_v3"
    "github.com/mikaello/go-iof-xml/pkg/marshallers"
)

func main() {
    data, err := os.ReadFile("ResultList.xml")
    if err != nil {
        panic(err)
    }

    list, err := marshallers.DecodeAs[iof_v3.ResultList](data)
    if err != nil {
        panic(err)
    }

    fmt.Println(list.Event.Name)
}
```

Decode a document when you don't:

```go
doc, err := marshallers.Decode(data)
if err != nil {
    panic(err)
}

switch d := doc.(type) {
case *iof_v3.ResultList:
    fmt.Println("got result list:", d.Event.Name)
case *iof_v3.StartList:
    fmt.Println("got start list:", d.Event.Name)
}
```

Encode back to XML or JSON:

```go
xmlBytes, err := marshallers.EncodeXML(list)
jsonBytes, err := marshallers.EncodeJSON(list)
```

## Related

- Java classes for IOF XML v3 and v2:
  [orienteering-oss/iof-xml](https://github.com/orienteering-oss/iof-xml)
- C# classes for IOF XML v3:
  [international-orienteering-federation/Dotnet-Client-IOF.XML.V3](https://github.com/international-orienteering-federation/Dotnet-Client-IOF.XML.V3)
- The official
  [v3 datastandard](https://github.com/international-orienteering-federation/datastandard-v3)
  repository.

## License

[GPL-3.0](./LICENSE)
