# randomtests

export GOPATH=path to randomtests

### workerpool vs goroutines test
+ cd src/workertest
+ go build
+ ./workertest
   
vary N to see how growing N affects tests, with N=10000 we get
> ./workertest  
> Elapsed time using async pool 12.95567ms  
> Elapsed time using sync pool 8.613112ms  
> Elapsed time using goroutine 4.752551ms  
