# Singleton Pattern
The Singleton pattern is easy to remember. As the name implies, it will provide you with a single instance of an object, and guarantee that there are no duplicates.
At the first call to use the instance, it is created and then reused between all the parts in the application that need to use that particular behavior.


## The Example
As an example of an object called Counter of which we must make sure there is only one instance, this object will contain the methods to add and subtract a unit, as well as to get the value of the accumulator. It should not matter how many instances we have of the counter, it must always return a single instance and if it does not exist it is created, each time the counter is called it must count the same value and it must be consistent between instances.


## Unit Tests

The next commands should be execute inside the folder named devices.
```sh
cd devices/
```


* Execute a single unit test per function name
```sh
go test -v -run=TestGetCounter
```


* Execute all unit tests
```sh
go test -v
```


* Generate code coverage by unit tests metrics and show these in HTML report

```sh
go test -coverprofile=coverage.out
go tool cover -html=coverage.out        
```
* Code coverage with unit tests

![alt coverage](/images/code_coverage.png)