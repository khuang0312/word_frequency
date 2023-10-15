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
https://stackoverflow.com/questions/55442878/organize-local-code-in-packages-using-go-modules

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

Cannot kill docker container
https://javahowtos.com/guides/124-docker/414-solved-cannot-kill-docker-container-permission-denied.html

Unit testing that is opening and reading a file
https://stackoverflow.com/questions/56397240/unit-test-function-that-is-opening-and-reading-a-file

Logging in Golang
https://betterstack.com/community/guides/logging/logging-in-go/

CLI apps
https://www.reddit.com/r/golang/comments/seg2sx/recommended_frameworklibrary_for_creating_cli/

File upload
https://freshman.tech/file-upload-golang/

CHMOD perms
https://chmodcommand.com/chmod-660/

URL validation
https://stackoverflow.com/questions/31480710/validate-url-with-standard-package-in-go

Defers
https://www.joeshaw.org/dont-defer-close-on-writable-files/

Commands
docker build --rm -t [YOUR IMAGE NAME]:alpha .

docker build --rm -t go_image:alpha .

docker run -d -it --name go-docker-app --mount type=bind,source="$(pwd)",target=/app go_image:alpha

go test -bench=. -count 5 -run=^#


TODO
- cleanup README
- add unit tests or integration tests for fileutils
- getting benchmarking to display in friendlier units
- add logging using Zap, slog, or Golang log
- add Jenkins?
- make a REST API?
- make a CLI?


- use a porter stemmer? (not important)


In future (Because these are still experimental)
- Use ErrorGroup
- Use maps.Keys


No GRPC -> because protobuffs aren't ideal for sending large message sizes ( > 1MB)
