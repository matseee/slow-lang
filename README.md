# slow-lang interpreter

a slow, fragile and simply designed interpreter for the programming language slow-lang.

Its only purpose is to teach myself more about interpreted programming languages and the structure of interpreters itself.
All work is based on Thorsten Ball's Book "Writing an interpreter in go".
You can find it [here](https://interpreterbook.com/).

## documentation

### start interpreter
- `make build` (if not already done)
- `./slow-lang`
- enter commands like `print(113 * 2)`

### execute file
- `make build` (if not already done)
- `./slow-lang -src=./example.sl`

### slow-lang syntax

#### types and variables
```
set a = true        // Booleans
set x = 12          // Integers
set y = "slow-lang" // Strings
```

#### operators
```
// prefix operators
-5;
!true;
!false;

// infix (or binary) arithmetic operators
5 + 5;
"Hello" + " " + "World";
5 - 5;
5 / 5;
5 * 5;

// infix comparison operators
foo == bar;
foo != bar;
foo < bar;
foo > bar;
```

#### if-else-expressions
```
if (x > y) { 
    // ...
} else {
    // ...
}
```

#### functions
```
set min = fn(x,y) {
    if (x < y) {
        return x;
    }

    // last statement is always interpreted as a return value, so 
    y;
    // is the same as 
    // return y;
}

min(10, 5); // => 5

set greeter = fn(x) {
    fn(y) {
        return x + " " + y;
    }
};

set helloGreeter = greeter("Hello")
helloGreeter("World");  // => "Hello World"
```

#### arrays
```
set arr = [1, 2, 3];
arr[0];                 // => 1

// builtin array functions:
len(arr);               // => 3
first(arr);             // => 1
last(arr);              // => 3
tail(arr);              // => [2, 3]
append(arr, 4);         // => [1, 2, 3, 4]

set anotherArr = [-13, "aString", fn(x) { x }];
```
