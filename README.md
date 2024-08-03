# Go  

## [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)  

```bash
cd /Users/<your username>/go
# location of go install at
mkdir go
# create new directory name: `go`
# `bin` `go` `pkg`
cd go
# navigate to go folder
mkdir hello
# create new directory name: `hello`
cd hello
# navigate to hello folder
```  
```plaintext
go
├── ./hello
│   ├── go.mod
│   ├── go.sum
│   └── hello.go
└── README.md
``` 
### 1. Enable dependency tracking for your code.  
Run the `go mod init command`, giving it your module path -- here, use ***example.com/greetings***. If you publish a module, this must be a path from which your module can be downloaded by Go tools. That would be your code's repository. 
```bash
go mod init example.com/hello
```  
> go: creating new go.mod: module example.com/hello  

### 2. In your text editor, create a file `hello.go` in which to write your code.  
### 3. Paste the following code into your hello.go file and save the file.  
```bash
package main

import (
	"fmt"
	"rsc.io/quote"
    # external package
)

func main() {
	fmt.Println("- Hello World!")
	fmt.Println("- " + quote.Hello())
	fmt.Println("- " + quote.Glass())
	fmt.Println("- " +  quote.Go())
	fmt.Println("- " + quote.Opt())
}
```  
- External package [rsc.io/quote](https://pkg.go.dev/rsc.io/quote)  

### 4. Add new module requirements and sums.  
```bash
go mod tidy
```  
> go: finding module for package rsc.io/quote  
> found rsc.io/quote in rsc.io/quote v1.5.2  

#### **When to use** ***`go mod tidy`***  
- **After Adding New Imports:** Whenever you add new imports to your Go files, running `go mod tidy` can ensure that all necessary dependencies are included in your `go.mod` file.  
- **Before Committing Changes:** It's a good practice to run `go mod tidy` before committing your changes to keep your module dependencies clean and up-to-date.  
- **Cleaning Up the Module:** If you suspect that your go.mod file has become cluttered with unused dependencies, running `go mod tidy` can help clean it up.

### 5. Run your code to see the greeting.  
```bash
go run .
```   

```bash
# Println()
hello % go run .
- Hello World!
- Hello, world.
- I can eat glass and it doesn't hurt me.
- Don't communicate by sharing memory, share memory by communicating.
- If a program is too slow, it must have a loop.
```  
---  

## [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)  

##  **Use your local package** 
- In the `hello` directory that will use ***local package***, run the following command
```bash
go mod edit -replace example.com/greetings=../greetings
```  
The command specifies that `example.com/greetings` should be replaced with `../greetings` for the purpose of locating the dependency. After you run the command, the `go.mod` file in the hello directory should include a ***replace directive***:   

```go
// go.mod
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings
```  

- From the command prompt in the **hello** directory, run the `go mod tidy` command to synchronize the ***example.com/hello*** module's dependencies, adding those required by the code, but not yet tracked in the module.  

```bash
hello % go mod tidy
go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000 
```  

- After the command completes, the ***example.com/hello*** module's go.mod file should look like this:  

```go
// go.mod
module example.com/hello

go 1.22.5

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
```  

- The command found the local code in the greetings directory, then added a require directive to specify that ***example.com/hello*** requires ***example.com/greetings***. You created this dependency when you imported the greetings package in **hello.go**.  

- To reference a published module, a ***go.mod*** file would typically omit the replace directive and use a require directive with a tagged version number at the end.  


## Automate your documentation  
### 📺 📺 📺 [Godocs - Effortless documentation for your go packages](https://www.youtube.com/watch?v=80VT3xexcWs)  

> Use : `godoc`  

- Install  
```bash
go install golang.org/x/tools/cmd/godoc@latest
```  

Kept in `/Users/<your-username>/go/bin`  

- Config
```bash
# In: ~/.zshrc  
export PATH=$PATH:$(go env GOPATH)/bin
```  

## [Developing and publishing modules](https://go.dev/doc/modules/developing)  


## [Publishing a module](https://go.dev/doc/modules/publishing)  


