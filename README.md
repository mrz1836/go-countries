# go-countries
> Complete go-ready list of countries in all standardized formats 

<div align="center">

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-countries.svg?logo=github&style=flat)](https://github.com/mrz1836/go-countries/releases)
[![Last Commit](https://img.shields.io/github/last-commit/mrz1836/go-countries?style=flat)](https://github.com/mrz1836/go-countries/commits/master)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mrz1836/go-countries?style=flat)](https://golang.org/)
[![Docs](https://pkg.go.dev/badge/github.com/mrz1836/go-countries.svg?style=flat)](https://pkg.go.dev/github.com/mrz1836/go-countries?tab=doc)

[![Build Status](https://img.shields.io/github/actions/workflow/status/mrz1836/go-countries/run-tests.yml?branch=master&logo=github&style=flat)](https://github.com/mrz1836/go-countries/actions)
[![CodeQL](https://github.com/mrz1836/go-countries/actions/workflows/codeql-analysis.yml/badge.svg?style=flat&logoColor=white)](https://github.com/mrz1836/go-countries/actions)
[![Code Coverage](https://codecov.io/gh/mrz1836/go-countries/branch/master/graph/badge.svg?style=flat)](https://codecov.io/gh/mrz1836/go-countries)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrz1836/go-countries?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-countries)

[![AGENTS.md](https://img.shields.io/badge/AGENTS.md-found-40b814?style=flat&logo=openai)](.github/AGENTS.md)
[![Security Policy](https://img.shields.io/badge/security-policy-blue?style=flat&logo=springsecurity&logoColor=white)](.github/SECURITY.md)
[![Mergify](https://img.shields.io/endpoint.svg?url=https://api.mergify.com/v1/badges/mrz1836/go-countries?style=flat&logoColor=white)](.github/mergify.yml)
[![Dependabot](https://img.shields.io/badge/dependencies-auto--updated-blue?logo=dependabot&style=flat)](.github/dependabot.yml)

[![Contributors](https://img.shields.io/github/contributors/mrz1836/go-countries?style=flat&logo=contentful&logoColor=white)](https://github.com/mrz1836/go-countries/graphs/contributors)
[![Sponsor](https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat)](https://github.com/sponsors/mrz1836)
[![Donate Bitcoin](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat)](https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-countries&utm_term=go-countries&utm_content=go-countries)

</div>

<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [AI Compliance](#ai-compliance)
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
 
### Code Generation

To generate the Go code for the country data, follow these steps:

1. **Navigate to the project root directory**:
   ```shell
   cd /path/to/your/project
   ```

2. **Run the `go generate` command**:
   ```shell
   go generate ./generate/...
   ```

This command will execute the code generation logic defined in the `generate.go` file located in the `/generate/` directory. The generated code will be written to `countries_data.go` in the project directory.

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

This project uses [goreleaser](https://github.com/goreleaser/goreleaser) for streamlined binary and library deployment to GitHub. To get started, install it via:

```bash
brew install goreleaser
```

The release process is defined in the [.goreleaser.yml](.goreleaser.yml) configuration file.

To generate a snapshot (non-versioned) release for testing purposes, run:

```bash
make release-snap
```

Before tagging a new version, update the release metadata in the `CITATION.cff` file:

```bash
make citation version=0.2.1
```

Then create and push a new Git tag using:

```bash
make tag version=x.y.z
```

This process ensures consistent, repeatable releases with properly versioned artifacts and citation metadata.

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
all                      Runs multiple commands
citation                 Update version in CITATION.cff (citation version=X.Y.Z)
clean                    Remove previous builds and any test cache data
clean-mods               Remove all the Go mod cache
coverage                 Shows the test coverage
diff                     Show the git diff
generate                 Runs the go generate command in the base of the repo
godocs                   Sync the latest tag with GoDocs
govulncheck-install      Install govulncheck for vulnerability scanning
help                     Show this help message
install                  Install the application
install-go               Install the application (Using Native Go)
install-releaser         Install the GoReleaser application
lint                     Run the golangci-lint application (install if not found)
release                  Full production release (creates release in GitHub)
release                  Runs common.release then runs godocs
release-snap             Test the full release (build binaries)
release-test             Full production test release (everything except deploy)
replace-version          Replaces the version in HTML/JS (pre-deploy)
tag                      Generate a new tag and push (tag version=0.0.0)
tag-remove               Remove a tag if found (tag-remove version=0.0.0)
tag-update               Update an existing tag to current commit (tag-update version=0.0.0)
test                     Runs lint and ALL tests
test-ci                  Runs all tests via CI (exports coverage)
test-ci-no-race          Runs all tests via CI (no race) (exports coverage)
test-ci-short            Runs unit tests via CI (exports coverage)
test-no-lint             Runs just tests
test-short               Runs vet, lint and tests (excludes integration tests)
test-unit                Runs tests and outputs coverage
uninstall                Uninstall the application (and remove files)
update-linter            Update the golangci-lint package (macOS only)
vet                      Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](examples) run via [GitHub Actions](https://github.com/mrz1836/go-countries/actions) and
 use [Go version 1.22.x](https://go.dev/doc/go1.22). View the [configuration file](.github/workflows/run-tests.yml).

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

## AI Compliance
This project documents expectations for AI assistants using a few dedicated files:

- [AGENTS.md](.github/AGENTS.md) — canonical rules for coding style, workflows, and pull requests.
- [CLAUDE.md](.github/CLAUDE.md) — quick checklist for the Claude agent.
- [.cursorrules](.cursorrules) — machine-readable subset of the policies for Cursor and similar tools.
- [sweep.yaml](.github/sweep.yaml) — rules for Sweep AI, a tool for code review and pull request management.

Edit `AGENTS.md` first when adjusting these policies, and keep the other files in sync within the same pull request.

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
