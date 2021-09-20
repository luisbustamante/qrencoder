# qrencoder
Encode any message to QR code image

## Get started

This is a [fiber](https://github.com/gofiber/fiber) based project. You just need to clone this repo and install dependendies

```
$ git clone https://github.com/luisbustamante/qrencoder.git
$ cd qrencoder
$ go install
```

Start the project and check http://localhost:3000

```
$ go run .

 ┌───────────────────────────────────────────────────┐
 │                    Fiber v2.3.0                   │
 │               http://127.0.0.1:3000               │
 │                                                   │
 │ Handlers ............. 3  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............. 72619 │
 └───────────────────────────────────────────────────┘


```

That's it! Now make a request to generate a QR image from text

```
curl -X GET "http://localhost:3000/generate?text=test" --output test.png
open test.png
```
You should see an image like this one:
![test](https://user-images.githubusercontent.com/10065321/134004371-96a09b23-3e81-4f49-9b93-64ac45fef46d.png)


Enjoy :) 
