# MS Pocket Word text extractor
Extracts plain text paragraphs from MS Pocket Word files (*.psw).

Supports Latin and Cyrillic symbols.

Doesn't support any formatting, pictures, other character sets.
## Usage
Build with `go build pwextractor.go` then run `pwextractor <PSW file>` it will produce file with the same name and `.txt` extension
## Contribution
If you have examples of PSW files with various formatting options or character sets please let me know, I'll try to turn this simple PoC extractor into more powerful converter