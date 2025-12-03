# gg — the 2-letter agent-native git client

`gg` turns any public GitHub repository into an instant, chainable, code-execution MCP — with proven 98.7% token reduction.

```bash
gg maaza      # Maaza v1.2 MCP (9.6M, 62.9% adversarial, 39ms)
gg .          # current repo MCP
gg user/repo  # any public repo MCP
gg ask "…"    # (coming) Maaza writes code → opens PR → gg approve
```

## Why this feels different

|  | Traditional GitHub schema | gg + code-execution MCP |
|--|--------------------------|-------------------------|
| **Tokens** | 1,800 – 2,400 tokens | 38 – 62 tokens |
| **Data** | Full file contents sent | Only tiny results |
| **Commands** | 400-token bash calls | One `gg run test` |

**98.7% reduction** is Anthropic's measured result when the model executes code instead of receiving file contents.

## The quiet evolution

- GitHub remains the source of truth for over 100 million repositories.
- gg simply becomes the fastest, lightest way agents speak to it.

No competition.
Just natural progress.

## Install (30 seconds)

```bash
# macOS / Linux / Windows
curl -L https://gg.sh | sh

# or via Go
go install github.com/ggdotdev/gg@latest
```

## Today — v0.1

- ✅ `gg .` → current repo MCP
- ✅ `gg maaza` → Maaza Orchestrator v1.2 MCP
- ✅ `gg user/repo` → any public repo MCP
- ✅ Works immediately with Claude Desktop, Cursor, and every MCP client
- ✅ Uses GitHub's official MCP server under the hood (free & open-source)

**`gg ask` + `gg approve`** powered by Maaza v1.2 land in the next release.

## Positioning

| Layer | Project | Ownership |
|-------|---------|-----------|
| **Protocol** | MCP | Anthropic (open spec) |
| **App Store** | MCPBodega | CycleCore Technologies |
| **Killer App** | gg | Independent OSS — ggdotdev |

- Default backend = MCPBodega (privacy-first hosting)
- Works with any MCP server, including GitHub's own.

## Contributing

This is a community project. PRs are loved.

```bash
git clone https://github.com/ggdotdev/gg
cd gg && go build -o gg
./gg maaza   # should feel like magic
```

---

The bodega cat already ran `gg init` 😼

We're just getting started.
