curl -X POST http://localhost:8080/db/mykey -H "Content-Type:application/octet-stream" -d "foo=bar"

curl -X GET http://localhost:8080/db/mykey

curl -X POST http://localhost:8080/db/dodo -H "Content-Type:application/octet-stream" -d $'"json":{"name":"Doittikorn","age":25}'

curl -X GET "http://localhost:8080/db/dodo?format=json"