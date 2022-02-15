package google_ads

const GOOGLE_ADS_CONFIG_INPUT_JSON_SCHEMA = `{
	  "$id": "https://example.com/person.schema.json",
	  "$schema": "https://json-schema.org/draft/2020-12/schema",
	  "title": "GoogleAdwordsSource",
	  "type": "object",
      "required": [ "developer_token", "customer_id", "start_date"],
	  "properties": {
		"developer_token": {
		  "type": "string",
		  "description": "Developer token granted by Google to use their APIs."
		},
		"customer_id": {
		  "type": "string",
		  "description": "Customer ID must be specified as a 10-digit number without dashes"
		},
		"start_date": {
		  "description": "UTC date and time in the format 2017-01-25. Any data before this date will not be replicated.",
		  "type": "string",
          "format": "date"
		}
	  }
	}`
