# 🗺️ go‑countries

> Go package providing comprehensive country data in all standard ISO formats

<table>
  <thead>
    <tr>
      <th>CI&nbsp;/&nbsp;CD</th>
      <th>Quality&nbsp;&amp;&nbsp;Security</th>
      <th>Docs&nbsp;&amp;&nbsp;Meta</th>
      <th>Community</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td valign="top" align="left">
        <a href="https://github.com/mrz1836/go-countries/releases">
          <img src="https://img.shields.io/github/release-pre/mrz1836/go-countries?logo=github&style=flat" alt="Latest release">
        </a><br/>
        <a href="https://github.com/mrz1836/go-countries/actions">
          <img src="https://img.shields.io/github/actions/workflow/status/mrz1836/go-countries/run-tests.yml?branch=master&logo=github&style=flat" alt="Build status">
        </a><br/>
		<a href="https://github.com/mrz1836/go-countries/actions">
          <img src="https://github.com/mrz1836/go-countries/actions/workflows/codeql-analysis.yml/badge.svg?style=flat" alt="CodeQL">
        </a><br/>
        <a href="https://github.com/mrz1836/go-countries/commits/master">
		  <img src="https://img.shields.io/github/last-commit/mrz1836/go-countries?style=flat&logo=clockify&logoColor=white" alt="Last commit">
		</a>
      </td>
      <td valign="top" align="left">
        <a href="https://goreportcard.com/report/github.com/mrz1836/go-countries">
          <img src="https://goreportcard.com/badge/github.com/mrz1836/go-countries?style=flat" alt="Go Report Card">
        </a><br/>
		<a href="https://codecov.io/gh/mrz1836/go-countries">
          <img src="https://codecov.io/gh/mrz1836/go-countries/branch/master/graph/badge.svg?style=flat" alt="Code coverage">
        </a><br/>
		<a href="https://scorecard.dev/viewer/?uri=github.com/mrz1836/go-countries">
          <img src="https://api.scorecard.dev/projects/github.com/mrz1836/go-countries/badge?logo=springsecurity&logoColor=white" alt="Security policy">
        </a><br/>
        <a href=".github/SECURITY.md">
          <img src="https://img.shields.io/badge/security-policy-blue?style=flat&logo=springsecurity&logoColor=white" alt="Security policy">
        </a><br/>
		<a href="https://www.bestpractices.dev/projects/10767">
		  <img src="https://www.bestpractices.dev/projects/10767/badge" alt="Best Practices">
		</a><br/>
      </td>
      <td valign="top" align="left">
        <a href="https://golang.org/">
          <img src="https://img.shields.io/github/go-mod/go-version/mrz1836/go-countries?style=flat" alt="Go version">
        </a><br/>
        <a href="https://pkg.go.dev/github.com/mrz1836/go-countries?tab=doc">
          <img src="https://pkg.go.dev/badge/github.com/mrz1836/go-countries.svg?style=flat" alt="Go docs">
        </a><br/>
        <a href=".github/AGENTS.md">
          <img src="https://img.shields.io/badge/AGENTS.md-found-40b814?style=flat&logo=openai" alt="AGENTS.md rules">
        </a><br/>
        <a href="Makefile">
          <img src="https://img.shields.io/badge/Makefile-supported-brightgreen?style=flat&logo=probot&logoColor=white" alt="Makefile Supported">
        </a><br/>
		<a href=".github/dependabot.yml">
          <img src="https://img.shields.io/badge/dependencies-automatic-blue?logo=dependabot&style=flat" alt="Dependabot">
        </a>
      </td>
      <td valign="top" align="left">
        <a href="https://github.com/mrz1836/go-countries/graphs/contributors">
          <img src="https://img.shields.io/github/contributors/mrz1836/go-countries?style=flat&logo=contentful&logoColor=white" alt="Contributors">
        </a><br/>
        <a href="https://github.com/sponsors/mrz1836">
          <img src="https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat" alt="Sponsor">
        </a><br/>
        <a href="https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-countries&utm_term=go-countries&utm_content=go-countries">
          <img src="https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat" alt="Donate Bitcoin">
        </a>
      </td>
    </tr>
  </tbody>
</table>

<br/>

## 🗂️ Table of Contents

* [Installation](#-installation)
* [Usage](#-usage)
* [Documentation](#-documentation)
* [Examples & Tests](#-examples--tests)
* [Benchmarks](#-benchmarks)
* [Code Standards](#-code-standards)
* [AI Compliance](#-ai-compliance)
* [Credits](#-credits)
* [Maintainers](#-maintainers)
* [Contributing](#-contributing)
* [License](#-license)

<br/>

## 📦 Installation

**go-countries** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).

```bash script
go get -u github.com/mrz1836/go-countries
```
<br/>

## 💡 Usage

Here is a basic example of how to use go-countries in your Go project:

```go
package main

import (
    "fmt"
    "github.com/mrz1836/go-countries"
)

func main() {
    country := countries.GetByAlpha2(countries.Alpha2US)
    if country != nil {
        fmt.Printf("Country: %s, Capital: %s\n", country.Name, country.Capital)
    } else {
        fmt.Println("Country not found")
    }
}
```

- Explore additional [usage examples](examples) for practical integration patterns
- Review [benchmark results](countries_test.go) to assess performance characteristics
- Examine the comprehensive [test suite](countries_test.go) for validation and coverage
- Inspect the [code generator](generate) for details on data sources and the build process

<br/>

## 📚 Documentation

View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-countries?tab=doc)

> **Heads up!** `go-countries` is intentionally light on dependencies. The only
external package it uses is the excellent `testify` suite—and that's just for
our tests. You can drop this library into your projects without dragging along
extra baggage.

<br/>

### Features

- Comprehensive list of all recognized countries, provided as a Go slice of structs for easy access and manipulation
- Direct access to raw country and currency JSON data in the [data package](data) for custom processing or validation
- Zero `init()` overhead—just import and use the `countries` package without side effects
- Fast, allocation-free lookups for all retrieval functions, ensuring optimal performance in production environments
- Includes region, subregion, capital, and currency information for each country
- Designed for extensibility—add or update country data via code generation from JSON sources
- Well-documented, tested, and benchmarked for reliability and speed

<br/>

### Functions
- [`GetAll()`](countries.go): Retrieve the entire slice of all known countries, including metadata such as names, codes, regions, capitals, and currencies
- [`GetByAlpha2(countries.Alpha2US)`](countries.go): Find a country by its [ISO 3166-1 alpha-2 code](https://en.wikipedia.org/wiki/ISO_3166-2)
- [`GetByAlpha3(countries.Alpha3USA)`](countries.go): Retrieve a country by its [ISO 3166-1 alpha-3 code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3)
- [`GetByCapital("Washington")`](countries.go): Find a country by its capital city in a case-insensitive search
- [`GetByCountryCode("840")`](countries.go): Lookup by [ISO 3166 numeric country code](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes), supporting string or integer input
- [`GetByISO31662("ISO 3166-2:US")`](countries.go): Retrieve a country by its [ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2)
- [`GetByName("United States of America")`](countries.go): Lookup a country by its [official name](https://en.wikipedia.org/wiki/ISO_3166), supporting case-insensitive queries

<br/>

### Code Generation

If you need to update the country data or regenerate the Go code, you can use the `go generate` command.
This will read the JSON data files and generate a Go file containing all the country data in a structured format.

To generate the Go code for the country data, follow these steps:

1. **Navigate to the project root directory**:

   ```bash
   cd /path/to/your/project
   ```

2. **Run the `go generate` command**:

   ```bash
   go generate -x ./generate/...
   ```

This command executes the code generation logic defined in the `generate.go` file located in the `/generate/` directory.
The generated code is written to `countries_data.go` in the project directory.

<br/>

### Additional Documentation & Repository Management

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

```bash script
make help
```

List of all current commands:

<!-- make-help-start -->
```text
all                   ## Runs multiple commands
bench                 ## Run all benchmarks in the Go application
build-go              ## Build the Go application (locally)
citation              ## Update version in CITATION.cff (use version=X.Y.Z)
clean-mods            ## Remove all the Go mod cache
coverage              ## Show test coverage
diff                  ## Show git diff and fail if uncommitted changes exist
generate              ## Run go generate in the base of the repo
godocs                ## Trigger GoDocs tag sync
govulncheck-install   ## Install govulncheck
help                  ## Display this help message
install-go            ## Install using go install with specific version
install-releaser      ## Install GoReleaser
install               ## Install the application binary
lint                  ## Run the golangci-lint application (install if not found)
release-snap          ## Build snapshot binaries
release-test          ## Run release dry-run (no publish)
release               ## Run production release (requires github_token)
run-fuzz-tests        ## Run fuzz tests for all packages
tag-remove            ## Remove local and remote tag (use version=X.Y.Z)
tag-update            ## Force-update tag to current commit (use version=X.Y.Z)
tag                   ## Create and push a new tag (use version=X.Y.Z)
test-ci-no-race       ## CI test suite without race detector
test-ci-short         ## CI unit-only short tests
test-ci               ## CI full test suite with coverage
test-no-lint          ## Run only tests (no lint)
test-short            ## Run tests excluding integration
test-unit             ## Runs tests and outputs coverage
test                  ## Run lint and all tests
uninstall             ## Uninstall the Go binary
update-linter         ## Upgrade golangci-lint (macOS only)
update-releaser       ## Reinstall GoReleaser
update                ## Update dependencies
vet                   ## Run go vet
```
<!-- make-help-end -->

</details>

<details>
<summary><strong><code>GitHub Workflows</code></strong></summary>
<br/>

| Workflow Name                                                                | Description                                                                                                            |
|------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| [auto-merge-on-approval.yml](.github/workflows/auto-merge-on-approval.yml)   | Automatically merges PRs after approval and all required checks, following strict rules.                               |
| [check-for-leaks.yml](.github/workflows/check-for-leaks.yml)                 | Runs gitleaks to detect secrets on a daily schedule.                                                                   |
| [codeql-analysis.yml](.github/workflows/codeql-analysis.yml)                 | Analyzes code for security vulnerabilities using GitHub CodeQL.                                                        |
| [delete-merged-branches.yml](.github/workflows/delete-merged-branches.yml)   | Deletes feature branches after their pull requests are merged.                                                         |
| [dependabot-auto-merge.yml](.github/workflows/dependabot-auto-merge.yml)     | Automatically merges Dependabot PRs that meet all requirements.                                                        |
| [pull-request-management.yml](.github/workflows/pull-request-management.yml) | Labels PRs by branch prefix, assigns a default user if none is assigned, and welcomes new contributors with a comment. |
| [release.yml](.github/workflows/release.yml)                                 | Builds and publishes releases via GoReleaser when a semver tag is pushed.                                              |
| [run-tests.yml](.github/workflows/run-tests.yml)                             | Runs all Go tests and dependency checks on every push and pull request.                                                |
| [scorecard.yml](.github/workflows/scorecard.yml)                             | Runs OpenSSF Scorecard to assess supply chain security.                                                                |
| [stale.yml](.github/workflows/stale.yml)                                     | Warns about (and optionally closes) inactive issues and PRs on a schedule or manual trigger.                           |
| [sync-labels.yml](.github/workflows/sync-labels.yml)                         | Keeps GitHub labels in sync with the declarative manifest at `.github/labels.yml`.                                     |

</details>

<details>
<summary><strong><code>Updating Dependencies</code></strong></summary>
<br/>

To update all dependencies (Go modules, linters, and related tools), run:

```bash
make update
```

This command ensures all dependencies are brought up to date in a single step, including Go modules and any tools managed by the Makefile. It is the recommended way to keep your development environment and CI in sync with the latest versions.

</details>

<br/>

## 🧪 Examples & Tests

All unit tests and [examples](examples) run via [GitHub Actions](https://github.com/mrz1836/go-countries/actions) and use [Go version 1.22.x](https://go.dev/doc/go1.22). View the [configuration file](.github/workflows/run-tests.yml).

Run all tests:

```bash script
make test
```

<br/>

## ⚡ Benchmarks

Run the Go [benchmarks](countries_test.go):

```bash script
make bench
```

<br/>

Performance benchmarks for the core functions in this library, executed on an Apple M1 Max (ARM64):

| Function                           | Ops/sec (approx) | Time per op | Allocations | Bytes Allocated |
|------------------------------------|------------------|-------------|-------------|-----------------|
| [`GetAll`](countries.go)           | 5,263,413        | 225.4 ns    | 1           | 2,048 B         |
| [`GetByAlpha2`](countries.go)      | 100,000,000      | 10.36 ns    | 0           | 0 B             |
| [`GetByAlpha3`](countries.go)      | 100,000,000      | 10.97 ns    | 0           | 0 B             |
| [`GetByCapital`](countries.go)     | 33,426,183       | 35.19 ns    | 1           | 16 B            |
| [`GetByCountryCode`](countries.go) | 100,467,049      | 11.10 ns    | 0           | 0 B             |
| [`GetByISO31662`](countries.go)    | 59,682,316       | 20.16 ns    | 0           | 0 B             |
| [`GetByName`](countries.go)        | 46,466,527       | 25.66 ns    | 0           | 0 B             |

> These benchmarks reflect fast, allocation-free lookups for most retrieval functions, ensuring optimal performance in production environments.

<br/>

## 🛠️ Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## 🤖 AI Compliance
This project documents expectations for AI assistants using a few dedicated files:

- [AGENTS.md](.github/AGENTS.md) — canonical rules for coding style, workflows, and pull requests used by [Codex](https://chatgpt.com/codex).
- [CLAUDE.md](.github/CLAUDE.md) — quick checklist for the [Claude](https://www.anthropic.com/product) agent.
- [.cursorrules](.cursorrules) — machine-readable subset of the policies for [Cursor](https://www.cursor.so/) and similar tools.
- [sweep.yaml](.github/sweep.yaml) — rules for [Sweep](https://github.com/sweepai/sweep), a tool for code review and pull request management.

Edit `AGENTS.md` first when adjusting these policies, and keep the other files in sync within the same pull request.

<br/>

## 🙏 Credits
Used: [ISO-3166-Countries-with-Regional-Codes](https://github.com/lukes/ISO-3166-Countries-with-Regional-Codes)     
Used: [List of Countries & Currencies](https://gist.github.com/tiagodealmeida/0b97ccf117252d742dddf098bc6cc58a)     

<br/>

## 👥 Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:------------------------------------------------------------------------------------------------:|
|                                [MrZ](https://github.com/mrz1836)                                 |

<br/>

## 🤝 Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and please follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:! 
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:. 
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap: 
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-countries&utm_term=go-countries&utm_content=go-countries) to ensure this journey continues indefinitely! :rocket:


[![Stars](https://img.shields.io/github/stars/mrz1836/go-countries?label=Please%20like%20us&style=social)](https://github.com/mrz1836/go-countries/stargazers)

<br/>

## 📝 License

[![License](https://img.shields.io/github/license/mrz1836/go-countries.svg?style=flat)](LICENSE)
