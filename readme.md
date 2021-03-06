# Slot Machine

Slot Machine game which is configured with Atkins diet.


## Setup
- verify you've go installed. `go -version` and $GOPATH
- clone the repo in `$GOPATH/github.com/devdinu/`
- run `make` which downloads dependencies, and run vet, lint and tests.
- `make run` to run the service, alternatively `SERVER_PORT=8001 ./scripts/run_server.sh`
- run `./scripts/make_request.sh` to make a sample request with JWT Token against the service.


### Configuration
`config.yaml` contains the configuration required to run the service.

### Response
* The response contains `jwt_token` field which contains `user` information along with new token
```json
{
    "jwt_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiZXQiOjUwMCwiY2hpcHMiOjExODQ1LCJ1aWQiOiJ1c2VyLWlkIiwiZXhwIjoxNTQ0MjY1Mzk0fQ", 
    "spins": [
        {
            "stops": [
                10,
                9,
                31,
                6,
                8
            ],
            "total": 0,
            "type": "main"
        }
    ],
    "total": 0,
    "user": {
        "bet": 500,
        "chips": 11845,
        "uid": "user-id"
    }
}
```

### Testing
- `make test` runs the unit tests
- `manual_mode` is needed to enabled the coded stopper, to stop the reels at specified `manual_positions` in stopper config
- set `manual_mode` to false to run in normal random stopper mode

### Pending:
- currently code handles single spin
- add logger with levels
- integration tests
