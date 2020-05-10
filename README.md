# test-blockchain

```
github.com/luisarmando-testcoder/test-blockchain
```

Testing a blockchain in a RESTAPI, with Mux in Golang

## Expected Requests

```
GET /blockchain

GET /block/{hash}

POST /block
body raw {
    data string
}
```