# capra
Convert between JSON, YAML and TOML.

## Features
- Pretty output with syntax highlighting
- Lots highlighting style (using [chroma](https://github.com/alecthomas/chroma?tab=readme-ov-file))

## Intallation
Run `go install`:
```console
go install github.com/9yokuro/capra@latest
```

Or Download the binary from the [release page](https://github.com/9yokuro/capra/releases)

## Usage
```console
capra --help

Usage of capra:
  -f, --format string   Set the output format (default "json")
  -H, --no-highlight    Disable syntax highlighting
  -s, --style string    Set the highlight style (default "monokai")
  -v, --version         Print version
```

To output converted data to stdout:
```console
capra --format=json foo.yaml bar.toml
```

You can use pipe:
```console
echo '{ "foo": "bar" }' | capra --format=yaml

foo: bar
```

> [!NOTE]
> When you redirect to a file or use capra with another software such as jq, you **MUST** specify the option --no-highlight.

To write converted data to a file:
```console
capra --format=json --no-highlight foo.yaml bar.toml > baz.json
```

With [jq](https://github.com/jqlang/jq):
```console
echo '{ "foo": [ { "bar": "baz" } ] }' | capra --format=json --no-highlight | jq -r '.foo[].bar'

baz
```

## Available format
- JSON
- YAML
- TOML

## Available style
See [Chroma Style Gallery](https://xyproto.github.io/splash/docs).

## License
This project is licensed under the [MIT License](https://github.com/9yokuro/capra/blob/main/LICENSE).
