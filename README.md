# Advent of Code Solutions
This are _my_ solutions to the [AdventOfCode](https://adventofcode.com) 2022 programming puzzles written in the Go programming language.

I would strongly suggest that you try to solve the puzzles before checking the solutions. 

# Setup instructions

Each user gets a different set if input data for each puzzle. To run the solutions yourself you will need first to get your session token exported.

You can get yours by logging into [AdventOfCode](https://adventofcode.com) and inspecting your cookies contents. You should have a cookie for the `.adventofcode.com` domain. Export the cookie value as:

```bash
export SESSION=536.........
```

Optionally you can install `direnv` on your terminal and then drop a `.envrc` file on the root of this git repo.

```
$ cat .envrc
export SESSION=536......
```

That way you won't need to export it manually ever again until your session expires and you need to update the session token :)

# Structure

Each puzzle has a solutions/dayN directory for the `main()` function and a `internal/dayN` for the libraries and unit tests.

Each puzzle test expresses as close as possible the given puzzle instructions. Run `go test -v` in the `internal/dayN` directory to run a particular puzzle test.

You can get _your_ solutions by running:

```
$ make dayN
```

where N is the day number. For example for the solutions for day1 you would run:

```
$ make day1
```

Each solution will show answers for Part1 and Part2.