# Credential harvester

## What it is

This application presents a legitimate looking login webpage which captures credentials of users trying to log in and saves them into a local file.

## Demo

1. Run 

```
go run main.go
```

2. Go to [localhost:8081](http://localhost:8081) and enter some credentials.

3. You will see a log `Login attempt captured` in the console.

4. Inspect the file [credentials.txt](./credentials.txt)
