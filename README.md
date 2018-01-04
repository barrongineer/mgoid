# mgoid

A simple cli to generate and validate Mongo ObjectID's.

Install/Update:

`go get -u github.com/barrongineer/mgoid`

Usage:
```
github.com/barrongineer/mgoid                                                                                                                                                                                                                         
▶ mgoid new --copy
5a4e5d29768a4174ea47d06f
copied to clipboard

github.com/barrongineer/mgoid                                                                                                                                                                                                                         
▶ mgoid validate 5a4e5d29768a4174ea47d06f
valid

github.com/barrongineer/mgoid                                                                                                                                                                                                                         
▶ mgoid validate something
invalid

```