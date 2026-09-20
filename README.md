# GlassFM

A CLI file manager for Linux that shows you the real command behind every
action — navigate with human-readable menus, learn the shell as you go.

## Why this exists

Most beginner-friendly tools work by hiding the system underneath them. This
one doesn't.

`filemgr` lets you navigate directories and pick plain-language actions —
"Create a new directory here" instead of typing `mkdir` — but every action
also shows you the actual command it just ran. The abstraction is there to
make the CLI usable from day one, not to replace the CLI itself.

If you don't want to see what's happening underneath, tools like Dolphin or
Nautilus already do that well. This project is for people who want the
opposite: a gentler on-ramp into the terminal, not a permanent replacement
for it.

## Design principles

- **Transparent, not simplified.** Every action shows the real command
  behind it. This is not a toggle or a "beginner mode" — it's the core
  mechanic.
- **No training wheels that come off.** The tool behaves the same on day
  100 as it does on day one. Nothing fades away as you get more familiar
  with it.