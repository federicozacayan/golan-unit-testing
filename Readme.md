# Teting (windows instructions)

To run test gouped by TAGS. 

For example `unit` tag

```BASH
clear ; go test .\...  -v --tags=unit -cover
```

For example `integration` tag
```BASH
clear ; go test .\... -v --tags=integration -cover
```


## To run all testing

Delete from `math1\calculator_test.go`
```
//go:build unit
// +build unit
```

And delete from `math1\calculator_integration_test.go` also.
```
//go:build integration
// +build integration
```

```BASH
clear ; go test .\... -v -cover
```


### More references

```
https://stackoverflow.com/questions/16353016/how-to-go-test-all-tests-in-my-project
```