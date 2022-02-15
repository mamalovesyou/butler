package linkedin_ads

const LINKEDIN_ADS_CONFIG_INPUT_JSON_SCHEMA = `{
	  "$id": "https://example.com/person.schema.json",
	  "$schema": "https://json-schema.org/draft/2020-12/schema",
	  "title": "LinkedinAdsSource",
	  "type": "object",
      "required": [ "account_id", "start_date"],
	  "properties": {
		"account_ids": {
		  "type": "array",
		  "description": "You must specify an account ID.",
          "items": { "type": "number" }
		},
		"start_date": {
		  "description": "UTC date and time in the format 2017-01-25. Any data before this date will not be replicated.",
		  "type": "string",
		  "format": "date"	
		}
	  }
	}`
