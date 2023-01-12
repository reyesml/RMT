# RMT
![Build & Test](https://github.com/reyesml/RMT/actions/workflows/build_and_test.yml/badge.svg)


The Relationship Management Tool (RMT) gives the power of CRM's to individuals.

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
- go v1.18.2
- make v3.81

### Running the project

Download the dependencies by running `make` in the project root.

Configure the server by copying the contents of `example.config.yml` into a new `config.yml` file in the same directory.
Replace the signing secret with your own secret. The signing secret is used for session validation.

Initialize the dev database by running `make init-dev-db`. This will create a new SQLite database with all of the tables
created. It will also add a new "admin" user with a default password of "not_secure".

Launch the backend service by running `make run-dev-server`.

Open a new terminal window, and run `make run-dev-client`.