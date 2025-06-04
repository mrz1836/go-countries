# AGENTS.md

## 🎯 Purpose & Scope

This file defines the **baseline standards, workflows, and structure** for *all contributors and AI agents* operating within this repository. It serves as the root authority for engineering conduct, coding conventions, and collaborative norms.

It is designed to help AI assistants (e.g., Codex, Claude, Cursor, Sweep AI) and human developers alike understand our practices, contribute clean and idiomatic code, and navigate the codebase confidently and effectively.

> Whether reading, writing, testing, or committing code, **you must adhere to the rules in this document.**

Additional `AGENTS.md` files **may exist in subdirectories** to provide more contextual or specialized guidance. These local agent files are allowed to **extend or override** the root rules to fit the needs of specific packages, services, or engineering domains—while still respecting the spirit of consistency and quality defined here.

---

## 🔍 Project Overview

`go-countries` is a Go library that provides a complete list of ISO-3166 countries and related metadata. It ships functions for retrieving countries by name or code and includes a generator that builds the country data from JSON. The project emphasizes simplicity and reliability and is continuously validated by linting and automated tests.

* Country data lives under `data/` and is transformed into Go structs by running `go generate ./generate/...`.
* The core package lives in the repository root (`countries.go`, `countries_data.go`).
* Examples and tests accompany the source code, demonstrating usage and maintain correctness.

---

## 📁 Directory Structure
| Directory   | Description                                                         |
|-------------|---------------------------------------------------------------------|
| `data/`     | Raw JSON datasets used for code generation                          |
| `generate/` | Code generation utility that produces `countries_data.go`           |
| `examples/` | Example program demonstrating package usage                         |
| `.github/`  | Issue templates, workflows, and community documentation             |
| `.make/`    | Shared Makefile targets used by `Makefile`                          |
| `.` (root)  | Source files, tests, and generated code for the `countries` package |

---

### 📚 Related Governance Documents

For more detailed guidance and supporting documentation, refer to the following project-level resources:

* `CITATION.cff` — Metadata for citing this project; GitHub uses it to render citation information
* `CODEOWNERS` - Ownership of the repository and various directories
* `CODE_OF_CONDUCT.md` — Expected behavior and enforcement
* `CODE_STANDARDS.md` — Style guides and best practices
* `CONTRIBUTING.md` — Guidelines for contributing code, issues, and ideas
* `README.md` — Project overview, goals, and setup instructions
* `SECURITY.md` — Vulnerability reporting and security process
* `SUPPORT.md` — How to get help or request new features

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

## 🧪 Development, Testing & Coverage Standards

All contributors—human or AI—must follow these standards to ensure high-quality, maintainable, and idiomatic Go code throughout the project.

### 🛠 Formatting & Linting

Code must be cleanly formatted and pass all linters before being committed.

```bash
go fmt ./...
goimports -w .
golangci-lint run
go vet ./...
```

> Refer to `.golangci.json` for the full set of enabled linters.

Editors should honor `.editorconfig` for indentation and whitespace rules, and
Git respects `.gitattributes` to enforce consistent line endings across
platforms.

### 🧪 Testing Standards

We use the `testify` suite for unit tests. All tests must follow these conventions:

* Name tests using the pattern: `TestFunctionName_ScenarioDescription`
* Use `testify/assert` for general assertions
* Use `testify/require` for:
    * All error or nil checks
    * Any test where failure should halt execution
    * Any test where a pointer or complex structure is required to be used after the check
* Use `require.InDelta` or `require.InEpsilon` for floating-point comparisons
* Prefer **table-driven tests** for clarity and reusability
* Use subtests (`t.Run`) to isolate and describe scenarios
* **Do not use** `t.Parallel()` unless explicitly testing concurrent behavior
* Avoid flaky, timing-sensitive, or non-deterministic tests

Run tests locally with:

```bash
go test ./...
```

> All tests must pass in CI prior to merge.

### 📈 Code Coverage

* Code coverage thresholds and rules are defined in `codecov.yml`
* Aim to provide meaningful test coverage for all new logic and edge cases
* Avoid meaningless coverage (e.g., testing getters/setters or boilerplate)

---

## ✍️ Naming Conventions

Follow Go naming idioms and the standards outlined in [Effective Go](https://go.dev/doc/effective_go):

### Packages

* Short, lowercase, one-word (e.g., `auth`, `rpc`, `block`)
* Avoid `util`, `common`, or `shared`
* Exception: standard lib wrapper like `httputil`
* Must have a clear concise package comment in a .go file with the same name as the package

### Files

* Naming using: snake_case (e.g., `block_header.go`, `test_helper.go`)
* Go file names are lowercase
* Test files: `_test.go`
* Generated: use prefix `zz_` and annotate with a `// Code generated...` header

### Functions & Methods

* `VerbNoun` naming (e.g., `CalculateHash`, `ReadFile`)
* Constructors: `NewXxx` or `MakeXxx`
* Getters: field name only (`Name()`)
* Setters: `SetXxx(value)`

### Variables

* Exported: `CamelCase` (e.g., `HTTPTimeout`)
* Internal: `camelCase` (e.g., `localTime`)
* Idioms: `i`, `j`, `err`, `tmp` accepted

### Interfaces

* Single-method: use `-er` suffix (e.g., `Reader`, `Closer`)
* Multi-method: use role-based names (e.g., `FileSystem`, `StateManager`)

---

## 📘 Commenting Standards

Great engineers write great comments. You’re not here to state the obvious—you’re here to document decisions, highlight edge cases, and make sure the next dev (or AI) doesn’t repeat your mistakes.

### 🧠 Guiding Principles

* **Comment the "why", not the "what"**

  > The code already tells us *what* it’s doing. Your job is to explain *why* it’s doing it that way—especially if it’s non-obvious, nuanced, or a workaround.

* **Explain side effects, caveats, and constraints**

  > If the function touches global state, writes to disk, mutates shared memory, or makes assumptions—write it down.

* **Don’t comment on broken code—fix or delete it**

  > Dead or disabled code with TODOs are bad signals. If it’s not worth fixing now, delete it and add an issue instead.

* **Your comments are part of the product**

  > Treat them like UX copy. Make them clear, concise, and professional. You’re writing for peers, not compilers.

---

### 🔤 Function Comments (Exported)

Every exported function **must** include a Go-style comment that:

* Starts with the function name
* States its purpose clearly
* Documents:
  * **Steps**: Include if the function performs a non-obvious sequence of operations.
  * **Parameters**: Always describe all parameters when present.
  * **Return values**: Document return types and their meaning if not trivially understood.
  * **Side effects**: Note any I/O, database writes, external calls, or mutations that aren't local to the function.
  * **Notes**: Include any assumptions, constraints, or important context that the caller should know.


Here is an example:

```go
// FunctionName does [what this function does] in [brief context].
//
// This function performs the following steps: [if applicable, describe the main steps in a bullet list]
// - [First major action or check performed]
// - [Second action or branching logic explained, if relevant]
//    - [Details about possible outcomes or internal branching]
// - [Additional steps with sub-bullets as needed]
// - [Final steps and cleanup, if applicable]
//
// Parameters:
// - ctx: [Purpose of context in this function]
// - paramName: [Explanation of each parameter and what it controls or affects]
//
// Returns:
// - [What is returned; error behavior if any]
//
// Side Effects:
// - [Any side effects, such as modifying global state, writing to disk, etc.]
//
// Notes:
// - [Caveats, assumptions, or recommendations—e.g., transaction usage, concurrency, etc.]
// - [Any implicit contracts with the caller or system constraints]
// - [Mention if this function is part of a larger workflow or job system]
```

---

### 📦 Package-Level Comments

* Each package **must** include a package-level comment in a file named after the package (e.g., `auth.go` for package `auth`).
* If no logical file fits, add a `doc.go` with the comment block.
* Use it to explain:
    * The package purpose
    * High-level API boundaries
    * Expected use-cases and design notes

---

### 🧱 Inline Comments

Use inline comments **strategically**, not excessively.

* Use them to explain "weird" logic, workarounds, or business rules.
* Prefer **block comments (`//`)** on their own line over trailing comments.
* Avoid obvious noise:

🚫 `i++ // increment i`
✅ `// Skip empty rows to avoid panic on CSV parse`

---

### ⚙️ Comment Style

* Use **complete sentences** with punctuation.
* Keep your tone **precise, confident, and modern**—you're not writing a novel, but you're also not writing legacy COBOL.
* Avoid filler like "simple function" or "just does X".
* Don’t leave TODOs unless:
    * They are immediately actionable
    * (or) they reference an issue
    * They include a timestamp or owner

---

### 🧬 AI Agent Directives

If you're an AI contributing code:

* Treat your comments like commit messages—**use active voice, be declarative**
* Use comments to **make intent explicit**, especially for generated or AI-authored sections
* Avoid hallucinating context—if you're unsure, omit or tag with `// AI: review this logic`
* Flag areas of uncertainty or external dependency (e.g., "// AI: relies on external config structure")

---

### 🔥 Comment Hygiene

* Remove outdated comments aggressively.
* Keep comments synced with refactoring.
* Use `//nolint:<linter> // message` only with clear, justified context and explanation.

---

## 📝 Modifying Markdown Documents

Markdown files (e.g., `README.md`, `AGENTS.md`, `CONTRIBUTING.md`) are first-class citizens in this repository. Edits must follow these best practices:

* **Write with intent** — Be concise, clear, and audience-aware. Each section should serve a purpose.
* **Use proper structure** — Maintain consistent heading levels, spacing, and bullet/number list formatting.
* **Full Table Borders** — Use full borders for tables to ensure readability across different renderers.
* **Table Border Spacing** — Make sure tables have appropriate spacing for clarity.
* **Preserve voice and tone** — Match the professional tone and style used across the project documentation.
* **Preview before committing** — Always verify rendered output locally or in a PR to avoid broken formatting.
* **Update references** — When renaming files or sections, update internal links and the table of contents if present.

> Markdown updates should be treated with the same care as code—clean, purposeful, and reviewed.

---

## 🚨 Error Handling (Go)

* Always check errors

```go
if err != nil {
  return err
}
```

* Prefer `errors.New()` over `fmt.Errorf`
* Use custom error types sparingly
* Avoid returning ambiguous errors; provide context

---

## 🔀 Commit & Branch Naming Conventions

Clear history ⇒ easy maintenance. Follow these rules for every commit and branch.

### 📌 Commit Message Format

```
<type>(<scope>): <imperative short description>

<body>  # optional, wrap at 72 chars
```

* **`<type>`** — `feat`, `fix`, `docs`, `test`, `refactor`, `chore`, `build`, `ci`
* **`<scope>`** — Affected subsystem or package (e.g., `api`, `countries`, `deps`). Omit if global.
* **Short description** — ≤ 50 chars, imperative mood ("add pagination", "fix panic")
* **Body** (optional) — What & why, links to issues (`Closes #123`), and breaking‑change note (`BREAKING CHANGE:`)

**Examples**

```
feat(countries): add Alpha-3 lookup by numeric code
fix(generator): handle malformed JSON input gracefully
docs(README): improve installation instructions
```

> Commits that only tweak whitespace, comments, or docs inside a PR may be squashed; otherwise preserve granular commits.

### 🌱 Branch Naming

| Purpose            | Prefix      | Example                            |
|--------------------|-------------|------------------------------------|
| Feature            | `feat/`     | `feat/pagination-api`              |
| Bug Fix            | `fix/`      | `fix/country-code-off-by-one`      |
| Documentation      | `docs/`     | `docs/agents-commenting-standards` |
| Refactor / Cleanup | `refactor/` | `refactor/remove-dead-code`        |
| Tests              | `test/`     | `test/generator-edge-cases`        |
| Chore / Meta       | `chore/`    | `chore/upgrade-go-1.23`            |
| Hotfix (prod)      | `hotfix/`   | `hotfix/rollback-broken-deploy`    |
| Prototype / Spike  | `proto/`    | `proto/iso3166-expansion`          |

* Use **kebab‑case** after the prefix.
* Keep branch names concise yet descriptive.
* PR titles should mirror the branch’s purpose (see [✅ Pull Request Conventions](#-pull-request-conventions)).

> Mergify and CI rely on these prefixes for auto‑labeling and workflow routing—stick to them.

---

## ✅ Pull Request Conventions

Pull Requests—whether authored by humans or AI agents—must follow a consistent structure to ensure clarity, accountability, and ease of review.

### 🔖 Title Format

```
[Subsystem] Imperative and concise summary of change
```

Examples:

* `[API] Add pagination to client search endpoint`
* `[DB] Migrate legacy rate table schema`
* `[CI] Remove deprecated GitHub Action for testing`

> Use the imperative mood ("Add", "Fix", "Update") to match the style of commit messages and changelogs.

---

### 📝 Pull Request Description

Every PR must include the following **four** sections in the description:

#### 1. **What Changed**

> A clear, bullet‑pointed or paragraph‑level summary of the technical changes.

#### 2. **Why It Was Needed**

> Context or motivation behind the change. Reference related issues, discussions, or bugs if applicable.

#### 3. **Testing Performed**

> Document:
>
> * Test suites run (e.g., `TestCreateOriginationAccount`)
> * Edge cases covered
> * Manual steps that were taken (if any)

#### 4. **Impact / Risk**

> Call out:
>
> * Breaking changes
> * Regression risk
> * Performance implications
> * Changes in developer experience (e.g., local dev setup, CI time)

---

### 💡 Additional PR Guidelines

* Link related issues with keywords like `Closes #123` or `Fixes #456` if there is a known issue.
* Keep PRs focused and minimal. Prefer multiple small PRs over large ones when possible.
* Use draft PRs early for feedback on in progress work.
* Releases are deployed using **goreleaser**.
* Rules for the release build are located in `.goreleaser.yml` and executed via `.github/workflows/release.yml`.

---

## 🚀 Release Workflow & Versioning

We follow **Semantic Versioning (✧ SemVer)**:  
`MAJOR.MINOR.PATCH` → `1.2.3`

| Segment   | Bumps When …                          | Examples        |
|-----------|---------------------------------------|-----------------|
| **MAJOR** | Breaking API change                   | `1.0.0 → 2.0.0` |
| **MINOR** | Back‑compatible feature / enhancement | `1.2.0 → 1.3.0` |
| **PATCH** | Back‑compatible bug fix / docs        | `1.2.3 → 1.2.4` |

### 📦 Tooling

* Releases are driven by **[goreleaser]** and configured in `.goreleaser.yml`.
* Install locally with Homebrew (Mac):  
```bash
  brew install goreleaser
````

### 🔄 Workflow

| Step | Command                         | Purpose                                                                                            |
|------|---------------------------------|----------------------------------------------------------------------------------------------------|
| 1    | `make release-snap`             | Build & upload a **snapshot** (pre‑release) for quick CI validation.                               |
| 2    | `make tag version=X.Y.Z`        | Create and push a signed Git tag. Triggers GitHub Actions.                                         |
| 3    | GitHub Actions                  | CI runs `goreleaser release` on the tag; artifacts and changelog are published to GitHub Releases. |
| 4    | `make release` (optional local) | Manually invoke the production release if needed.                                                  |

> **Note for AI Agents:** Do not create or push tags automatically. Only the repository [codeowners](CODEOWNERS) are authorized to tag and publish official releases.

[goreleaser]: https://github.com/goreleaser/goreleaser

---

## 🏷️ Labeling Conventions (GitHub)

Labels serve as shared vocabulary for categorizing issues, pull requests, and discussions. Proper labeling improves triage, prioritization, automation, and clarity across the engineering lifecycle.

Current labels are located in `.github/labels.yml` and automatically synced into GitHub upon updating the `master` branch.

### 🎨 Standard Labels & Usage

| Label Name         | Color     | Description                                                | When to Use                                                                 |
|--------------------|-----------|------------------------------------------------------------|-----------------------------------------------------------------------------|
| `documentation`    | `#0075ca` | Improvements or additions to project docs                  | Updates to `README`, onboarding docs, usage guides, code comments           |
| `bug-P1`           | `#b23128` | **Critical bug**, highest priority, impacts all users      | Regressions, major system outages, critical service bugs                    |
| `bug-P2`           | `#de3d32` | **Moderate bug**, medium priority, affects a subset        | Broken functionality with known workaround or scoped impact                 |
| `bug-P3`           | `#f44336` | **Minor bug**, lowest priority, limited user impact        | Edge case issues, cosmetic UI glitches, legacy bugs                         |
| `feature`          | `#0e8a16` | Any new **major feature or capability**                    | Adding new API, CLI command, UI section, or module                          |
| `hot-fix`          | `#b60205` | Time-sensitive or production-impacting fix                 | Used with `bug-P1` or urgent code/config changes that must ship immediately |
| `idea`             | `#cccccc` | Suggestions or brainstorming candidates                    | Feature ideas, process improvements, early-stage thoughts                   |
| `prototype`        | `#d4c5f9` | Experimental work that may be unstable or incomplete       | Spike branches, POCs, proof-of-concept work                                 |
| `question`         | `#cc317c` | A request for clarification or feedback                    | Use for technical questions, code understanding, process queries            |
| `test`             | `#c2e0c6` | Changes to tests or test infrastructure                    | Unit tests, mocks, integration tests, CI coverage enhancements              |
| `ui-ux`            | `#fbca04` | Frontend or user experience-related changes                | CSS/HTML/JS updates, UI behavior tweaks, design consistency                 |
| `chore`            | `#006b75` | Low-impact, internal tasks                                 | Dependency bumps, code formatting, comment fixes                            |
| `update`           | `#006b75` | General updates not tied to a specific bug or feature      | Routine code changes, small improvements, silent enhancements               |
| `refactor`         | `#ffa500` | Non-functional changes to improve structure or readability | Code cleanups, abstraction, splitting monoliths                             |
| `automerge`        | `#fef2c0` | Safe to merge automatically (e.g., from CI or bot)         | Label added by automation or trusted reviewers                              |
| `work-in-progress` | `#fbca04` | Not ready to merge, actively under development             | Blocks `automerge`, signals in-progress discussion or implementation        |
| `stale`            | `#c2e0c6` | Inactive, obsolete, or no longer relevant                  | Used for automated cleanup or manual archiving of old PRs/issues            |

---

### 🧠 Labeling Best Practices

* Apply labels at the time of PR/issue creation, or during triage.
* Use **only one priority label** (`bug-P1`, `P2`, `P3`) per item.
* Combine labels as needed (e.g., `feature` + `ui-ux` + `test`).
* Don’t forget to remove outdated labels (e.g., `work-in-progress` → after merge readiness).

---

## 🤖 Mergify Automation

This repository uses [Mergify](https://docs.mergify.com) to automate and streamline pull request workflows. Mergify enables intelligent merging, auto-labeling, and CI-based gatekeeping to ensure high-quality contributions with reduced manual overhead.

### 🔧 Core Behaviors

* **Automatic Merging for Dependabot PRs**
  Minor version bumps from Dependabot are automatically reviewed and merged once all required checks pass.
  Major version bumps require manual review and trigger an alert.

* **Standard Approval-Based Merging**
  Pull requests with:

    * At least one approval
    * No requested changes
    * Passing required checks
      are merged automatically unless marked as a draft or labeled `work-in-progress`.

* **Auto-Labeling by Branch Prefix**
  Labels such as `feature`, `bug-P3`, `test`, `idea`, and `hot-fix` are added automatically based on the branch name pattern (e.g., `feat/`, `fix/`, `hotfix/`, `test/`, etc.).

* **Automated Housekeeping**

    * Stale PRs (older than 21 days without activity) are closed with a `stale` label.
    * Merged branches are deleted automatically.
    * New contributors are welcomed with a friendly message.

### 🏷️ Automation Notes

* The `automerge` label allows for non-interactive merging when all conditions are satisfied.
* The `work-in-progress` label or a `wip/` branch name disables auto-merge.
* All required CI jobs must succeed before merging is permitted.

---

## 🧩 CI & Validation

CI automatically runs on every PR to verify:

* Formatting (`go fmt` and `goimports`)
* Linting (`golangci-lint run`)
* Tests (`go test ./...`)
* This codebase uses GitHub Actions; test workflows reside in `.github/workflows/run-tests.yml`

Failing PRs will be blocked. AI agents should iterate until CI passes.

---

## 🔐 Dependency Management

Dependency hygiene is critical for security, reproducibility, and developer experience. Follow these practices to ensure our module stays stable, up to date, and secure.

### 📦 Module Management

* All dependencies must be managed via **Go Modules** (`go.mod`, `go.sum`)

* After adding, updating, or removing imports, run:

  ```bash
  go mod tidy
  ```

* Periodically refresh dependencies with:

  ```bash
  go get -u ./...
  ```

> Avoid unnecessary upgrades near release windows—review major version bumps carefully for breaking changes.

### 🛡️ Security Scanning

* Use [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck) to identify known vulnerabilities:
```bash
  govulncheck ./...
```

* Run via make command: 
```bash
  make govulncheck
```

* Address critical advisories before merging changes into `master`

* Document any intentionally ignored vulnerabilities with clear justification and issue tracking

### 📁 Version Control

* Never manually edit `go.sum`
* Do not vendor dependencies; we rely on modules for reproducibility
* Lockstep upgrades across repos (when applicable) should be coordinated and noted in PRs

> Changes to dependencies should be explained in the PR description and ideally linked to the reason (e.g., bug fix, security advisory, feature requirement).

---

## 🛡️Security Considerations & Vulnerability Reporting

Security is a first-class requirement. If you discover a vulnerability—no matter how small—follow our responsible disclosure process:

* **Do not** open a public issue or pull request.
* Follow the instructions in [`SECURITY.md`](SECURITY.md).
* Include:
  * A clear, reproducible description of the issue
  * Proof‑of‑concept code or steps (if possible)
  * Any known mitigations or workarounds
* You will receive an acknowledgment within **72 hours** and status updates until the issue is resolved.

> For general hardening guidance (e.g., `govulncheck`, dependency pinning), see the [🔐Dependency Management](#-dependency-management) section.

---

## 🕓 Change Log (AGENTS.md)

This section tracks notable updates to `AGENTS.md`, including the date, author, and purpose of each revision. 
All contributors are expected to append entries here when making meaningful changes to agent behavior, conventions, or policies.

| Date       | Author   | Summary of Changes                                                             |
|------------|----------|--------------------------------------------------------------------------------|
| 2025-06-03 | @mrz1836 | Major rewrite: clarified commenting standards and merged scope/purpose         |
| 2025-06-03 | @mrz1836 | Combined testing and development sections; improved formatting & test guidance |
| 2025-06-03 | @mrz1836 | Enhanced dependency management practices and security scanning advice          |
| 2025-06-04 | @mrz1836 | Documented citation and configuration files for contributors                   |

> For minor edits (typos, formatting), this log update is optional. For all behavioral or structural changes, log entries are **required**.
