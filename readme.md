# Attendence Time Clocker (for slack)

This application is meant to be used in conjunction with slack clock in/out tools.

This application is meant to connect to slack as a user at set random times in the
morning and at night and send a message to a particular channel.

This application requires the following setup:

- Create an app in slack
- Gather required startup parameters
- Schedule application to auto start on your computer

## Creating an app in slack

1. Navigate to https://api.slack.com/apps/
2. Click `Create New App` button, Select `From an app Manifest`, then choose your workspace and fill in the app manifest as below.
```yaml
display_information:
  name: AttendenceTC
oauth_config:
  scopes:
    user:
      - chat:write
settings:
  org_deploy_enabled: false
  socket_mode_enabled: false
  token_rotation_enabled: false
```
3. Once the app has been created from the `Basic Information` page click `Install to workspace` and `Accept`.

## Gather Required Startup Parameters

1. From your Slack App page go to the `OAuth & Permissions` page and record your `User OAuth Token`
2. Go back to Slack and go the channel you would like to post messages to.
3. Open the channel menu at the top of the channel and copy the `ChannelID`

## Schedule Application to start on your computer.

This will depend on your operating system please look for the way to do this on your own.  The app should be run by the scheduler with the following options.

### Minimum Options

```
attendance-tc -slackToken <YOUR_TOKEN> -slackChannelID <CHANNEL_ID>
```

### Full Options

HOUR here is the hour of the day for which the `start` work or `end` work can apply. This defaults to 8-11 for start time and 18-21 for end time.

```
attendance-tc -slackToken <YOUR_TOKEN> -slackChannelID <CHANNEL_ID> -startMin <HOUR> -startMax <HOUR> -endMin <HOUR> -endMax <HOUR>
```

## Limitations

On execution this application does not schedule a `start` or `end` message in the event that the `startMax`-1 or `endMax`-1 is greater than the current hour of the day respectively.  However it will schedule the messages for tomorrow at 00:00.  This might cause a bug where a message doesn't get sent if you start it halfway through the min/max range and the randomly picked value is older than the current time.  This only affects the initial startup afterwards all messages will be scheduled properly at 00:00.

If the application fails or you shutdown/restart your computer the message will not go out.
