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
httpsClient := &http.Client{Transport: &http.Transport{TLSClientConfig: {InsecureSkipVerify: true}}}

hook := logrus_logstash.NewHookWithFields("https://logz.io:9891", "MyApp", fields)
hook.SetClient(httpsClient)
...
logrus.AddHook(hook)
```
