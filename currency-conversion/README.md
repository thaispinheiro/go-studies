# Currency Conversion CLI

A command-line application developed in Go for converting amounts from Brazilian Real (BRL) to other currencies using predefined exchange rates.

---

## Project Overview

The project contains a CLI tool capable of receiving an amount in BRL and a target currency code as arguments, returning the converted value based on fixed exchange rates.

---

## Features

- Currency conversion from BRL to multiple currencies
- Command-line execution
- Fixed exchange rates configuration
- Formatted output with two decimal places
- Input validation and error handling
- Support for case-insensitive currency codes

---

## Supported Currencies

| Currency | Code |
|---|---|
| US Dollar | USD |
| Euro | EUR |
| Japanese Yen | JPY |
| British Pound | GBP |
| Swiss Franc | CHF |
| Australian Dollar | AUD |

---

## Exchange Rates

```json
{
  "base": "BRL",
  "date": "2025-04-14",
  "rates": {
    "USD": 0.151,
    "EUR": 0.137,
    "JPY": 16.29,
    "GBP": 0.13,
    "CHF": 0.1402,
    "AUD": 0.2712
  }
}
```

---

## Command Syntax

```bash
./convert [amount_in_brl] [target_currency]
```

---

## Examples

### Convert 10 BRL to USD

```bash
$ ./convert 10 USD
1.51
```

### Convert 12 BRL to JPY

```bash
$ ./convert 12 JPY
195.48
```

---

## Project Structure

```bash
currency-conversion/
├── go.mod
├── main.go
└── rates.json
```

---

## Concepts Practiced

- Command-line arguments with `os.Args`
- Data structures using maps
- JSON file processing
- Type conversion and parsing
- String manipulation
- Conditional validation
- Error handling
- Output formatting

---

## Error Handling

The application validates:

- Incorrect number of arguments
- Invalid numeric values
- Unsupported currency codes
- Case-insensitive currency inputs

---

## How to Run

### Prerequisites

- Go installed on the machine

Installation guide:

https://golang.org/doc/install

### Running the Application

Navigate to the project folder and execute:

```bash
go run main.go [amount_in_brl] [target_currency]
```

Example:

```bash
go run main.go 100 USD
```
