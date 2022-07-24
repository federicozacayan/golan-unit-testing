# To run all testing (windows instructions)

Run all test by using all tags

```
clear ; go test .\...  -v --tags "unit integration" -cover
```

## Testing by tags

To run test gouped by TAGS. 

For example `unit` tag

```BASH
clear ; go test .\...  -v --tags=unit -cover
```

For example `integration` tag
```BASH
clear ; go test .\... -v --tags=integration -cover
```




### More references

```
https://stackoverflow.com/questions/16353016/how-to-go-test-all-tests-in-my-project
```