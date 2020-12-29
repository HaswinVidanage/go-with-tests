### Structs, methods & interfaces
Source : 
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#wrapping-up



This was more TDD practice, iterating over our solutions to basic mathematic problems and learning new language features motivated by our tests.

Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer
Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)

Adding methods so you can add functionality to your data types and so you can implement interfaces

Table based tests to make your assertions clearer and your suites easier to extend & maintain

This was an important chapter because we are now starting to define our own types. In statically typed languages like Go, being able to design your own types is essential for building software that is easy to understand, to piece together and to test.

Interfaces are a great tool for hiding complexity away from other parts of the system. In our case our test helper code did not need to know the exact shape it was asserting on, only how to "ask" for it's area.

As you become more familiar with Go you start to see the real strength of interfaces and the standard library. 

You'll learn about interfaces defined in the standard library that are used everywhere and by implementing them against your own types you can very quickly re-use a lot of great functionality.