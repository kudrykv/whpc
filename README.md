# Webhook Proxy Client

This application helps to route webhooks to localhost.

```
  ┌──────────────────┐            ┌──────────────────┐
  │                  │            │                  │
  │       whps       │<──webhook──│     Service      │
  │                  │            │                  │
  └──────────────────┘            └──────────────────┘
            ^
            ┃
       websocket
            ┃
┌───────────v──────────────────────────────Local machine
│ ╔══════════════════╗            ┌──────────────────┐ │
│ ║                  ║            │                  │ │
│ ║       whpc       ║──webhook──>│   Application    │ │
│ ║                  ║            │                  │ │
│ ╚══════════════════╝            └──────────────────┘ │
└──────────────────────────────────────────────────────┘
```

This is a client for the [`whps`](https://github.com/kudrykv/whps).
It connects to the server using websocket and relays incoming messages
to the `Application`.