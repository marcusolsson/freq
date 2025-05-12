# freq

Command-line utility application for displaying frequencies as histograms or bar charts, from the standard input.

For text input, `freq` returns a bar chart:

```shell
$ cat file.txt | freq
bar  ██████████▏ 1
foo  ████████████████████▏ 2
test ██████████████████████████████▏ 3
```

For numerical datasets, add the `--histogram` option to display a frequency distribution:

```shell
$ cat dataset.txt | freq --histogram --justify
  2-3 █████ 28
  3-4 █████████████████████████████▌ 167
  4-5 ██████████████████████████████▏ 170
  5-6 ████████████████████████████▌ 161
  6-7 █████████████ 73
  7-8 ██████████▌ 59
  8-9 █████████▎ 52
 9-10 ████▋ 26
10-11 ███▏ 17
11-12 ██▊ 15
```

## Installation

```bash
go install github.com/marcusolsson/freq@latest
```

## Usage

Show top contributors to a Git repository:

```bash
git --no-pager log --format='%aN' | freq
```

Show most frequently used words in a document:

```bash
cat file.txt | tr -d '[:punct:]' | tr '[:upper:]' '[:lower:]' | tr '[:space:]' '\n' | freq
```

Show distribution of word lengths in a document:

```bash
cat file.txt | tr -d '[:punct:]' | tr '[:upper:]' '[:lower:]' | tr '[:space:]' '\n' | xargs -I'%' -n1 sh -c "echo % | wc -m" | freq --histogram
```

Show distribution of pull request age for a GitHub project, using [hub](https://hub.github.com/):

```
hub pr list -s opened -f '%ct%n' | xargs -n1 -I'{}' sh -c 'echo $(($(date +%s)-{}))' | freq --histogram --buckets=20
```
