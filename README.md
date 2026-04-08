# Go EU VAT Validation Client

![EuroValidate](https://img.shields.io/badge/Powered%20by-EuroValidate-6d5cff)

Go client library for EuroValidate API — validate EU VAT, IBAN, EORI

## Quick Start

1. Get a free API key at [eurovalidate.com](https://eurovalidate.com)
2. Clone this repo: `git clone https://github.com/eurovalidate/go-vat-client.git`
3. Copy `.env.example` to `.env` and add your API key
4. Install dependencies and run (see below)

## Setup

```bash
go run main.go
```

## What This Does

Go client library for EuroValidate API — validate EU VAT, IBAN, EORI. Uses the [EuroValidate API](https://api.eurovalidate.com/docs) for real-time validation against official EU government sources (VIES, HMRC, GLEIF).

## API Response Example

```json
{
  "vat_number": "NL820646660B01",
  "country_code": "NL",
  "status": "valid",
  "company_name": "ABN AMRO BANK N.V.",
  "meta": {
    "confidence": "high",
    "source": "vies_cached",
    "cached": true,
    "response_time_ms": 2
  }
}
```

## Test VAT Numbers

| VAT | Country | Expected |
|-----|---------|----------|
| NL820646660B01 | Netherlands | Valid |
| FR40303265045 | France | Valid |
| IT00159560366 | Italy | Valid |
| DE000000000 | Germany | Invalid |

## Links

- [EuroValidate API Docs](https://api.eurovalidate.com/docs)
- [Get Free API Key](https://eurovalidate.com) (no credit card)
- [Python SDK](https://pypi.org/project/eurovalidate/) (`pip install eurovalidate`)
- [Node.js SDK](https://www.npmjs.com/package/@eurovalidate/sdk) (`npm install @eurovalidate/sdk`)
- [PHP SDK](https://packagist.org/packages/eurovalidate/sdk) (`composer require eurovalidate/sdk`)

## License

MIT
