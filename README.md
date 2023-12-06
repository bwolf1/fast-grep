# fast-grep

A faster grep tool!

[GNU grep](https://www.gnu.org/software/grep/) is a command line tool for
finding matching text patterns. GNU grep can operate on shell input or file
input, and is particularly useful for highlighting matching text patterns within
a large body of text.

fast-grep re-implements the GNU grep tool using Go's built in concurrency to
make fast-grep faster.

## building

```shell
go build
```

## usage

```text
fast-grep search_pattern file_path [flags]
```

## flags

```text
--config string   config file (default is $HOME/.fast-grep.yaml)
-h, --help            help for fast-grep
-t, --toggle          Help message for toggle
-v, --version         version for fast-grep
```

## example

```shell
./fast-grep et testfile.txt
```

### result

![image info](./images/readme_results.png)
