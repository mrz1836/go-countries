# AGENTS.md

## Purpose

This file defines the coding conventions, contribution workflows, and repository structure guidelines for AI agents collaborating on the codebase. It is designed to assist AI assistants (e.g., Codex, Claude, Cursor, Sweep AI) in understanding our practices, contributing safe and idiomatic code, and navigating the project effectively.

AI agents must adhere to the instructions herein when reading, writing, testing, or committing code.

---

## 🔍 Project Overview

`go-countries` is a Go library that provides a complete list of ISO-3166 countries and related metadata. It ships functions for retrieving countries by name or code and includes a generator that builds the country data from JSON. The project emphasizes simplicity and reliability and is continuously validated by linting and automated tests.

* Country data lives under `data/` and is transformed into Go structs by running `go generate ./generate/...`.
* The core package lives in the repository root (`countries.go`, `countries_data.go`).
* Examples and tests accompany the source to demonstrate usage and maintain correctness.

---

## 📁 Directory Structure

| Directory   | Description
| ----------- | ------------------------------------------------------------------------
| `data/`     | Raw JSON datasets used for code generation
| `generate/` | Code generation utility that produces `countries_data.go`
| `examples/` | Example program demonstrating package usage
| `.github/`  | Issue templates, workflows, and community documentation
| `.make/`    | Shared Makefile targets used by `Makefile`
| `.` (root)  | Source files, tests, and generated code for the `countries` package

---

## 🪨 Contribution Workflow

1. Format code with `go fmt` and ensure `golangci-lint` passes. Use `make lint` if unsure.
2. Run `make test` (or `make test-ci`) before committing to execute vet, lint, and tests.
3. Do **not** manually edit `countries_data.go`; run `go generate ./generate/...` if the data needs updating.
4. Add or update unit tests for any new functionality or bug fix.
5. Commit messages should be concise yet descriptive.
6. When opening a pull request, assign it to **mrz1836**.

---

## 🛠 Makefile Overview

The repository's `Makefile` includes reusable targets from `.make/common.mk` and
`.make/go.mk`. The root file exposes a few high-level commands while the files
under `.make` contain the bulk of the build logic.

`common.mk` provides utility tasks for releasing with GoReleaser, tagging
releases, and updating the releaser tool. It also offers the `diff` and `help`
commands used across projects.

`go.mk` supplies Go-specific helpers for linting, testing, generating code,
building binaries, and updating dependencies. Targets such as `lint`, `test`,
`test-ci`, and `coverage` are defined here and invoked by the root `Makefile`.

Use `make help` to view the full list of supported commands.

---

