# Klasifikasi for Golang

Official [Klasifikasi](https://klasifikasi.com/) API Client Library

## Requirements
Go v1.13
## Installation
install klasifikasi-go with:

`go get -u github.com/bahasa-ai/klasifikasi-go`

then, you can import it like this
```go
import (
	klasifikasi "github.com/bahasa-ai/klasifikasi-go"
)
```

## Quick Start

You will need valid `clientId` & `clientSecret` of your model. You can get those
from credential section at your model page, which is both unique per model.

```go
clientCredentials := []klasifikasi.ClientBuildParams{
  {
    ClientId:     "client-id",
    ClientSecret: "client-secret"
  }
}
klasifikasiInstance := klasifikasi.Build(clientCredentials)

```
You can pass multiple `clientId` & `clientSecret` too

```go
clientCredentials := []klasifikasi.ClientBuildParams{
  {
    ClientId:     "client-id-1",
    ClientSecret: "client-secret-1",
  },
  {
    ClientId:     "client-id-2",
    ClientSecret: "client-secret-2", 
  }
}
klasifikasiInstance := klasifikasi.Build(clientCredentials)

```

## Classify
You will need you model `publicId` to start classifying with your model. You can get your model `publicId` from your model page, or you can get it from here :
```go
models := klasifikasiInstance.GetModels()
for publicId := range models {
  fmt.Println(publicId)
}
```

Classifying example
```go
result, _ := klasifikasiInstance.Classify("publicId", "your query")
for _, labelResult := range result.Result {
  fmt.Printf("%+v\n", labelResult)
  // Output example
  // {
  //   Label:Sedih
  //   Score:0.94172746
  // }
}
```

## Logs
You can get your classifying logs based on your model `publicId`
```go

startedAt,_ := time.Parse("January 2 2006", "December 1 2020")
endedAt, _ := time.Parse("January 2 2006", "December 4 2020")

logs, _ := klasifikasiInstance.Logs("publicId", klasifikasi.LogsParams{
  StartedAt: startedAt,
  EndedAt:   endedAt,
  Take:      10,
  Skip:      0,
})
for _, data := range logs.Logs {
  fmt.Printf("%+v\n", data)
  // Output example
  // {
  //   CreatedAt:2020-12-03T15:36:18+07:00
  //   UpdatedAt:2020-12-03T15:36:18+07:00
  //   Id:1348
  //   Query:sucksee
  //    ModelResult:[
  //     {
  //        Label:Sedih 
  //        Score:0.94172746
  //     }
  //     {
  //       Label:Bahagia
  //       Score:0.05827257
  //     }
  //   ]
  // }
}
```

## Error
`Classify` & `Logs` function will return an error if something bad happen, always check the error variable
```go
result, err := klasifikasiInstance.Classify("publicId", "your query")
if err != nil {
  log.Fatal(err)
}
```

```go
logs, err := klasifikasiInstance.Logs("publicId", klasifikasi.LogsParams{
  StartedAt: startedAt,
  EndedAt:   endedAt,
  Take:      10,
  Skip:      0,
})
if err != nil {
  log.Fatal(err)
}
```