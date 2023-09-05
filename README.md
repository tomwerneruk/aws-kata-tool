# aws-kata-tool

`aws-kata-tool` is a 'stub' that can be used as a placeholder service for demos or gamedays, or as a test harness at different points of your project.

This tool is designed to grow over time with various AWS themed functions for different AWS architectures.

## Mimic Server

An HTTP Server that simply listens and responds back with what was sent in the headers. This is useful for stubbing servers behind reverse proxies, or for quick and dirty HTTP load tests.

`aws-kata-tool mimic` - Starts on a default non-encrypted port of 8080, on all interfaces.
`aws-kata-tool mimic -p 8888` - Starts on a specified non-encrypted port of 8888, on all interfaces.

The mimic service can be called using GET `/mimic`.

An optional query parameter can be passed to influence the output of the command. Pass

`/mimic?format=pretty` for an HTML table formatted output
`/mimic?format=json` for a machine readable JSON response