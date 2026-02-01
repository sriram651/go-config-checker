# go-config-checker

A small CLI tool to validate JSON config files against a simple schema.

## What it does

* Checks for missing required keys
* Validates value types (`string`, `number`, `bool`)
* Exits non-zero when config is invalid

## Schema format

```json
{
  "required": {
    "PORT": "number",
    "ENV": "string",
    "DEBUG": "bool"
  }
}
```

## Usage

```bash
go run . --config config.json --schema schema.json
```

## Exit codes

* `0` → config is valid
* `1` → config is invalid
* `2` → tool error (file read / parse failure)
