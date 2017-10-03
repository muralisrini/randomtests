# randomtests

export GOPATH=path to randomtests

### workerpool vs goroutines test
+ cd src/workertest
+ go build
+ ./workertest
   
vary N to see how growing N affects tests, with N=10000 we get
> ./workertest  
> Elapsed time using pool 12.69563ms  
> Elapsed time using goroutine 4.274119ms  
