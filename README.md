# WAMock (Web API Mock)

is a localhost micro-server made for mocking calls to web api.

WAMock is:

- made to be _light_, so it can be used as part of a testing, CI/CD toolchain...
- _offline_, so you can work from a cave... and also so it can be used as part of testing, CI/CD... :)
- _easy to use_, by using either **a single JSON describing the whole API, or a set of JSONs, each describing a route to mock**.
- _lazy sometimes_, you can allow a coroutine to reload the json config every X seconds.

My personal use-case for starting this tool was to work from my commute train with bad internet connection, and also so I would not clutter my DB with trash data.

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

## JSON!

### Using a single JSON config file:

```json
{
  "PUT/api/doggo/{id}": {
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
  },
  "ALL/": {
    "response": "Available routes: PUT/api/doggo/{id}, GET/api/doggo, POST /api/doggo"
  }
}
```

### Using a collection of JSON config files:

`PUT:api/doggo/{id}.json`

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

## EXAMPLES!

### Examples of command use:

- `wamock path/to/config.json` (if we to use a single JSON for the whole configuration)
- `wamock path/to/routes_config_dir/` (if we want to use several JSONs)

### Examples of JSON file naming (in case of using a directory as argument):

- :api:doggo:(id).json => would mock a route such as /api/doggo/2
- :.json => /
