# aws-kata-tool

![alt text](ninja-gopher.png)

`aws-kata-tool` is a 'stub' that can be used as a placeholder service for demos or gamedays, or as a test harness at different points of your project.

This tool is designed to grow over time with various AWS themed functions for different AWS architectures.

## Start the restapi server

The Mimic Server and Simulated API Endpoint are hosted by the `restapi` sub-command of the tool. 

`aws-kata-tool restapi` - Starts on a default non-encrypted port of 8080, on all interfaces.
`aws-kata-tool restapi -p 8888` - Starts on a specified non-encrypted port of 8888, on all interfaces.

## Mimic Endpoint

An HTTP response that simply replays back with what was sent in the headers. This is useful for stubbing servers behind reverse proxies, or for quick and dirty HTTP load tests.

The mimic service can be called using GET `/mimic`.

An _optional_ query parameter can be passed to influence the output of the command. Pass

`/mimic?format=pretty` for an HTML table formatted output
`/mimic?format=json` for a machine readable JSON response
`/mimic?format=prettyJson` for a human readable JSON response (with extra \n that may break programmatic parsing)

## Simulated API Endpoint

You may want to execute a call that actually creates CPU load, simulating an actual application. This allows for a realistic experience in a kata, playing with metrics to scale ASGs or see how Cloudwatch Metrics work.

The simulated API can be reached by issuing a GET `/restapi/messages`. Response will be delayed, as the function synchronously blocks while calculating a prime between 1000 and 1500 digits long. As this causes CPU load, the delay will scale with the number of simulataneous connections and instance size.