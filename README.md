# RMT
![Build & Test](https://github.com/reyesml/RMT/actions/workflows/build_and_test.yml/badge.svg)


RMT gives the power of CRM's to individuals.

## Key features:
### Journal
Record your thoughts. Track your state of mind over time.

### People
Manage relationships like a pro.

### Initiatives
Coming soon.
___

## Development

### Recommended pre-requisites
- npm v8.19.2+
- go v1.18.2+
- make v3.81+

### Running the project

Download the dependencies by running `make` in the project root.

Configure the server by copying the contents of `app/example.config.yml` into a new `app/config.yml` file.
Replace the signing secret with your own secret. The signing secret is used for session validation.

Initialize the dev database by running `make init-dev-db`. This will create a new SQLite database with all the tables
created. It will also add a new "admin" user with a default password of "not_secure".

Launch the backend service by running `make run-dev-server`.

Open a new terminal window, and run `make run-dev-client`.


### Project Architecture

![image](https://user-images.githubusercontent.com/4985056/212031430-2119355d-3695-4986-bab6-413eef292c74.png)

#### Backend
The backend of RMT is split into two distinct segments: Core, and the HTTP server.

`app/core/` is the heart of the service.  It is responsible for database interactions (`database/`, `repos/`), defining  the models used throughout the app (`models/`), and defining the business logic (`interactors/`).  `core/` is protocol agnostic.

`app/httpserver` is an HTTP server that exposes `app/core`'s functionality to web clients. `app/httpserver` defines different HTTP controllers to handle incoming requests (`controllers/`), and those controllers delegate work to `app/core`.

#### Frontend
The frontend of RMT lives in `web-client/`. The web client is implemented using [SvelteKit](https://kit.svelte.dev/), and the project project loosely follows [SvelteKit project conventions](https://kit.svelte.dev/docs/project-structure).
