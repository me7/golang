from https://plus.google.com/110008741064576119995/posts/63hquvmm6jn

I am new to golang. Employer send me a test exercise and turned me down without further explanations.
I'll appreciate if someone would tell me what's wrong with my code, there're only 110 lines of it. Here's a link with test description and implementation

https://gist.github.com/dmokhov/6539a2d920dec7e140ba02ea6f917ed8
* Your code seems fine, only you did use a global `MaxWorkers`. 
* The parm used for the http.Client timeout is silly, why use 1e9 instead of "time.Second", as a side note consider using a higher timeout as I get >1sec latency to golang.org. Finally the http client errors are not handled but ignored, you should log the network errors to stderr even though the test doesn't call for that.
* Not sure why the employer would fail you other than the global. It's also not a very good test, too easy :P
* The function names don't really convey their functionality (e.g. "get" does get a url but it also counts).
* You don't check whether the scanner terminates due to an error (check the documentation).


