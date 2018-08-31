# Prepare env

* godep
```
$ cd $GOPATH/bin
$ wget https://github.com/tools/godep/releases/download/v80/godep_linux_amd64
```

* make a directory for GOPATH
```
# 
$ mkdir -p gopath
$ cd gopath
$ ln -s ../../../../../../.. src/github.com/koketani/playground/go/tools/godep
```

* move to development path
```
$ cd gopath
$ export GOPATH=`pwd`
$ export PATH=$PATH:$GOPATH/bin
$ cd $GOPATH/src/github.com/koketani/playground/go/tools/godep
```

# godep commands

* list and copy dependencies into Godeps
  * `$GOPATH/src` -> `Godeps.json`, `vendor`
  * this means it would make `vendor` directory if need
```
$ godep save
```

* check out listed dependency versions in GOPATH
  * `Godeps.json`, `vendor` -> `$GOPATH/src`
  * if missing ones in `vendor`, fetch from internet (really?)
```
$ godep restore
```

# Questions to be cleared

* referenced packages b/w go list and run differ?
  * even if there is no lib under `GOPATH`, exsitence `vendor` enables us go run.
  * go list outcome wouldn't show its directory. but, go list debug option includes it as deps and specify `vendor` directory.
* actually, removing vendor directory, go list commands also crash as well as go run command.

```
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master<>)$ go list ./...
github.com/koketani/playground/go/tools/godep
github.com/koketani/playground/go/tools/godep/gopath/src/github.com/mitchellh/mapstructure
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master<>)$ go run main.go
{person koke}
main.Person
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master<>)$ go list ./...
github.com/koketani/playground/go/tools/godep
github.com/koketani/playground/go/tools/godep/gopath/src/github.com/mitchellh/mapstructure
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master<>)$ rm -rf gopath/src/github.com/mitchellh
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master<>)$ go list ./...
github.com/koketani/playground/go/tools/godep
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master<>)$ go run main.go
{person koke}
main.Person
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master<>)$ go list --json
{
        "Dir": "/home/koketani/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep",
        "ImportPath": "github.com/koketani/playground/go/tools/godep",
        "Name": "main",
        "Target": "/home/koketani/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/bin/godep",
        "Root": "/home/koketani/Developments/git/github.com/koketani/playground/go/tools/godep/gopath",
        "Stale": true,
        "StaleReason": "target missing",
        "GoFiles": [
                "main.go"
        ],
        "Imports": [
                "fmt",
                "github.com/koketani/playground/go/tools/godep/vendor/github.com/mitchellh/mapstructure",
                "reflect"
        ],
        "Deps": [
                "bytes",
                "encoding",
                "encoding/base64",
                "encoding/binary",
                "encoding/json",
                "errors",
                "fmt",
                "github.com/koketani/playground/go/tools/godep/vendor/github.com/mitchellh/mapstructure",
                "internal/cpu",
                "internal/poll",
                "internal/race",
                "internal/testlog",
                "io",
                "math",
                "os",
                "reflect",
                "runtime",
                "runtime/internal/atomic",
                "runtime/internal/sys",
                "sort",
                "strconv",
                "strings",
                "sync",
                "sync/atomic",
                "syscall",
                "time",
                "unicode",
                "unicode/utf16",
                "unicode/utf8",
                "unsafe"
        ]
}
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master<>)$ rm -rf vendor
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master *<>)$ go list --json
{
        "Dir": "/home/koketani/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep",
        "ImportPath": "github.com/koketani/playground/go/tools/godep",
        "Name": "main",
        "Target": "/home/koketani/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/bin/godep",
        "Root": "/home/koketani/Developments/git/github.com/koketani/playground/go/tools/godep/gopath",
        "Stale": true,
        "StaleReason": "target missing",
        "GoFiles": [
                "main.go"
        ],
        "Imports": [
                "fmt",
                "github.com/mitchellh/mapstructure",
                "reflect"
        ],
        "Deps": [
                "errors",
                "fmt",
                "github.com/mitchellh/mapstructure",
                "internal/cpu",
                "internal/poll",
                "internal/race",
                "internal/testlog",
                "io",
                "math",
                "os",
                "reflect",
                "runtime",
                "runtime/internal/atomic",
                "runtime/internal/sys",
                "strconv",
                "sync",
                "sync/atomic",
                "syscall",
                "time",
                "unicode",
                "unicode/utf8",
                "unsafe"
        ],
        "Incomplete": true,
        "DepsErrors": [
                {
                        "ImportStack": [
                                "github.com/koketani/playground/go/tools/godep",
                                "github.com/mitchellh/mapstructure"
                        ],
                        "Pos": "main.go:7:2",
                        "Err": "cannot find package \"github.com/mitchellh/mapstructure\" in any of:\n\t/home/koketani/.gvm/gos/go1.10/src/github.com/mitchellh/mapstructure (from $GOROOT)\n\t/home/koketani/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/mitchellh/mapstructure (from $GOPATH)"
                }
        ]
}
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master *<>)$ go run main.go
main.go:7:2: cannot find package "github.com/mitchellh/mapstructure" in any of:
        /home/koketani/.gvm/gos/go1.10/src/github.com/mitchellh/mapstructure (from $GOROOT)
        /home/koketani/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/mitchellh/mapstructure (from $GOPATH)
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master *<>)$
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep/gopath/src/github.com/koketani/playground/go/tools/godep (master *<>)$ go list
github.com/koketani/playground/go/tools/godep
```

# Appendix

* godep help
```
koketani:~/Developments/git/github.com/koketani/playground/go/tools/godep (master *+=)$ godep
Godep is a tool for managing Go package dependencies.

Usage:

        godep command [arguments]

The commands are:

    save     list and copy dependencies into Godeps
    go       run the go tool with saved dependencies
    get      download and install packages with specified dependencies
    path     print GOPATH for dependency code
    restore  check out listed dependency versions in GOPATH
    update   update selected packages or the go version
    diff     shows the diff between current and previously saved set of dependencies
    version  show version info

Use "godep help [command]" for more information about a command.
```
