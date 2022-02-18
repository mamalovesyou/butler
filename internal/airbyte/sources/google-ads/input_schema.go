package google_ads

const GOOGLE_ADS_SECRETS_INPUT_JSON_SCHEMA = `{
    "type": "object",
    "title": "Google Ads Spec",
    "$schema": "http://json-schema.org/draft-07/schema#",
    "properties": {
      "developer_token": {
		"type": "string",
		"title": "Developer Token",
		"description": "Developer token granted by Google to use their APIs."
	  }
    }
  }`

const GOOGLE_ADS_CONFIG_INPUT_JSON_SCHEMA = `{
    "type": "object",
    "title": "Google Ads Spec",
    "$schema": "http://json-schema.org/draft-07/schema#",
    "required": [
      "start_date",
      "customer_id"
    ],
    "properties": {
      "start_date": {
        "type": "string",
        "order": 2,
        "title": "Start Date",
        "pattern": "^[0-9]{4}-[0-9]{2}-[0-9]{2}$",
		"default": "2021-06-01",
        "examples": [
          "2017-01-25"
        ],
        "description": "UTC date and time in the format 2017-01-25. Any data before this date will not be replicated."
      },
      "customer_id": {
        "type": "string",
        "order": 1,
        "title": "Customer ID",
		"pattern": "^[0-9]{10}$",
        "description": "Customer ID must be specified as a 10-digit number without dashes."
      },
      "conversion_window_days": {
        "type": "integer",
        "order": 5,
        "title": "Conversion Window (Optional)",
        "default": 14,
        "maximum": 1095,
        "minimum": 0,
        "description": "A conversion window is the period of time after an ad interaction (such as an ad click or video view) during which a conversion, such as a purchase, is recorded in Google Ads."
      }
    },
    "additionalProperties": true
  }`

const GOOGLE_ADS_AIRBYTE_CONNECTION_SCHEMA = `{
  "documentationUrl":"https://docs.airbyte.com/integrations/sources/google-ads",
  "connectionSpecification":{
    "$schema":"http://json-schema.org/draft-07/schema#",
    "title":"Google Ads Spec",
    "type":"object",
    "required":[
      "credentials",
      "start_date",
      "customer_id"
    ],
    "additionalProperties":true,
    "properties":{
      "credentials":{
        "type":"object",
        "title":"Google Credentials",
        "order":0,
        "required":[
          "developer_token",
          "client_id",
          "client_secret",
          "refresh_token"
        ],
        "properties":{
          "developer_token":{
            "type":"string",
            "title":"Developer Token",
            "description":"Developer token granted by Google to use their APIs. More instruction on how to find this value in our <a href=\"https://docs.airbyte.com/integrations/sources/google-ads#setup-guide\">docs</a>",
            "airbyte_secret":true
          },
          "client_id":{
            "type":"string",
            "title":"Client ID",
            "description":"The Client ID of your Google Ads developer application. More instruction on how to find this value in our <a href=\"https://docs.airbyte.com/integrations/sources/google-ads#setup-guide\">docs</a>"
          },
          "client_secret":{
            "type":"string",
            "title":"Client Secret",
            "description":"The Client Secret of your Google Ads developer application. More instruction on how to find this value in our <a href=\"https://docs.airbyte.com/integrations/sources/google-ads#setup-guide\">docs</a>",
            "airbyte_secret":true
          },
          "access_token":{
            "type":"string",
            "title":"Access Token",
            "description":"Access Token for making authenticated requests. More instruction on how to find this value in our <a href=\"https://docs.airbyte.com/integrations/sources/google-ads#setup-guide\">docs</a>",
            "airbyte_secret":true
          },
          "refresh_token":{
            "type":"string",
            "title":"Refresh Token",
            "description":"The token for obtaining a new access token. More instruction on how to find this value in our <a href=\"https://docs.airbyte.com/integrations/sources/google-ads#setup-guide\">docs</a>",
            "airbyte_secret":true
          }
        }
      },
      "customer_id":{
        "title":"Customer ID(s)",
        "type":"string",
        "description":"Comma separated list of (client) customer IDs. Each customer ID must be specified as a 10-digit number without dashes. More instruction on how to find this value in our <a href=\"https://docs.airbyte.com/integrations/sources/google-ads#setup-guide\">docs</a>. Metrics streams like AdGroupAdReport cannot be requested for a manager account.",
        "pattern":"^[0-9]{10}(,[0-9]{10})*$",
        "examples":[
          "6783948572,5839201945"
        ],
        "order":1
      },
      "start_date":{
        "type":"string",
        "title":"Start Date",
        "description":"UTC date and time in the format 2017-01-25. Any data before this date will not be replicated.",
        "pattern":"^[0-9]{4}-[0-9]{2}-[0-9]{2}$",
        "examples":[
          "2017-01-25"
        ],
        "order":2
      },
      "conversion_window_days":{
        "title":"Conversion Window (Optional)",
        "type":"integer",
        "description":"A conversion window is the period of time after an ad interaction (such as an ad click or video view) during which a conversion, such as a purchase, is recorded in Google Ads.",
        "minimum":0,
        "maximum":1095,
        "default":14,
        "examples":[
          14
        ],
        "order":5
      }
    }
  }
}`
