# slow-lang interpreter

a slow, fragile and simply designed interpreter for the programming language slow-lang.

Its only purpose is to teach myself more about interpreted programming languages and the structure of interpreters itself.
All work is based on Thorsten Ball's Book "Writing an interpreter in go".
You can find it [here](https://interpreterbook.com/).

## start interpreter
- `make build` (if not already done)
- `./slow-lang`
- enter commands like `print(113 * 2)`

## execute file
- `make build` (if not already done)
- `./slow-lang -src=./example.sl`
