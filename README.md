
# GOLANG (GIN) DEMO API
This is a golang app to show gin framework development & testing.
- Author: [**Cristian Pardo**](http://cristianpardodeveloper.net) (cristianpardoluna@gmail.com)
- Based on: https://go.dev/doc/tutorial/web-service-gin

built with:
* go v1.18.1
    * gin v1.8.2

## Start
Start the development server with:
```bash
. start.sh
```

## Run tests
Run the available tests:
```go
go test 
```

## API
1. get records:
    ```bash
    curl 127.0.0.1:8080/records
    ```
2. post record
    ```bash
    curl -X POST 127.0.0.1:8080/records -d '{"id": RECORD_ID, "name": RECORD_NAME, "val": RECORD_VALUE}'
    ```
3. get record by ID
    ```bash
    curl 127.0.0.1:8080/RECORD_ID
    ```

## Key operations
* Get go version
    ```bash
    go version 
    ```
* Get go modules versions
    ```bash
    go list -m all
    
    # looking for specific module?
    go list -m all | egrep "gin"
    ```