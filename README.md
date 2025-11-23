# idconv - uuid conversion tool

idconv allows to convert string uuid to base64, rawBase64, or vice versa. Also it allows to generate random uuid (and output it in all 3 formats).

### Usage

```
idconv [uuid | base64 | rawBase64] | random
```

### Examples
First option - convert basic string format of uuid: 

```
idconv 9ca5de1c-c01a-4cc1-96d6-1017b40d5843
```
Second option - convert base64 string to uuid (if it's valid):

```
idconv nKXeHMAaTMGW1hAXtA1YQw==
```

Next option - convert rawBase64 (no padding) string to uuid (if it's valid):

```
idconv nKXeHMAaTMGW1hAXtA1YQw
```

And finally - just generate random uuid:

```
idconv random
```

In all cases there will be the following output:

```
 ----------------------------------------------------
|   uuid    |  9ca5de1c-c01a-4cc1-96d6-1017b40d5843  |
 ----------------------------------------------------
|  base64   |        nKXeHMAaTMGW1hAXtA1YQw==        |
 ----------------------------------------------------
| rawBase64 |         nKXeHMAaTMGW1hAXtA1YQw         |
 ----------------------------------------------------
```

