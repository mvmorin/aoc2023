# Advent of Code 2023
Written in Go to learn

Each day is a separate package in this one module. Each solution to each day (package) can be run as a test, for
instance with `go test ./day01 -v -count 1 -run Test1` where `-v` means verbose to show stdout, `-count 1` means to run
once to avoid caching, and `-run <regex>` is used to select specific day. Instead of a path, the package/day can be
specified with the package identifier, e.g., `github.com/mvmorin/aoc2023/day01`, but that is a lot to write.
Alternatively, they main function of the module will run all solutions for all days.
