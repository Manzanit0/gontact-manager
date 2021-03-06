# Gontact Manager

This project is a very simple contact manager application whose purpose is learn go.
I wrote it as I was getting aquainted with the language and as such I've tried both
to maintain the complexity as low as possible but at the same time, use as many features
as possible.

As for the application, it's an contact manager in the form of a HTTP server.
It spawns up a [`gin`](https://github.com/gin-gonic/gin) server which allows users to
add, edit, delete and list contacts in a generic addressbook.

## Getting started

The following command will run all the tests in the current directory and sub-directories:

```
go test ./...
```

To start the gin server, run:

```
go run main.go
```

## Notes to anyone learning Go

The project is structured in a main directory which contains the `main.go` file, which is
the entry point for the program and the `/addressbook` directory which contains a go module
with both the business logic and the gin server code.

As you can see, both `addressbook.go` and `web.go` are in the same package (`package addressbook`),
so there is no need to add imports when mentioning structs or functions of one file in the
other file.

The `go.mod` and `go.sum` files are the files which have all the information regarding the go
module. They track which version of both Go and all the libraries has been used so all builds
are deterministic. [Go modules](https://github.com/golang/go/wiki/Modules) is the alternative
to the single [`GOPATH` workspace pattern](https://golang.org/doc/code.html), introduced in Go 1.12.

### Regarding tests...

The project has been TDD-ed, so you can see how the different parts have been tested. 

To my surprise, the first thing I found when learning Go is that *assertions* are not common.
The pattern idiomatic go uses is to
[throw an error upon an unexpected result](https://golang.org/doc/code.html#Testing) in a test:

```go
if len(addressbook.Contacts) > 1 {
	t.Errorf("The addressbook has one too many contacts")
}
```

The second interesting thing I learnt is that what determines a test from a non-test is the
name: test functions are expected to start like `func TestXxx(*testing.T)`, camel-cased.
As long as they have that, they will be run by the test runner. If you come from Java,
it's like adding the JUnit `@Test` tag to a method.

```go
// web_test.go
func TestListContacts(t *testing.T) {
	// Some code
}
```


Furthermore, in order for the compiler to pick up test files and package them separately,
it must end with `_test.go`. As the documentation says,

> Test files that declare a package with the suffix "_test" will be compiled as a separate package, and then linked and run with the main test binary.

Apart from the naming of the tests, all test methods accept the parameter `t *testing.T`. It's
basically a structure with all the testing utilities you might need – It contains methods to
log messages to the console, to throw errors upon unexpected results, etc. For more information,
checkout the [documentation](https://golang.org/pkg/testing/#pkg-overview).

To implement before/after hooks, when logic starts to be duplicated, golang/testing provides
with the [main](https://golang.org/pkg/testing/#hdr-Main) function. It's a global hook, so it's
run once:

```go
func TestMain(m *testing.M) {
	// setup code
	retCode := m.Run()
	// teardown code
	os.Exit(retCode)
}
```

Lastly, on testing `gin` code, I have seen multiple approaches online. There were some people who
suggested implementing the handler functions with interfaces, to be able to stub the `*gin.Context`
structure, others which simply served mock requests... I went with the latter simply because I
found it to be simpler. As the code grows it could be interesting to look approaches like
[this](https://stackoverflow.com/questions/41742988/make-mock-gin-context-in-golang) as it's cleaner.
You will find the testing code under `addressbook/web_test.go`.
