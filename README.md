# Secret-Santa
Secret-Santa is an application specifically designed to simplify the process of organizing gift exchanges. With this app, users can easily organize a gift exchange without the need for manually drawing names from a hat or pieces of paper.

## Features

 - [X] Rule sets
 - [X] Send per Email

## Get started

1. Setup Configuration file
2. Setup Env variables
3. Run application

### Configuration
`config.yaml`
```yaml
users: # Register Users
  - name: Clara
    email: jonasbender2005@gmail.com
  - name: Tom
    email: jonasbender2005@gmail.com
  - name: Jonas
    email: jonasbender2005@gmail.com
  - name: Max
    email: jonasbender2005@gmail.com
  - name: Lena
    email: jonasbender2005@gmail.com
  - name: Leo
    email: jonasbender2005@gmail.com
  - name: Jan
    email: jonasbender2005@gmail.com

forbiddenCombinations: # Write down permitted combinations
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
emailConfig:
  subject: SecretSanta 2020
  content: Hi %v, your gift goes to %v  # The first %v is the mail receiver name and the second %v is the where the gift goes to
```

### Env variables
Set as environment variables, or write in a `.env` file
````
MAIL_SENDER=example@domain.example
MAIL_PW=123
MAIL_SMTP_SERVER=smtp.domain.example
MAIL_SMTP_SERVER_PORT=587
````