# kaalin

CLI tool for the Karakalpak language. Script conversion, numbers to words, letter casing.

## Install

```bash
brew tap dontbeidle/tap
brew install kaalin
```

## Convert

Latin to Cyrillic:

```bash
kaalin convert "Assalawma áleykum"
# Ассалаўма әлейкум
```

Cyrillic to Latin:

```bash
kaalin convert "Ассалаўма әлейкум"
# Assalawma áleykum
```

Explicit direction:

```bash
kaalin convert --to-cyr "Sálem"
kaalin convert --to-lat "Сәлем"
```

From file:

```bash
kaalin convert -f document.txt -o converted.txt
kaalin convert -f document.txt --in-place
```

Pipe:

```bash
echo "Sálem" | kaalin convert
```

## Number to words

```bash
kaalin number 123
# bir júz jigirma úsh

kaalin number 123 --cyr
# бир жүз жигирма үш

kaalin number 12.75
# on eki pútin júzden jetpis bes

kaalin number -5
# minus bes
```

## Case

```bash
kaalin case upper "sálem álem"
# SÁLEM ÁLEM

kaalin case lower "SÁLEM ÁLEM"
# sálem álem
```

## Shell completion

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
| `--quiet` | `-q` | Print result only |
