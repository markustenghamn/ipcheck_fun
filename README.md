# IPcheck.fun

A simple website for checking your IP written in Golang.

See the project live at https://ipcheck.fun

## Run for development

Use `go-bindata` as described below and then type `go run main.go app.go api.go bindata.go` and go to `localhost:8009` in your browser

## Run for production

Use go-bindata to include the template files in your binary

```
go-bindata assets/... views/...
```

Then build

```
go build
```

And then simply run using `./bcrypt_fun` and go to `localhost:8009` in your browser

The following arguments are available


```
-siteurl - The site url
-sitename - The name of the site
-host - The host and port (localhost:8009 or :80)
```

## Contributors

Created by [Markus Tenghamn](https://ma.rkus.io)

Other contributors welcome!

