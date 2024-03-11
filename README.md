- [OpenPass](#openpass)
  - [About the project](#about-the-project)
    - [API docs](#api-docs)
    - [Design](#design)
  - [Getting started](#getting-started)
    - [Setting up](#setting-up)
    - [Building the project](#building-the-project)
    - [Running the tests](#running-the-tests)
    - [Layout](#layout)
  - [Notes](#notes)

# OpenPass

## About the project

OpenPass is a password manager that is designed to be secure, easy to use, and open source. Blah blah blah idk

### API docs

Coming soon...

### Design

The template follows project convention doc.

* [Repository Conventions](https://github.com/caicloud/engineering/blob/master/guidelines/repo_conventions.md)

## Getting started

Below we describe the conventions or tools specific to golang project.

### Setting up

```bash
git clone https://github.com/skjdfhkskjds/openpass-v2.git
cd openpass-v2
```

### Building the project

```bash
make
```

### Generating protobuf files and mocks

```bash
make generate
```

### Running the tests

Unit tests are located in the same package as the code they are testing, and integration
and e2e tests are located in the `test` directory.

For unit tests, run:
```bash
make test-unit
```

Integration and e2e tests coming soon...

### Contributing

More dedicated docs for contributing guidelines coming soon...

#### Linting

Prior to committing, run the following to ensure the code is formatted and linted:
```bash
make format && make lint
```
This handles all license adding and formatting of the code.

#### Testing

Please ensure that all tests pass before committing. See [Running the tests](#running-the-tests) for more details.
Additionally, please ensure that all new code has tests (within reason).

#### Pull requests

Please ensure that all pull requests have a description of the changes made, and request a review from at least one
other contributor (CODEOWNER whenever I set that up).


### Layout

```tree
├── .gitignore
├── .golangci.yaml
├── .mockery.yaml
├── cover.out
├── Makefile
├── LICENSE
├── LICENSE.header
├── README.md
├── build
│   ├── tools
│   └── scripts
├── client
│   ├── cmd
│   └── config
├── config
├── core
│   ├── generator
│   └── keychain
├── crypto
│   ├── argon2
│   └── chacha20
├── docs
├── proto
│   ├── services
│   └── types
├── server
│   ├── cmd
│   └── config
├── store
│   └── local
│       └── db
├── test
└── types
    ├── key
    ├── password
    ├── proto
    └── salt
```

A brief description of the layout:

* `.gitignore` varies per project, but all projects need to ignore `bin` directory.
* `.golangci.yaml` is the golangci-lint config file.
* `.mockery.yaml` is the mockery config file.
* `Makefile` is used to build and manage the project.
<!-- * `OWNERS` contains owners of the project. -->
* `README.md` is a detailed description of the project.
* `bin` is to hold build outputs.
* `build` contains scripts, yaml files, dockerfiles, etc, to build and package the project.
* `client` is the package for client side code.
* `config` contains all root-level configuration code.
* `core` contains the core business logic.
* `crypto` contains the code for all cryptographic operations.
* `docs` for project documentations.
* `proto` contains all protobuf files.
* `server` is the package for server side code.
* `store` contains all code to manage data storage.
* `test` contains all tests (except unit tests), e.g. integration, e2e tests.
* `types` contains all types used in the project.
