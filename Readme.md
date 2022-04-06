# Roman Numeral Conversion

## Definition (Wikipedia)

Roman numerals are a numeral system that originated in ancient Rome and remained the usual way of writing numbers throughout Europe well into the Late Middle Ages. Numbers in this system are represented by combinations of letters from the Latin alphabet. Modern style uses seven symbols, each with a fixed integer value (Wikipedia)

### Some Roman Numerals

| Symbol | Value |
|--------|-------|
| I      | 1     |
| V      | 5     |
| X      | 10    |
| L      | 50    |
| C      | 100   |
| D      | 500   |
| M      | 1000  |

## How to build and run the project

### INFO
Extensions are implemented.

- This project handles all of the roman numerals ranged between 1 - 3999.
- This project converts roman numerals between a range in parallel.
- This project is containerized.


### Environment Variables

There are 2 different ways of running the project.

#### With Docker and Docker-compose

Docker-compose gets 2 environment variables.

1. `ADOBE_APP_PORT`: The port you want the container to run the http server. If it is not provided, then the value is 8080.
2. `ADOBE_PROXY_PORT`: The host port from localhost to docker container. You can set this variable from `.env`. Default is 8080. 

There are some scripts to easily run the project. However, to need to project from these scripts
make sure that you have installed [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/)
in your development environment.

After installing necessary requirements, you can simply run

```
make run
```

which will run the project in port `8080` in your local environment.

You can simply send a request to get a roman numeral conversion,

```bash
curl --location --request GET 'http://localhost:8080/romannumeral?query=1'
```

and the response

```json
{
    "input": "1",
    "output": "I"
}
```

To get the range conversions, simply run

```bash
curl --location --request GET 'http://localhost:8080/romannumeral?min=1&max=3'
```

and the response

```json
{
    "conversions": [
        {
            "input": "1",
            "output": "I"
        },
        {
            "input": "2",
            "output": "II"
        },
        {
            "input": "3",
            "output": "III"
        }
    ]
}
```

To stop the project, simply run

```
make stop
```

and to remove the container

```
make down
```


#### If you do not have Docker

You can provide an environment variable called the `PORT` or the program can use default 8000.

You can simply run

```
PORT=<YOUR_PORT> go run cmd/main.go
```

It will run the project in the `PORT`.

## Testing of the project

Project has unit tests for both the http handler and the roman conversion function.

For the http handlers, project tests if the correct values are retrieved from different 
http requests. 

To run the tests, you can simply run

```
make test
```

## Packaging Layout

| Folder  | Description                                |
| ------------- |--------------------------------------------|
| cmd  | contains the main function of the program. |
| internal  | roman conversion related files             |
| pkg | http handlers                              |

## Tests

There are tests both in `pkg` and `internal` folders. For the `pkg`,
the http requests are tested to see if the returned responses are correct.

Tests in `internal` package is used to test if roman conversion is right. 

## Error Handling

I added some error handling logics below.

1. Check if any query parameter or min max query parameters are provided.
2. Check if the provided query parameters are integers
3. Check if min value is not bigger than the max value

## REFERENCES

[Wikipedia - Roman Numerals](https://en.wikipedia.org/wiki/Roman_numerals)