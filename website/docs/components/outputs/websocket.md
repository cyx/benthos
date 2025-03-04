---
title: websocket
type: output
status: stable
categories: ["Network"]
---

<!--
     THIS FILE IS AUTOGENERATED!

     To make changes please edit the contents of:
     lib/output/websocket.go
-->

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

Sends messages to an HTTP server via a websocket connection.


<Tabs defaultValue="common" values={[
  { label: 'Common', value: 'common', },
  { label: 'Advanced', value: 'advanced', },
]}>

<TabItem value="common">

```yml
# Common config fields, showing default values
output:
  label: ""
  websocket:
    url: ""
```

</TabItem>
<TabItem value="advanced">

```yml
# All config fields, showing default values
output:
  label: ""
  websocket:
    url: ""
    tls:
      enabled: false
      skip_cert_verify: false
      enable_renegotiation: false
      root_cas: ""
      root_cas_file: ""
      client_certs: []
    oauth:
      enabled: false
      consumer_key: ""
      consumer_secret: ""
      access_token: ""
      access_token_secret: ""
    basic_auth:
      enabled: false
      username: ""
      password: ""
    jwt:
      enabled: false
      private_key_file: ""
      signing_method: ""
      claims: {}
      headers: {}
```

</TabItem>
</Tabs>

## Fields

### `url`

The URL to connect to.


Type: `string`  
Default: `""`  

### `tls`

Custom TLS settings can be used to override system defaults.


Type: `object`  

### `tls.enabled`

Whether custom TLS settings are enabled.


Type: `bool`  
Default: `false`  

### `tls.skip_cert_verify`

Whether to skip server side certificate verification.


Type: `bool`  
Default: `false`  

### `tls.enable_renegotiation`

Whether to allow the remote server to repeatedly request renegotiation. Enable this option if you're seeing the error message `local error: tls: no renegotiation`.


Type: `bool`  
Default: `false`  
Requires version 3.45.0 or newer  

### `tls.root_cas`

An optional root certificate authority to use. This is a string, representing a certificate chain from the parent trusted root certificate, to possible intermediate signing certificates, to the host certificate.


Type: `string`  
Default: `""`  

```yml
# Examples

root_cas: |-
  -----BEGIN CERTIFICATE-----
  ...
  -----END CERTIFICATE-----
```

### `tls.root_cas_file`

An optional path of a root certificate authority file to use. This is a file, often with a .pem extension, containing a certificate chain from the parent trusted root certificate, to possible intermediate signing certificates, to the host certificate.


Type: `string`  
Default: `""`  

```yml
# Examples

root_cas_file: ./root_cas.pem
```

### `tls.client_certs`

A list of client certificates to use. For each certificate either the fields `cert` and `key`, or `cert_file` and `key_file` should be specified, but not both.


Type: `array`  
Default: `[]`  

```yml
# Examples

client_certs:
  - cert: foo
    key: bar

client_certs:
  - cert_file: ./example.pem
    key_file: ./example.key
```

### `tls.client_certs[].cert`

A plain text certificate to use.


Type: `string`  
Default: `""`  

### `tls.client_certs[].key`

A plain text certificate key to use.


Type: `string`  
Default: `""`  

### `tls.client_certs[].cert_file`

The path to a certificate to use.


Type: `string`  
Default: `""`  

### `tls.client_certs[].key_file`

The path of a certificate key to use.


Type: `string`  
Default: `""`  

### `oauth`

Allows you to specify open authentication via OAuth version 1.


Type: `object`  

### `oauth.enabled`

Whether to use OAuth version 1 in requests.


Type: `bool`  
Default: `false`  

### `oauth.consumer_key`

A value used to identify the client to the service provider.


Type: `string`  
Default: `""`  

### `oauth.consumer_secret`

A secret used to establish ownership of the consumer key.


Type: `string`  
Default: `""`  

### `oauth.access_token`

A value used to gain access to the protected resources on behalf of the user.


Type: `string`  
Default: `""`  

### `oauth.access_token_secret`

A secret provided in order to establish ownership of a given access token.


Type: `string`  
Default: `""`  

### `basic_auth`

Allows you to specify basic authentication.


Type: `object`  

### `basic_auth.enabled`

Whether to use basic authentication in requests.


Type: `bool`  
Default: `false`  

### `basic_auth.username`

A username to authenticate as.


Type: `string`  
Default: `""`  

### `basic_auth.password`

A password to authenticate with.


Type: `string`  
Default: `""`  

### `jwt`

BETA: Allows you to specify JWT authentication.


Type: `object`  

### `jwt.enabled`

Whether to use JWT authentication in requests.


Type: `bool`  
Default: `false`  

### `jwt.private_key_file`

A file with the PEM encoded via PKCS1 or PKCS8 as private key.


Type: `string`  
Default: `""`  

### `jwt.signing_method`

A method used to sign the token such as RS256, RS384 or RS512.


Type: `string`  
Default: `""`  

### `jwt.claims`

A value used to identify the claims that issued the JWT.


Type: `object`  
Default: `{}`  

### `jwt.headers`

Add optional key/value headers to the JWT.


Type: `object`  
Default: `{}`  


