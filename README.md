Working the example in https://codegangsta.gitbooks.io/building-web-apps-with-go/content/index.html

The code is using out dated libraries and commands

Problems
- go mod
-- when I tried to initialize a module, I kept getting an error. Issuing the following helped. 
--- export GO111MODULE=on
- blackfriday library
-- There is a new version of the library out. 
--- changed the code from using .MarkdownCommon (v1) to .Run (v2)
-- Got an error about not being able to load the package
--- error message: build markdownapp: cannot load gopkg.in/russross/blackfriday.v2: cannot find module providing package gopkg.in/russross/blackfriday.v2
--- resolution: go mod edit -replace=gopkg.in/russross/blackfriday.v2@v2.0.1=github.com/russross/blackfriday/v2@v2.0.1
