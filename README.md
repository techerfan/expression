# Expression

Expression is a basic expression parser and compiler, implemented in Golang for handling simple sytanxes such as mathematical operations (+, -, *, /, %), logical operations (==, <, >, <=, >=, !=, &&, ||), unary operations (-, !), and bitwise operations (>>, <<, ^, |, &).

## Installation

To use the package inside your golang project, simply run the following command:

```sh
go get github.com/techerfan/expression
```

## How to use?

To parse a string expression, first you need a to make a sytanx tree. During this process, lexer will tokenize the expression and makes tree for you:

```Golang
tree := syntax.Parse("1 + 2")
```

The result of the parsing, contains an array named `Diagnostics` that helps you to diagnose errors that happens during the process. So basically when the length of the `Diagnostics` array is greater than zero, the provided expression is not valid and has syntax issues. Otherwise, we are good to continue:

```Golang
if len(tree.Diagnostics) > 0 {
  for _, d := range tree.Diagnostics {
    fmt.Println(d.Message, d.Span, d.Span.Length, d.Span.Start)
  }
  panic(1)
}
```

If any variable is used in the expression, the parser will return them to us:

```Golang
// Imagine the expression was "variable1 + variable2":
for _, v := range tree.Variables() {
  fmt.Println(v)
}
// The result would be:
// variable1
// variable2
```

After parsing the expression, it needs to be compiled:

```Golang
compilationResult := expression.NewCompilation(tree)
```

Lastly, the compilation result must be evaluated. If variables are used inside the express, we must tell the evaluator. Otherwise, pass an empty map:

```Golang
// No variables used 
result := compilationResult.Evaluate(map[*contracts.VariableSymbol]interface{}{})

// Variables used:
var variables = map[*contracts.VariableSymbol]interface{}{
  // for each variable, you must specify the type of it:
  contracts.NewVariableSymbol("variable1", reflect.Float64): 10.0,
  contracts.NewVariableSymbol("variable2", reflect.Float64): 11.0,
}

result := compilationResult.Evaluate(variables)
```

It is possible to face errors in the evaluation process. Just like the syntax tree, the result of the evaluation contains a `Diagnostics` array (if it is empty, we are good to go). To get the result value, there are mainly two ways, get it as an`interface{}` or casted as a `float64`:

```Golang
if len(result.Diagnostics) > 0 {
  for _, d := range tree.Diagnostics {
    fmt.Println(d.Message, d.Span, d.Span.Length, d.Span.Start)
  }
  panic(1)
}

// Printing the result as an interface{}
fmt.Println(result.Value)

// Printing the result as a float64
fmt.Println(result.FloatCastedValue)
```

## Sample Code

The whole code with the using of variables:

```Golang
tree := syntax.Parse("variable1 + variable2")

// Check if parsing is done without any error
if len(tree.Diagnostics) > 0 {
  for _, d := range tree.Diagnostics {
    fmt.Println(d.Message, d.Span, d.Span.Length, d.Span.Start)
  }
  panic(1)
}

variables := make(map[*contracts.VariableSymbol]interface{})
for _, v := range tree.Variables() {
  variables[contracts.NewVariableSymbol(v, reflect.Float64)] = 10.0  
}

result := compilationResult.Evaluate(variables)

// Check if evaluation is done without any error
if len(result.Diagnostics) > 0 {
  for _, d := range tree.Diagnostics {
    fmt.Println(d.Message, d.Span, d.Span.Length, d.Span.Start)
  }
  panic(1)
}

// Printing the result as an interface{}
fmt.Println(result.Value)

// Printing the result as a float64
fmt.Println(result.FloatCastedValue)
```

## Acknowledgements

This package is built based on the [Immo Landwerth](https://github.com/terrajobst)'s tutorial on [youtube](https://www.youtube.com/playlist?list=PLRAdsfhKI4OWNOSfS7EUu5GRAVmze1t2y). Many thanks to him for making such a great content.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](/LICENSE)
