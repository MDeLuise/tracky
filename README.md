# Tracky
Tracky is a self-hosted, open source tracker service.
It is used to monitor how a `target`'s values change over time.

## Why Tracky?
_I'm Something of a Scientist Myself_ as Norman Osborn says. I enjoy keeping track of a variety of things, including gas prices, mileage, weight, etc.
I ended up with numerous excel files that were remarkably similar to one another.
Therefore, I made the decision to create an application that can combine all the functionality I had in the excel files with additional features.

## Prerequisite
In order to make the service works the following are needed.

### Snapshot version
If you want to use the snapshot version (i.e. simply use the service from the _main_ branch):
* [Go](https://go.dev/) `v1.16.0` or above
* [Buffalo](https://gobuffalo.io/documentation/getting_started/installation/) `v0.18.8` or above
* [Buffalo Pop](https://gobuffalo.io/pt/documentation/database/pop/)
* [PostgreSQL](https://www.postgresql.org/)

### Release version
If you want to use the release version (i.e. simply use the service from one of the _release-X_ branches):
* [Docker](https://www.docker.com/) `v17.05` or above

## How to run
In order to make the service run follow the following steps.

### Snapshot version
* create the needed [configuration variables](#configuration-variables) inside an `.env` file
* start a postgres server on `5432`
* [optional] if the database is not initialized yet, run `buffalo pop create -a && buffalo pop migrate`
* run `buffalo dev`
* create the [application user](#user-creation)

### Release version
* create the needed [configuration variables](#configuration-variables) inside an `.env` file
* from the `deployment` directory run `docker-compose --env-file <env-file> up` (adding `--profile debug` before `up` will startup even a `pgAdmin` instance)
* create the [application user](#user-creation)

### Configuration variables
```
JWT_SECRET = "<secret key used to encrypt the access token>"
ACCESS_TOKEN_EXPIRATION_SECONDS = <access token expiration expressed in seconds>

JWT_REFRESH_SECRET = "<secret key used to encrypt the refresh token>"
REFRESH_TOKEN_EXPIRATION_SECONDS = <refresh token expiration expressed in seconds>

LOG_LEVEL = "<verbosity>" # can be trace, debug, info, warn, error
```

### User creation
In order to access the application resources, a user is needed.
The user can be created via CLI or via database administration system (e.g. `pgAdmin`):
* via CLI executes the following commands, replacing the `USERNAME` and `PASSWORD` as wanted:
    ```
    $ psql -U postgres -d tracky_development -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\""
    $ psql -U postgres -d tracky_development -c "CREATE EXTENSION IF NOT EXISTS \"pgcrypto\""
    $ psql -U postgres -d tracky_development -c "INSERT INTO users (id, username, password, created_at, updated_at) VALUES (uuid_generate_v1(), '<USERNAME>', crypt('<PASSWORD>', gen_salt('bf')), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);"
    ```
    If postgres is running in a container (e.g. becose run via `docker-compose-up`), the prefix the above command with `docker exec -it deployment-db-1`.

* via database administration system: execute the above commands from the system user interface after connecting to the db.

### Authentication
Authentication can be done via jwt or via api key:
* jwt: pass `Authentication: Bearer <token>` in the requests' header
* api key: pass `key=<api_key>` in the requests' url

## Client
It's possible to handle the `targets` and the `values` via [rest-api](#endpoints) by using simply this repository, which hosts the server's backend.
Visit the [provided client](https://github.com/MDeLuise/tracky-client) repository to utilize the service via mobile app and web app.

## Endpoints
These are the provided enpoints

### Login
* `POST /login` `{username: "foo", password: "bar"}` - get access and refresh token
* `POST /refresh` `{username: "foo", password: "bar"}` - user the refresh token to take new access and refresh tokens

### Api Key
* `GET /key` - get all the keys
* `GET /key/{id}` - get the key with specified id
* `DELETE /key/{id}` - delete the key with specified id
* `POST /key/` - create a key

### Target
* `GET /target` - get all the targets
* `GET /target/{id}` - get the target with specified id
* `DELETE /target/{id}` - delete the target with specified id
* `PUT /target/{id}` `{name: "foo", description: "bar", unit: "$"}` - update the target with specified id
* `POST /target/` `{name: "foo", description: "bar", unit: "$"}` - create a target

### Value
* `GET /value` - get all the values
* `GET /value/{id}` - get the value with specified id
* `DELETE /value/{id}` - delete the value with specified id
* `PUT /value/{id}` `{target_id: "foo", value: 42, time: "2023-09-28T14:08:59Z"}` - update the value with specified id
* `POST /value/` `{target_id: "foo", value: 42, time: "2023-09-28T14:08:59Z"}` - create a target

### Stats
* `GET /stats/mean/{target_id}` - get the mean of the target with specified id
* `GET /mean/{target_id}/{at}` - get the the mean of X if the target with specified id (e.g. mean of last 10 values)
* `GET /increment/{target_id}` - get the last value increment respect the previous of the target with specified id

## How to contribute
Fell free to contribute! Just a few useful information below.


This project use the Trunk-Based Development as source-control branching model. The usual workflow for contributing is the following:
1. create a new branch starting from `main` branch,
1. work on that,
1. create a pull request to merge the branch in the `main`.

Once the pull request is approved, please rebase the branch upon `main` and squash the commits before merging providing a meaningful commit message as follow:
```
Short (72 chars or less) summary

More detailed explanatory text. Wrap it to 72 characters. The blank
line separating the summary from the body is critical (unless you omit
the body entirely).

Write your commit message in the imperative: "Fix bug" and not "Fixed
bug" or "Fixes bug." This convention matches up with commit messages
generated by commands like git merge and git revert.

Further paragraphs come after blank lines.

- Bullet points are okay, too.
- Typically a hyphen or asterisk is used for the bullet, followed by a
  single space. Use a hanging indent.
```


## Future update
- [ ] Multi users support
- [ ] Export of values
- [ ] Filter values by start and end dates
- [ ] Get last _n_ target's values
- [x] Add targets unit of measurement 