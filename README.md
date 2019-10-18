# freq

Command-line utility application for displaying frequencies as histograms or bar charts, from the standard input.

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

Show distribution of word length in a document:

```bash
cat file.txt |tr -d '[:punct:]' | tr '[:upper:]' '[:lower:]' | tr '[:space:]' '\n' | xargs -I'%' -n1 sh -c "echo % | wc -m" | freq -i
```
