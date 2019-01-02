# Getini

This is a simple command line utility that extracts a value from an INI file.

It is designed to be simple, stupid and have no dependencies outside the Go 
standard library.

## Usage:

```
getini [FILE] [SECTION] KEY

```

This will output the value of the KEY in section SECTION in the INI file given 
by FILE.

If FILE is not given or is "-", standard input is used.

If SECTION is not given, it defaults to returning matching keys before the 
first section header in the INI file.

KEY and SECTION are case-insensitive.

If the key was found, it is output to standard output.
If the key was not found, `getini` exits with code 1.

