# Advent of Go 🎄🐹

This repository is a self-guided learning journey through the Go programming language,
inspired by the **Advent of Code** format.

Each day introduces a new problem that focuses on a specific Go concept,
with:
- Clear problem descriptions
- Required function signatures
- Edge cases
- Automated test suites

The goal is to solve each problem using **idiomatic Go** and **test-driven development**.

---

## Structure

Each day lives in its own folder:

```text
dayXX/
├── README.md        # Problem description
├── solution.go     # Your implementation
└── solution_test.go

To see detailed test output:

```bash
go test -v ./...

