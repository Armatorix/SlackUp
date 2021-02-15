# SlackUp

StandUp report generator - no one is listening, so why bother?

## Installation

1. `go get github.com/Armatorix/SlackUp`

2. Get slack OAuth token
   - go to
   - On https://api.slack.com/ : Your apps -> Create new app -> Install App -> copy oAuth token to config file
3. Set up write permissions
   - On you app page: OAuth & Permissions -> (Scroll down) -> User Token Scopes -> Add an OAuth Scope -> `channels:write`
4. [JUST DO IT](https://www.youtube.com/watch?v=ZXsQAXx_ao0&ab_channel=MotivaShian)
   - `SlackUp gen` and watch your channel :rocket:

## Roadmap

- service for running daily
- gaming mode - reponse to anyone with msg that hardly working on ticket
