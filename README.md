# Constellix

## This tool allows you to create or delete CNAME and A records based on domain id provided.

### Usage:
  `constellix-cli [command]`

### Available Commands:
  `create`      Create a dns records of types CNAME or A
  
  `delete`      Delete a dns record by its name
  
  `get`         Get one or all dns records/domains

### Flags:
  `-c, --command-config [string]`   Provide a config file (default "config.toml")
&nbsp;  
## Create a dns records of types CNAME or A

### Usage:
  `constellix-cli create [flags]`
  
  `constellix-cli create [command]`

### Available Commands:
  `a`           Create a dns record of type A
  
  `cname`       Create a dns record of type CNAME

### Global Flags:
  `-c, --command-config [string]`   Provide a config file (default "config.toml")
&nbsp;  
## Delete a dns record by its name

### Usage:
  `constellix-cli delete [flags]`
  
  `constellix-cli delete [command]`

### Available Commands:
  `a`           Delete a dns record of type A
  
  `cname`       Delete a dns record of type CNAME

### Global Flags:
  `-c, --command-config [string]`   Provide a config file (default "config.toml")


&nbsp;  
## Get records or fetch all domains  

### Usage:
  `constellix-cli get [flags]`
  
  `constellix-cli get [command]`

### Available Commands:
  `a-id`        Get an A record id from a domain
  
  `all`         Get all records in a domain
  
  `cname-id`    Get a CNAME record id from a domain
  
  `domain`      Get all domains

### Global Flags:
  `-c, --command-config [string]`   Provide a config file (default "config.toml")
 

&nbsp;  
## TOML Config file
```
[constellix]
apiKey = "constellix-api-key"
secretKey = "constellix-secret-key"
domain = "domain-id"
```

