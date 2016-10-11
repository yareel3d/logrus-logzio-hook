# logrus-logzio-hook

Ship logs to Logz.io over HTTP.

**Note**: This is a PoC and the API probably is going to be changed.

## Usage:

```go
fields := logrus.Fields{
    "ID": "89fabde223",
    "Host": os.Getenv("HOST"),
    "Username": os.Getenv("USER"),
}
tr := &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
httpsClient := &http.Client{Transport: tr}

hook := logzio.New(os.Getenv("LOGZ_HOST"), "CuantoQuedaBot", fields)
hook.SetClient(httpsClient)
...
logrus.AddHook(hook)
```
