# README

## About

This is the official Wails Vue template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Yangfan UI Client

This project is the UI execution node client for Yangfan.

### Tech Stack
- **Frontend**: Vue 3 + Arco Design
- **Backend**: Go (Wails)
- **Integration**: Local `httprunner` module

### Setup
1. Ensure `wails` CLI is installed.
2. Run `wails dev` to start the application.

### Dependencies
- The project depends on the local `../httprunner` module.
- `go.work` in the parent directory handles the workspace configuration.
