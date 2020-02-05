# go-recordlocator

**A Go package to encode integers into a short and easy to read string and decode the strings back to the initial numbers**

This is the Go equivalent to [php-recordlocator](https://github.com/jakoubek/php-recordlocator), [lua-recordlocator](https://github.com/jakoubek/lua-recordlocator), [postgresql-recordlocator](https://github.com/jakoubek/postgresql-recordlocator) and [sqlanywhere-recordlocator](https://github.com/jakoubek/sqlanywhere-recordlocator).

A RecordLocator is an alphanumerical string that represents an integer. This package encodes integers to RecordLocators and decodes RecordLocators back to integers.

A RecordLocator is shorter than the corresponding integer and easier to read and memorize. You can use it to encode autoincrement primary keys from an database into an human-readable representation for your users.

## Examples

- The integer 5,325 encodes to the RecordLocator 78G.
- The integer 3,512,345 encodes to the RecordLocator 5E82T.

Both RecordLocators are shorter than their integer equivalent. You can encode more than 33.5 million integers in an 5-char RecordLocator: the largest 5-char RecordLocator, ZZZZZ, represents the integer 33,554,431.

With 6 chars you can encode integers up to more than 1 billion.

## Installation

```
go get -u github.com/jakoubek/go-recordlocator
```

## Usage

Encoding integers to recordlocators:

```go
recloc := recordlocator.NewRecordLocator()
str, _ := recloc.Encode(5325)
fmt.Printf("The number %d encodes to %s.\n", 5325, str)  // The number 5325 encodes to 78G.
```

Decoding recordlocators back into the integer values:

```go
recloc := recordlocator.NewRecordLocator()
nr, _ := recloc.Decode("78G")
fmt.Printf("The recordlocator %s decodes back to %d.\n", "78G", nr) // The recordlocator 78G decodes back to 5325.
```