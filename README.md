# kaalin-cli

A CLI tool for the Karakalpak language. Convert text between Latin and Cyrillic scripts, turn numbers into words, and change letter casing — all with proper Karakalpak alphabet support.

## Install

### Homebrew (macOS / Linux)

```bash
brew tap dontbeidle/tap
brew install kaalin
```

## Usage

### Convert text between scripts

```bash
# Latin to Cyrillic (auto-detected)
kaalin convert "Assalawma áleykum"
# → Ассалаўма әлейкум

# Cyrillic to Latin (auto-detected)
kaalin convert "Ассалаўма әлейкум"
# → Assalawma áleykum

# Explicit direction
kaalin convert --to-cyr "Sálem"
kaalin convert --to-lat "Сәлем"

# Pipe
echo "Sálem" | kaalin convert

# File input/output
kaalin convert -f document.txt -o converted.txt

# Edit file in place
kaalin convert -f document.txt --in-place
```

### Numbers to words

```bash
kaalin number 123
# → bir júz jigirma úsh

kaalin number 123 --cyr
# → бир жүз жигирма үш

kaalin number 12.75
# → on eki pútin júzden jetpis bes

kaalin number -5
# → minus bes
```

### Change case

```bash
kaalin case upper "sálem álem"
# → SÁLEM ÁLEM

kaalin case lower "SÁLEM ÁLEM"
# → sálem álem
```

### Shell completion

```bash
kaalin completion bash >> ~/.bashrc
kaalin completion zsh >> ~/.zshrc
kaalin completion fish > ~/.config/fish/completions/kaalin.fish
kaalin completion powershell >> $PROFILE
```

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--help` | `-h` | Show help |
| `--json` | `-j` | Output as JSON |
| `--no-color` | — | Disable colored output |
| `--quiet` | `-q` | Only print the result |

The `NO_COLOR` environment variable is also respected.

