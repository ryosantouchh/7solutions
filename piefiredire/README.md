## Instruction

### Challenge #1

- find the most value in all possible paths
- using tree1.json , tree2.json and tree3.json
- you can change how the function read json file at `func readJSONFile` to read specific file

```plain
cd mostValuePath
go run main.go
```

---

---

### Challenge #2

- decode from text to number

```plain
cd catchMeIfYouCan
go run main.go
```

---

---

### Challenge #3

#### HTTP (Gin Framework)

##### Run the server

```plain
cd piefiredire/cmd/http
go run main.go
```

`GET localhost:8080/beef/summary`

```json
## response example ::
{
    "beef": {
        "loin": 1,
        "sirloin": 5,
        "bacon": 10,
    }
}
```

##### Run the test

```plain
cd piefiredire
go test ./...
```

---

#### gRPC

- using grpcurl to call the server (MacOS)

```plain
brew install grpcurl
```

```plain
cd piefiredire/cmd/grpc
go run main.go
```

- open another terminal to run grpcurl

```plain
grpcurl --plaintext localhost:50051 BeefService/GetSummary
```

```json
## response example ::
{
    "beef": {
        "loin": 1,
        "sirloin": 5,
        "bacon": 10,
    }
}
```

---

Big PS : I want to thank you for the challenges. They make me to know that I have much more things to learn along the programming journey, and your challenges give me motivation and inspiration to be the better programmer.
