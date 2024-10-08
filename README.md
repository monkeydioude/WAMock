# WAMock (Web API Mock)

is a localhost micro-server made for mocking calls to web APIs.
**For example, with WAMock, you could mock calls like `POST /api/doggo`, and always get an answer, without the need for an internet connection or a running DB.**

WAMock is:

- made to be _light_, so it can be used as part of a testing, CI/CD toolchain...
- _offline_, so you can work from a cave... and also so it can be used as part of testing, CI/CD... :)
- _easy to use_, by using either **a single JSON describing the whole API, or a set of JSONs, each describing a route to mock**.
- _lazy sometimes_, you can allow a coroutine to reload the json config every X seconds.
- _dockerized_, see below.
My personal use-case for starting this tool was to work from my commute train with bad internet connection, and also so I would not clutter my DB with trash data.

![Go](https://github.com/monkeydioude/wamock/actions/workflows/go.yml/badge.svg)

## PARAMETERS!

### Using a single JSON as parameter:

1st argument should be the path to a single JSON file holding the complete configuration of the Mock API (ie: `path/to/config.json`). This is conveniant if we want to mock a few routes, or if we plan to use different configurations for different cases.

### Using a directory as parameter:

1st argument should then be the path of a directory holding a collection of JSON files, named after the HTTP methods and the routes they mock. For example:

```sh
drwxr-xr-x  6 mkd  staff  192  4 aoû 18:14 .
drwxr-xr-x  3 mkd  staff   96  4 aoû 18:11 ..
-rw-r--r--  1 mkd  staff    0  4 aoû 18:11 ALL:.json
-rw-r--r--  1 mkd  staff    0  4 aoû 18:12 GET:api:doggo.json
-rw-r--r--  1 mkd  staff    0  4 aoû 18:12 POST:api:doggo.json
-rw-r--r--  1 mkd  staff    0  4 aoû 18:12 PUT:api:doggo:{id}.json
```

`PUT:api:doggo:{id}.json` would handle a `PUT` request made to the route `/api/doggo/21`

In cURL lingo, such filename would cover such call:

```sh
curl -H 'Content-Type: application/json' \
      -d '{ "title":"foo","body":"bar", "id": 1}' \
      -X PUT \
      http://localhost:8088/api/doggo/21
```

### Provided flags:

- -x `{s}`: enable the config auto-reload coroutine and also set it to trigger every `s` seconds
- -p `{port}`: choose the port you wish to run the server on. _Default is `8088`_

## JSON!

### Using a single JSON config file:

```json
{
  "PUT/api/doggo/{id}": {
    "response": {
      "id": 25,
      "size": "HUGE",
      "weight": 2000
    }
  },
  "ALL/": {
    "response": "Available routes: PUT/api/doggo/{id}, GET/api/doggo, POST /api/doggo"
  }
}
```

### Using a collection of JSON config files:

`PUT:api:doggo:{id}.json`

```json
{
  "request": {
    "size": "",
    "weight": 0,
    "name": ""
  },
  "response": {
    "id": 25,
    "size": "HUGE",
    "weight": 2000
  }
}
```

## DOCKER!

[Here is the image's repo](https://hub.docker.com/repository/docker/drannoc/wamock/general).

To run this image, a **PATH to a valid config JSON file is required**. Such JSON file should be **[mounted as a volume (-v)](https://docs.docker.com/storage/volumes/#choose-the--v-or---mount-flag) and used as an ARG to the `docker run` command**
As such:

`docker run -p 8088:8088 -v $PWD/api_mock.json:/app/api_mock.json drannoc/wamock /app/api_mock.json -x 30`

Note that the `/app/api_mock.json` PATH provided as an ARG matches the second part of the [mount volume (-v)](https://docs.docker.com/storage/volumes/#choose-the--v-or---mount-flag) parameter.

## EXAMPLES!

### Examples of command use:

- `wamock path/to/config.json -x 30` (if we to use a single JSON for the whole configuration and refresh the config every 30s)
- `wamock path/to/routes_config_dir/` (if we want to use several JSONs)

### Examples of JSON file naming (in case of using a directory as argument):

- :api:doggo:(id).json => would mock a route such as /api/doggo/2
- :.json => /

## TODO!

- ~~Handle config directory~~
- ~~Docker image~~
- ~~Handle hot reload of the config through goroutines~~
- Match request defined in config with the request's payload
- (MAYBE NOT) Handle dynamic parameters in path that should be reflected in the response
