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

I've always loves to keep track of different things: weight, bills, car mileage, GPL prices, etc.
So I ended up creating multiple excel files almost identical each other.

Tracky takes care of this problem providing a system used to track the values and visualize them in a simple graph. It also provides a Graphical User Interface and the REST API, so the values can be updated from everywhere.

## Features highlight
* Create tracker in order to group different value with the same meaning
* Add new values for the saved trackers
* List all the tracker's values over time
* View all the tracker's value in a chart
* Add, modify and delete values and trackers via REST API or GUI
* 🔜 Group together values over time
* 🔜 Compare multiple trackers together


## Quickstart 
Installing Tracky is pretty straight forward, in order to do so follow these steps:

1. Create a folder where you want to place all Tracky related files.
1. Inside that folder, create the following files:
    * `docker-compose.yml`:
    ```yaml
    version: "3"
    name: tracky
    services:
      backend:
        image: msdeluise/tracky-backend:latest
        env_file: backend.env
        depends_on:
          - db
        restart: unless-stopped
        volumes:
          - "certs:/certificates"
        ports:
          - "8080:8080"

      db:
        image: mysql:8.0
        restart: always
        env_file: backend.env
        volumes:
          - "./db:/var/lib/mysql"

      frontend:
        image: msdeluise/tracky-frontend:latest
        env_file: frontend.env
        links:
          - backend
        ports:
          - "3000:3000"
        volumes:
          - "certs:/certificates"
    volumes:
      certs:
        driver: local
        driver_opts:
          type: none
          o: bind
          device: ./certificates
    ```
    * `backend.env`:
    ```properties
    #
    # DB
    #
    MYSQL_HOST=db
    MYSQL_PORT=3306
    MYSQL_USERNAME=root
    MYSQL_PSW=root
    MYSQL_ROOT_PASSWORD=root
    MYSQL_DATABASE=bootdb
    
    #
    # JWT
    #
    JWT_SECRET=putTheSecretHere
    JWT_EXP=1
    
    #
    # Server config
    #
    USERS_LIMIT=-1 # < 0 means no limit

    #
    # SSL
    #
    SSL_ENABLED=false
    CERTIFICATE_PATH=/certificates/
    ```
    * `frontend.env`:
    ```properties
    PORT=3000
    API_URL=http://localhost:8080/api
    WAIT_TIMEOUT=10000
    BROWSER=none
    SSL_ENABLED=false
    CERTIFICATE_PATH=/certificates/
    ```

1. Run the docker compose file (`docker compose -f docker-compose.yml up -d`), then the service will be available at `localhost:3000`, while the REST API will be available at `localhost:8080/api` (`localhost:8080/api/swagger-ui/index.html` for the documentation of them).

Run the docker compose file (`docker compose -f <file> up -d`), then the service will be available at `localhost:8080`, while the REST API will be available at `localhost:8080/api` (`localhost:8080/api/swagger-ui/index.html` for the documentation of them).

<details>

  <summary>Run on a remote host (e.g. from mobile)</summary>

  Please notice that running the `docker-compose` file from another machine change the way to connect to the server. For example, if you run the `docker-compose` on the machine with the local IP `192.168.1.100` then you have to change the backend url in the `API_URL` variable to `192.168.1.100:8080/api`. In this case, the frontend of the system will be available at `192.168.1.100:8080`, and the backend will be available at `192.168.1.100:8080/api`.
</details>

## Contribute
Feel free to contribute and help improve the repo.

### Bug Report, Feature Request and Question
You can submit any of this in the [issues](https://github.com/MDeLuise/tracky/issues/new/choose) section of the repository. Chose the right template and then fill the required info.

### Bug fix
If you fix a bug, please follow the [contribution-guideline](https://github.com/MDeLuise/tracky#contribution-guideline) in order to merge the fix in the repository.

### Feature development
Let's discuss first possible solutions for the development before start working on that, please open a [feature request issue](https://github.com/MDeLuise/tracky/issues/new?assignees=&labels=&projects=&template=fr.yml).

### Contribution guideline
To fix a bug or create a feature, follow these steps:
1. Fork the repo
1. Create a new branch (`git checkout -b awesome-feature`)
1. Make changes or add new changes.
1. Commit your changes (`git add -A; git commit -m 'Awesome new feature'`)
1. Push to the branch (`git push origin awesome-feature`)
1. Create a Pull Request

#### Conventions
* Commits should follow the [semantic commit](https://www.conventionalcommits.org/en/v1.0.0/) specification, although not mandatory.
