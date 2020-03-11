# Note for Golang Modules

* concept
  * compared to python
    * module is a python file
    * package is a collection of modules
    * import module and access module var/class/def
    * pakcage is provided as a part of library via PyPI/deb (wheel/egg.info)
    * two types of importing module module can be imported from 

* commands

  * go mod 
```shell-session

$ go mod
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

        go mod <command> [arguments]

The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph
        init        initialize new module in current directory
        tidy        add missing and remove unused modules
        vendor      make vendored copy of dependencies
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```

  * go mod init

```shell-session
$ go mod init github.com/koketani/playground/go/golang/modules

$ cat go.mod
module github.com/koketani/playground/go/golang/modules

go 1.14
$ cat <<EOF > test
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
}
EOF

$ cat go.mod
module github.com/koketani/playground/go/golang/modules

go 1.14

require rsc.io/quote v1.5.2

$ cat go.sum
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=

$ go mod why
# github.com/koketani/playground/go/golang/modules
github.com/koketani/playground/go/golang/modules
$ go mod graph
github.com/koketani/playground/go/golang/modules rsc.io/quote@v1.5.2
rsc.io/quote@v1.5.2 rsc.io/sampler@v1.3.0
rsc.io/sampler@v1.3.0 golang.org/x/text@v0.0.0-20170915032832-14c0d48ead0c

$ go mod verify
all modules verified
```


