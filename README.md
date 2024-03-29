# Secret-Santa

Secret-Santa is an application specifically designed to simplify the process of organizing gift exchanges. 
With this app, users can easily organize a gift exchange without the need for manually drawing names from a hat or pieces of paper.

## Features

 - [X] Rule sets
 - [X] Send per Email

## Installation

You only need to download the binary for your OS, from the latest release.
After extracting the ``zip`` or ``tar`` archive, you can use the binary.

## Get started

1. Setup Configuration file
2. Setup Env variables
3. Run application

### Configuration

`config.yaml`

```yaml
users: # Register Users
  - name: Clara
    email: example@domain.example
  - name: Tom
    email: example@domain.example
  - name: Jonas
    email: example@domain.example
  - name: Max
    email: example@domain.example
  - name: Lena
    email: example@domain.example
  - name: Leo
    email: example@domain.example
  - name: Jan
    email: example@domain.example

forbiddenCombinations: # Write down forbidden combinations
  - combination:
      - Max    # Max can never get Lena
      - Lena   # Lena can never get Max
  - combination:
      - Jonas
      - Jan
  - combination:
      - Clara # Clara can never get Leo or Tom
      - Leo
      - Tom

summaryEmail: example@domain.example

emailConfig:
  subject: SecretSanta 2023
  content: Hi %v, your gift goes to %v # The first %v is the mail receiver name and the second %v is the name where the gift goes to
```

### Flags
| Flag   | Description          |
|------- | -------------------- |
| `y`    | Sends emails directly  |


### Env variables

Configure email credentails.
Set as environment variables, or write in a `.env` file

```env
MAIL_SENDER=example@domain.example
MAIL_PW=123
MAIL_SMTP_SERVER=smtp.domain.example
MAIL_SMTP_SERVER_PORT=587
```
