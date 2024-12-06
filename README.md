Attempt to use JS [JSONPath-Plus/JSONPath](https://github.com/JSONPath-Plus/JSONPath) npm package in Go.

I have checked 3 Go packages that can run JS:
* [rogchap/v8go](https://github.com/rogchap/v8go)
  * failed because of "Fatal javascript OOM in CALL_AND_RETRY_LAST" after 1000 call in a row
* [robertkrimen/otto](https://github.com/robertkrimen/otto)
  * don't support some JS keywords like `class`
* [dop251/goja](https://github.com/dop251/goja)
  * worked!
