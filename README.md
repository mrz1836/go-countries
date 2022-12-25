# go-countries
> Complete go-ready list of countries in all standardized formats 

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-countries.svg?logo=github&style=flat)](https://github.com/mrz1836/go-countries/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/mrz1836/go-countries/run-tests.yml?branch=master&logo=github&v=3)](https://github.com/mrz1836/go-countries/actions)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-countries?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-countries)
[![codecov](https://codecov.io/gh/mrz1836/go-countries/branch/master/graph/badge.svg)](https://codecov.io/gh/mrz1836/go-countries)
[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-countries)](https://golang.org/)
[![Sponsor](https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat&v=3)](https://github.com/sponsors/mrz1836)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat)](https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-countries&utm_term=go-countries&utm_content=go-countries)

<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Credits](#credits)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

<br/>

## Installation

**go-countries** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-countries
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-countries?tab=doc)

[![GoDoc](https://godoc.org/github.com/mrz1836/go-countries?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-countries?tab=doc)
 
### Features
- All known countries in a friendly Go slice of structs
- All JSON is still available in the [data package](data)
- No `init()` method required for using the `countries` package
- `GetAll()` will return the entire slice of all known countries
- `GetByName("Nigeria")` will return the country by its [proper name](https://en.wikipedia.org/wiki/ISO_3166)
- `GetByAlpha2("NG")` will return the country by its [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)
- `GetByAlpha3("NGA")` will return the country by its [ISO 3166-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3)
- `GetByCountryCode("566")` will return the country by its [ISO 3166 country code](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes)
- `GetByISO31662("ISO 3166-2:NG")` will return the country by its [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                  Runs multiple commands
clean                Remove previous builds and any test cache data
clean-mods           Remove all the Go mod cache
coverage             Shows the test coverage
generate             Runs the go generate command in the base of the repo
godocs               Sync the latest tag with GoDocs
help                 Show this help message
install              Install the application
install-go           Install the application (Using Native Go)
lint                 Run the golangci-lint application (install if not found)
release              Full production release (creates release in Github)
release              Runs common.release then runs godocs
release-snap         Test the full release (build binaries)
release-test         Full production test release (everything except deploy)
replace-version      Replaces the version in HTML/JS (pre-deploy)
tag                  Generate a new tag and push (tag version=0.0.0)
tag-remove           Remove a tag if found (tag-remove version=0.0.0)
tag-update           Update an existing tag to current commit (tag-update version=0.0.0)
test                 Runs vet, lint and ALL tests
test-ci              Runs all tests via CI (exports coverage)
test-ci-no-race      Runs all tests via CI (no race) (exports coverage)
test-ci-short        Runs unit tests via CI (exports coverage)
test-short           Runs vet, lint and tests (excludes integration tests)
uninstall            Uninstall the application (and remove files)
update-linter        Update the golangci-lint package (macOS only)
vet                  Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](examples) run via [Github Actions](https://github.com/mrz1836/go-countries/actions) and
uses [Go version 1.15.x](https://golang.org/doc/go1.15). View the [configuration file](.github/workflows/run-tests.yml).

Run all tests (including any integration tests)
```shell script
make test
```

<br/>

## Benchmarks
Run the Go [benchmarks](countries_test.go):
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## Usage
- View the [examples](examples)
- View the [benchmarks](countries_test.go)
- View the [tests](countries_test.go)
- View the [generator](generate)

<br/>

## Credits
Used: [ISO-3166-Countries-with-Regional-Codes](https://github.com/lukes/ISO-3166-Countries-with-Regional-Codes)     
Used: [List of Countries & Currencies](https://gist.github.com/tiagodealmeida/0b97ccf117252d742dddf098bc6cc58a)     

<br/>

## Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:------------------------------------------------------------------------------------------------:|
|                                [MrZ](https://github.com/mrz1836)                                 |

<br/>

## Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and please follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:! 
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:. 
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap: 
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-countries&utm_term=go-countries&utm_content=go-countries) to ensure this journey continues indefinitely! :rocket:


[![Stars](https://img.shields.io/github/stars/mrz1836/go-countries?label=Please%20like%20us&style=social)](https://github.com/mrz1836/go-countries/stargazers)

<br/>

## License

[![License](https://img.shields.io/github/license/mrz1836/go-countries.svg?style=flat)](LICENSE)
