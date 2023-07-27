<p align="center">
  <img width="200px" src="images/tracky-logo.png" title="Tracky">
</p>
<p align="center">
  <img src="https://img.shields.io/github/checks-status/MDeLuise/tracky/main?style=for-the-badge&label=build&color=%239400D3" />
<img src="https://img.shields.io/github/v/release/MDeLuise/tracky?style=for-the-badge&color=%239400D3" />
</p>

<p align="center">Tracky is a <b>self-hosted tracker service.</b><br>Useful to keep track of things like: gas prices, mileage, weight, bills, blood glucose levels, grocery prices, etc.</p>

<p align="center"><a href="https://github.com/MDeLuise/tracky/#why">Why?</a> • <a href="https://github.com/MDeLuise/tracky/#features-highlight">Features highlights</a> • <a href="https://github.com/MDeLuise/tracky/#getting-started">Getting started</a> • <a href="https://github.com/MDeLuise/tracky/#configuration">Configuration</a></p>

<p align="center">
  <img src="/images/screenshot-desktop.png" width="45%" />
  <img src="/images/screenshot-mobile.png" width="45%" /> 
</p>

## Why?
Tracky is a track application that helps you keep track of different values over time.

I've always loves to keep track of different things: weight, car mileage, GPL prices, etc.
So I ended up creating multiple excel files almost identical each other.

Tracky takes care of this problem providing a system used to track the values and visualize them in a simple graph. It also provides a Graphical User Interface and the REST API, so the values can be updated from everywhere.

## Features highlight
* Create tracker in order to group different value with the same meaning
* Add new values for the saved trackers
* List all the tracker's values over time
* View all the tracker's value in a chart
* Add, modify and delete values and trackers via REST API or GUI
* 🔜 Group together values over time
* 🔜 Dark/Light mode


## Getting started
Tracky provides multiple ways of installing it on your server.
* [Setup with Docker](https://www.tracky.org/docs/v1/setup/setup-with-docker/) (_recommended_)
* [Setup without Docker](https://www.tracky.org/docs/v1/setup/setup-without-docker/)

### Setup with docker
Working with Docker is pretty straight forward. To make things easier, a [docker compose file](#) is provided in the repository which contain all needed services, configured to just run the application right away.

There are two different images for the service:
* `msdeluise/tracky-backend`
* `msdeluise/tracky-frontend`

This images can be use indipendently, or they can be use in a docker-compose file.
For the sake of simplicity, the provided docker-compose.yml file is reported here:
```
version: "3"
name: tracky
services:
  backend:
    image: msdeluise/tracky-backend:latest
    env_file: backend.env
    depends_on:
      - db
    restart: unless-stopped

  db:
    image: mysql:8.0
    restart: always
    env_file: backend.env

  frontend:
    image: msdeluise/tracky-frontend:latest
    env_file: frontend.env
    links:
      - backend

  reverse-proxy:
    image: nginx:stable-alpine
    ports:
      - "8080:80"
    volumes:
      - ./default.conf:/etc/nginx/conf.d/default.conf
    links:
      - backend
      - frontend
```

Run the docker compose file (`docker compose -f <file> up -d`), then the service will be available at `localhost:8080`, while the REST API will be available at `localhost:8080/api` (`localhost:8080/api/swagger-ui/index.html` for the documentation of them).

<details>

  <summary>Run on a remote host</summary>

  Please notice that running the `docker-compose` file from another machine change the way to connect to the server. For example, if you run the `docker-compose` on the machine with the local IP `192.168.1.100` then you have to change the backend url in the [REACT_APP_API_URL](#configurations) variable to `http://192.168.1.100:8080/api`. In this case, the frontend of the system will be available at `http://192.168.1.100:8080`, and the backend will be available at `http://192.168.1.100:8080/api`.
</details>

### Setup without docker
The application was developed with being used with Docker in mind, thus this method is not preferred.

#### Requirements
* [JDK 19+](https://openjdk.org/)
* [MySQL](https://www.mysql.com/)
* [React](https://reactjs.org/)

#### Run
1. Be sure to have the `mysql` database up and running
1. Run the following command in the terminal inside the `backend` folder
  `./mvnw spring-boot:run`
1. Run the following command in the terminal inside the `frontend` folder
  `npm start`

Then, the frontend of the system will be available at `http://localhost:3000`, and the backend at `http://localhost:8085/api`.


## Configuration

There are 2 configuration file available:
* `deployment/backend.env`: file containing the configuration for the backend. An example of content is the following:
  ```
  MYSQL_HOST=db
  MYSQL_PORT=3306
  MYSQL_USERNAME=root
  MYSQL_PSW=root
  JWT_SECRET=putTheSecretHere
  JWT_EXP=1
  MYSQL_ROOT_PASSWORD=root
  MYSQL_DATABASE=bootdb
  USERS_LIMIT=-1 # including the admin account, so <= 0 if undefined, >= 2 if defined
  ```
  Change the properties values according to your system.

* `deployment/frontend.env`: file containing the configuration for the frontend. An example of content is the following:
  ```
  REACT_APP_API_URL=http://localhost:8080/api
  BROWSER=none
  REACT_APP_PAGE_SIZE=25
  ```
  Change the properties values according to your system.

