# telegrammy

telegrammy is a simple Telegram bot that can be configured via YAML. It supports
response jobs (triggered by messages) and cron jobs (executed on a schedule).
Any messages that do not match a configured trigger are forwarded to ChatGPT,
and the context of the conversation is maintained.

## Configuration

telegrammy is configured using environment variables and a YAML file.

### Environment Variables

| Variable                        | Description                                      |
|---------------------------------|--------------------------------------------------|
| `TELEGRAMMY_CONFIG_PATH`        | Path to the YAML configuration file              |
| `TELEGRAM_CHAT_ID`              | Telegram chat ID for the bot                     |
| `TELEGRAM_BOT_TOKEN`            | Telegram bot API token                           |
| `OPENAI_API_KEY`                | OpenAI API key for ChatGPT usage                 |
| `CHAT_GPT_CONVERSATION_PATH`    | Path to the ChatGPT conversation file            |

### YAML Configuration

Both `responseJobs` and `cronJobs` follow a similar structure, differing only in
their trigger mechanism:

- `responseJobs` are triggered by specific messages
- `cronJobs` are executed on a schedule

| Parameter                           | Description                                                                                      | Default                           |
|-------------------------------------|--------------------------------------------------------------------------------------------------|-----------------------------------|
| `shellPath`                         | Path to the shell used for executing commands                                                    | _required_                        |
| `pollInterval`                      | Interval in seconds for polling new messages                                                     | _required_                        |
| `responseJobs[].trigger`            | Message that triggers the response job                                                           | _required_ (for each responseJob) |
| `responseJobs[].message`            | Message sent by the bot                                                                          | _required_ (for each responseJob) |
| `responseJobs[].parseMode`          | Formatting mode (e.g., HTML, MarkdownV2)                                                         | HTML                              |
| `responseJobs[].command`            | Shell command executed when triggered (the output can be inserted into the `message` using `%s`) | -                                 |
| `responseJobs[].escapeCharacters[]` | Characters that should be escaped in the response                                                | -                                 |
| `cronJobs[].schedule`               | Cron schedule for job execution                                                                  | _required_ (for each cronJob)     |
| `cronJobs[].message`                | Message sent by the bot                                                                          | _required_ (for each cronJob)     |
| `cronJobs[].parseMode`              | Formatting mode (e.g., HTML, MarkdownV2)                                                         | -                                 |
| `cronJobs[].command`                | Shell command executed when triggered (the output can be inserted into the `message` using `%s`) | -                                 |
| `cronJobs[].escapeCharacters[]`     | Characters that should be escaped in the response                                                | -                                 |

For an example YAML configuration, see [example_config/config.yaml](./example_config/config.yaml).

## Build and run

The simplest way is to install [Task](https://taskfile.dev/) and use the [Taskfile.yaml](./Taskfile.yaml)
in this repository.

## Build a new release

The [release.yaml](./.github/workflows/release.yaml) workflow is triggered by
tags with a leading "v" (e.g. `v1.0.0`) and builds and uploads a new Docker image
of telegrammy. For an even cleaner process, I recommend creating the new tag via
Github's [release page](https://github.com/kymzky/telegrammy/releases/new).
