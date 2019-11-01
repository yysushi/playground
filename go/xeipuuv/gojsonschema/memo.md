```
koketani:gojsonschema (master +=)$ go run main.go
true
#0 json vaildation succeed
#1 json vaildation succeed
#2 json vaildation failed with [value.1.id: Does not match format 'myidfmt'] on {"value":[{"name":"key1","value":true}, {"name":"key2","value":true,"id":"123"}]}
#3 json vaildation failed with [value.1.value: Invalid type. Expected: boolean, given: array] on {"value":[{"name":"key1","value":true}, {"name":"key2","value":[true]}]}
koketani:gojsonschema (master +=)$ 
```
