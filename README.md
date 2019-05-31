# Crumpet

CLI test suite written in Go

## Building a crumpet

Run ```make``` in the root of the directory, this will compile the application and drop the ```crumpet``` binary in the ```bin/``` folder.

## Running tests

Right now crumpet only performs ```HTTP GET``` requests. I'll add more methods as when I need them.

You can either run test from the CLI by providing arguments, or you can provide a json file by passing the file path to the ```-spec-file``` cli arg.

## Specfile schema

Specfile are .json files that adhere to the following schema:

```
{
  "host": "https://tempuri.org",
  "paths": [
    "/",
    "/abc",
    "/abc?xyz=123"
  ],
  "iterations": 2500,
  "concurrency": 10,
  "maxDelayMs": 2500,
  "options": {
    "httpRequestHeaders": {
      "Authorization": "bearer abc-123"
    }
  }
}
```
### What does it all mean?
| specfile property | purpose | allowed values |
|------------------|---------|----------------|
| host | identifies the hostname of your api | {HTTP/S}://{any string} |
| paths | the HTTP paths you want to test. These will be randomly chosen for testing | array of strings |
| iterations | the total number of tests to perform | any whole number > 0 |
| concurrency | the number of concurrent 'test runners' to use | any whole number > 0 |
| maxDelayMs | the maximum amount of time to delay between requests | any non-negative whole number |
| options | a dynamic collection of options | a JSON object - the content of it will vary but breaking changes will be gaurded against within the codebase |
