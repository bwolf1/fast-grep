# fast-grep

## TODO

### Phase 1

- [x] Write functionality to identify patterns in a line of text
- [x] Implement a way to handle both input files and shell input
- [x] Create a function to display matched lines
- [x] Add concurrency

### Phase 2

- [x] Highlight the matched pattern with a different color
- [ ] An a `-i`, `--ignore-case` flag to perform case insensitive matching
- [ ] Add a `-p`, `--preserve-order` flag for enforcing results order
  preservation
- [ ] Add tree view
- [ ] Add flags to ignore files based on a glob pattern
- [ ] Add a `-v`, `--invert-match` flag to invert matching (show lines that do not match)
- [ ] Add a `-e pattern`, `--regexp=pattern` flag to search for matches using a regular expression
- [ ] Add the ability to get file input from a Linux pipe
