# SPAN Developer Challenge Submission

## Overview

This repository is a submission for the SPAN developer challenge.

I have chosen to go with golang as this is my preferred language and ideally suited to CLI applications.

The pattern loosely followed within this repository is to use the so called [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) pattern whereby `entities` or `models` are a central part of the dependency flow which then moves outwards to more aggregating packages. Where it was more difficult to follow this pattern, sub-packages were used to ensure there were no circular dependencies effectively pushing certain parts of the concern more to the center of the dependency graph.

The CLI tool utilizes a library called [cobra](https://github.com/spf13/cobra) which has preset conventions as well as completions and ‘help’ output. Cobra is well-used and liked in the golang community and has many well known projects using it for example 

To invoke the cli you simply need to run `go run ./cli space rank <input-file(s)>`

The solution diverges a little from the described solution in the sense that if a team has zero points they do not receive the same rank number but rather have a sequential number. it was noted that in the output example when two teams had the same rank the following rank would switch back to a  sequential number. This is illustrated in the sample output below (note the jump from 3-5):

```csv
1. Tarantulas, 6 pts
2. Lions, 5 pts
3. FC Awesome, 1 pt
3. Snakes, 1 pt
5. Grouches, 0 pts
```

Essentially any teams with 0 points are simply listed alphabetically rather than given the same rank. This made more intuitive sense as an outcome.

## Coverage

The output of `go test -cover ./...` is as follows:

```bash
?       span-challenge  [no test files]
?       span-challenge/cmd      [no test files]
?       span-challenge/cmd/answer       [no test files]
?       span-challenge/cmd/ask  [no test files]
ok      span-challenge/internal/usecase 0.004s  coverage: 83.5% of statements
ok      span-challenge/pkg/config       0.002s  coverage: 100.0% of statements
?       span-challenge/pkg/errs [no test files]
?       span-challenge/pkg/fixtures     [no test files]
ok      span-challenge/pkg/models       0.005s  coverage: 100.0% of statements
ok      span-challenge/pkg/parse        0.014s  coverage: 87.5% of statements
?       span-challenge/pkg/platform     [no test files]
ok      span-challenge/pkg/serialize    0.003s  coverage: 87.5% of statements
ok      span-challenge/pkg/util 0.003s  coverage: 75.9% of statements
ok      span-challenge/pkg/validate     0.005s  coverage: 100.0% of statements

```

## CLOC
```shell
      47 text files.
      45 unique files.                              
      11 files ignored.

github.com/AlDanial/cloc v 1.82  T=0.02 s (1537.7 files/s, 97459.8 lines/s)
--------------------------------------------------------------------------------
Language                      files          blank        comment           code
--------------------------------------------------------------------------------
Go                               29            128             55           1764
XML                               6              0              0            306
Markdown                          1             28              0             57
Bourne Again Shell                1              1              0              6
--------------------------------------------------------------------------------
SUM:                             37            157             55           2133
--------------------------------------------------------------------------------

```

## Sample Output Comparison

To compare the output of this cli with the sample provided in the assignment you can run the following

```shell
go run ./cli rank pkg/fixtures/provided-example-input.csv -d file -n test.csv &&
diff test.csv pkg/fixtures/provided-example-output.csv
```

There is no difference between the two files.

## Assumptions

- The input file is a csv file but the code is written so as not to care about the file extension
- The input file is a valid csv file although some validation is done and strings are trimmed of leading and trailing spaces
- Parsing of a team score for example `Team A 2` is simplistic and I have avoided any type of regex extraction in favour of a simple split on the last space and parse the last chunk as an integer which should perform better for larger files.
- 
## Response to cautions within the assignment

- The input files are specified by the user of the CLI so any type of file join delimiter is handled by golang and the os package.
- Line endings are provided with specially named files in the platform package which only build on certain operating systems.

## Niceties

The CLI doesn't enforce the memorization of commandline args. You can provide them as a shortcut but if you don't provide any required flags, the CLI will prompt you for the values.

Log level can be adjusted with the -l flag e.g. `go run ./cli space rank -l debug <input-file(s)>`

The CLI handles any number of files provided and gives a single consolidated result

## Future Improvements

- Add validation on accepted valued for flags such as the output format flag
- Add more format options for the output e.g. tableprint where the output would be more user-friendly on the terminal
