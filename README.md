# SLACKWIENER BACKEND

Slackwiener is the Nakkiservo slack backend utility server thingy thing. Yeah dog. 

Currently it just makes the files permalink_public available via slack api whenever a file is created.

## Building

### What you need.

- Go 1.9 or newer
- Git
- Working internet


### What you need to do

At the root directory: 

  $ go get
  $ make


And you're done.


## Running

The binaries are under build/
You can run them.

Configure application using the app_config.toml, make sure to include your bearer token. See below.

### Required Authorization Scopes

The backend currently requires a token with the following tokens to function properly:

- workspace admin
- channels:history
- files:read
- files:write:user

Without the files:read and files:write:user we can't make files public or get their info at all.

