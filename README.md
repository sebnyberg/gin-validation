# gin-validation

Example of custom Gin validation

Run the code:

```bash
go run main.go v8_to_v9.go fielderr.go
```

Try the errors with:

```bash
# this will fail on "red" because the `binding:"oneof=blue yellow" condition is not met
curl -G \
'http://localhost:5000/locations' \
-d name=volvo \
-d color=red
```
