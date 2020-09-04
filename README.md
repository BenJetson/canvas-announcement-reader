# canvas-announcement-reader

Mark all of those pesky Canvas announcements that you read in your email as
read in the web interface.

## What does this do

It will mark all announcements as read for _all_ of your active Canvas courses.
This action is permanent and irreversible.

## Why

Canvas can be configured to send emails, text messages, push notifications, and
web interface notifications when you receive an announcement.

If you're like me, you read these in your email and the little notification
badge on the web interface simply keeps going up all semester long. This just
clears those unread messages from the web interface.

Aggravatingly, Canvas does not provide a "mark all as read" button, hence this
program was created.

## Usage

This software comes with ABSOLUTELY NO WARRANTY! Use this
program at your own risk. See [LICENSE](LICENSE) for details.

Download the appropriate binary for your platform from the Releasees page or
compile it yourself like so:

```sh
go build cmd/main/main.go
```

Then simply run the binary and follow the instructions.

If you are doing this frequently, you may skip the prompts for the host and
token by setting the `CANVAS_HOST` and `CANVAS_TOKEN` environment variables.
