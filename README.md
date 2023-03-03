# Secret-Santa
Secret-Santa is an application specifically designed to simplify the process of organizing gift exchanges. With this app, users can easily organize a gift exchange without the need for manually drawing names from a hat or pieces of paper.

## Features

 - [X] Rule sets
 - [ ] Send per Email _(coming soon)_

## Get started

1. Setup Configuration file
2. Run application

### Configuration
`config.yaml`
```yaml
users:    # Register Users
  - name: Clara
    email: clara@domain.example
  - name: Tom
    email: tom@domain.example
  - name: Jonas
    email: jonas@domain.example
  - name: Max
    email: max@domain.example
  - name: Lena
    email: lena@domain.example
  - name: Leo
    email: leo@domain.example
  - name: Jan
    email: jan@domain.example
    
comboBans: # Write down permitted combinations 
  - combo: 
      - Max    # Max can never get Lena
      - Lena   # Lena can never get Max
  - combo:
      - Jonas
      - jan
  - combo:
      - Clara # Clara can never get Leo or Tom
      - Leo
      - Tom
```