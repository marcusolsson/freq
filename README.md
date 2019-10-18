# freq

Command-line utility application for displaying frequencies as histograms or bar charts, from the standard input.

For text input, `freq` returns a bar chart:

```shell
$ cat file.txt | freq
bar  ██████████ 1
foo  ████████████████████ 2
test ██████████████████████████████ 3
```

For numerical datasets, add the `--histogram` option to display a frequency distribution:

```shell
$ cat dataset.txt | freq --histogram --justify
  [2.0, 3.0] ████ 28
  [3.0, 4.0] █████████████████████████████ 167
  [4.0, 5.0] ██████████████████████████████ 170
  [5.0, 6.0] ████████████████████████████ 161
  [6.0, 7.0] ████████████ 73
  [7.0, 8.0] ██████████ 59
  [8.0, 9.0] █████████ 52
 [9.0, 10.0] ████ 26
[10.0, 11.0] ███ 17
[11.0, 12.0] ██ 15
```

## Installation

```bash
go get github.com/marcusolsson/freq
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
