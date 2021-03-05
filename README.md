# Normalize Vanguard Sector Breakdown Data

The lambda function to normalize `Vanguard Canada ETF` overview raw data and extract sector breakdown

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Overview](#overview)
  - [Technical Summary](#technical-summary)
- [Project Structure](#project-structure)
- [Usage](#usage)
  - [Build lambda function](#build-lambda-function)
  - [Build cmd](#build-cmd)
  - [Clean up](#clean-up)
- [How To](#how-to)
  - [Add new build environment](#add-new-build-environment)
- [Contributing](#contributing)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

#### Technical Summary

This lambda function is the 2nd step of a `state machine`. After the first step `Scrape Vanguard Canada ETF funds` is completed, this lambda function will be triggered.

This lambda function normalize the fund overview data in `vanguard_fund_overview` collection and save them into `fund_sectors` collection.

## Project Structure

This project follows the clean architecture with minor variation

```
├── api
│   └── lambda
├── bin
│   ├── cmd
│   └── lambda
├── cmd
├── config
├── consts
├── entities
├── infrastructure
│   └── repositories
│       └── mongodb
│           ├── models
│           └── repo
└── usecase
    └── breakdown
```

## Usage

To build the project for different environment set env variable `LIBRARY_ENV`. Default environment is `dev`. Check [`How To`](#how-to) section to learn how to add new build env.

#### Build lambda function

Bellow command will build a `lambda function` to deploy to `AWS`. The output binary file is in `./bin/api/lambda/main

```bash
# Build lambda function
make build
```

#### Build cmd

Bellow command is to build a `command line` to run `locally`. The output binary file is in `./bin/cmd/main

```bash
# Build cmd
make build-cmd
```

#### Clean up

Bellow command is to clean up the build

```bash
# Clean
make clean
```

## How To

### Add new build environment

- Make a copy of `config_dev.go` and change the suffix to new environment name. For example, adding `staging` the new file should be named `config_staging.go`.
- Change the build tag (the first line) to new environment name. For example, changing `// +build dev` to `// +build staging`
- Finally, update the variable values in new config file accordingly.
- To build the project for the newly added environment don't forget to set env variable `LIBRARY_ENV`

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

This project is not an opensource project. Please contact the owner for permission.
