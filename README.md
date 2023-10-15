Picking a Docker image
https://pythonspeed.com/articles/base-image-python-docker-images/

Using Go in Docker
https://www.docker.com/blog/developing-go-apps-docker/

Reading files
https://gobyexample.com/reading-files

Reading text files
https://www.jeremymorgan.com/tutorials/go/how-to-read-text-file-go/

Download file in Golang
https://golangcode.com/download-a-file-from-a-url/

String methods
https://gobyexample.com/string-functions

Golang info
- exported names begin with capitals

Go maps
https://go.dev/blog/maps

Buffered vs unbuffered IO
https://www.reddit.com/r/golang/comments/15r962b/when_should_you_consider_using_bufio/
https://stackoverflow.com/questions/1450551/buffered-vs-unbuffered-io#1450563

Create an IO Reader from file
https://stackoverflow.com/questions/25677235/create-a-io-reader-from-a-local-file


Concurrent reading of file?
https://stackoverflow.cohm/questions/11867348/concurrent-reading-of-a-file-java-preferred

Packages Go
https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules

Capitalization case
https://nlp.stanford.edu/IR-book/html/htmledition/capitalizationcase-folding-1.html

Error handling
https://earthly.dev/blog/golang-errors/
https://go.dev/blog/defer-panic-and-recover

Returning multiple values
https://gobyexample.com/multiple-return-values

Testing in Go
https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/#the-testify-package

Test setup in Go
https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go

Using Subtests and Sub-benchmarks
https://go.dev/blog/subtests
https://blog.logrocket.com/benchmarking-golang-improve-function-performance/

Concurrent programming
https://www.freecodecamp.org/news/concurrent-programming-in-go/
https://golangbot.com/goroutines/
https://gobyexample.com/waitgroups

Error handling
https://earthly.dev/blog/golang-errors/
https://preslav.me/2023/04/14/golang-error-handling-is-a-form-of-storytelling/

Package naming
https://go.dev/doc/effective_go#package-names
https://stackoverflow.com/questions/61845013/package-xxx-is-not-in-goroot-when-building-a-go-project

Iterating through map in order
https://stackoverflow.com/questions/37695209/golang-sort-slice-ascending-or-descending


Commands
docker build --rm -t [YOUR IMAGE NAME]:alpha .

docker build --rm -t go_image:alpha .

docker run -d -it --name go-docker-app --mount type=bind,source="$(pwd)"/src,target=/app go_image:alpha


TODO
- use ErrorGroup for error handling
- use temporary files
- figure out way to change output
- write to CSV



- use a porter stemmer? (not important)
