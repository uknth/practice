# NPM Dependency Problem

There is NPM JSON API for getting NPM packages info. For example the following URL allows for getting information about the latest version of "forever" package:
http://registry.npmjs.org/forever/latest
This request will result in a JSON, containing many fields, including dependencies field

```
{
    "dependencies":{
        cliff: "~0.1.9",
        clone: "^1.0.2",
        colors: "~0.6.2",
        flatiron: "~0.4.2",
        forever-monitor: "~1.7.0",
        ...
    }    
}
```

This is a list of direct dependencies of an NPM package.

Write a function getAllDependencies(packageName) which takes in packageName parameter as a string and returns an array of strings of both direct and all indirect (recursive) dependencies of the given package, fetched from the API described above. For example, if A depends on B, and B depends on C and D, getAllDependencies('A') should return ['B', 'C', 'D']. The result should not contain duplicates.

In a correct implementation, getAllDependencies("forever") should return an array with length about 200+ (as of the time we wrote this question and might be different in the future).

Include a list of tools that needs to be installed to run your code and instructions on how to run your program.


**NOTE**
- Don't code solution in browser environment
- You shouldn't care about package version
- You shouldn't care about development dependencies
- Function should return the array, instead of printing the result

**Expectations**
- Code Correctness
- Code Readability
- Error Handling
