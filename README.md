# go-json-diff
Check if two given JSON files contain the same data.

## Up and running

To get the project up and running, you first need to clone the repository.
To do so, type the following on your terminal:

```sh
$ git clone https://github.com/fernandoporazzi/go-json-diff.git
```

Next, you have to `cd` into the project folder. To run the application, type the following command:

```sh
$ go run main.go
```

## Supported shapes(so far)

Since the JSON inputs can come in any shape and have multiple levels of nesting, the complexity of the `compare()` function can grow exponentially. In order to achieve a better functionality, the compare function should be able to handle recursive calls to verify equality of deeply nested objects.

So far the only available handlers verify the first level of nesting of the files.

Shapes supported:

```json
{
  "data": {
    "name": "Peter",
    "id": "123"
  }
}
```

or then:

```json
[
  {
    "id":"jkl",
    "name":"Peter"
  },
  {
    "id":"xyz",
    "name":"Mary"
  },
  {
    "id":"abc",
    "name":"Charles"
  }
]
```

Play around by changing the order of the objects, changing the order of the keys of an object or even changing types(`string` -> `int`).

## Known issues

If there is a space between objects, the program will assume files have different sizes. This could be solved by using `json.Compact`

It still does not support deeply nested objects yet.

## Testing

To run the tests locally, run the following command:

```sh
$ go test ./... 
```

If you want to see how covered the project is, you can run the following command to get the coverage report"

```sh
$ go test ./... -coverprofile=coverage.out
```

Once the above has been run, it's time to see it in your browser. The following command will open a new tab in your browser with the code coverage.

```sh
$ go tool cover -html=coverage.out
```

## TODO

- Support deeply nested objects and types
