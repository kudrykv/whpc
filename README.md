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

# Get

```sh
go get -u github.com/kudrykv/whpc/whpc
```

# Usage

Application has two mandatory parameters: channel and route.
Channel is your random unique name on what id to operate, and
route is the path where to route requests:

```bash
./whpc -channel betazoid -route http://localhost:8080/webhook
```

3rd-app can send webhooks here:
```
https://whps.herokuapp.com/webhook/betazoid
```

Here, the `betazoid` is that channel name we made up for ourselves.
