calculate pi using gopherjs is faster than native go!!
http://www.gopherjs.org/blog/2015/09/28/surprises-in-gopherjs-performance/

* `go get -u github.com/gopherjs/gopherjs` need go 1.7 for gopherjs 1.7
* `go run main.go` get 8 sec
* `gopherjs build` then `node pi.js` got 3 sec
