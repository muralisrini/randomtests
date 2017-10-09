# randomtests

export GOPATH=path to randomtests

### workerpool vs goroutines test
+ cd src/workertest
+ go build
+ ./workertest
   
vary N to see how growing N affects tests, with N=10000 we get
> ./workertest  
> Elapsed time using async pool 11.361809ms  
> Elapsed time using sync pool 7.422936ms  
> Elapsed time using goroutine 4.898461ms  
> Elapsed time using goroutine with sem and weight == Num CPU (8) -  10.965778ms  
> Elapsed time using goroutine with sem and full weight == all N (basically no blocking) 10000 -  5.192229ms  
