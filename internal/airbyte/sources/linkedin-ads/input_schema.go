package linkedin_ads

const LINKEDIN_ADS_SECRETS_INPUT_JSON_SCHEMA = `{
    "type": "object",
    "title": "Google Ads Credential Spec",
    "$schema": "http://json-schema.org/draft-07/schema#",
    "required": [],
    "properties": {}
  }`

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
		  "format": "date",
          "default": "2021-06-01"
		}
	  }
	}`

const LINKEDIN_ADS_AIRBYTE_CONNECTION_SCHEMA = `{
  "documentationUrl":"https://docs.airbyte.io/integrations/sources/linkedin-ads",
  "connectionSpecification":{
    "$schema":"http://json-schema.org/draft-07/schema#",
    "title":"Linkedin Ads Spec",
    "type":"object",
    "required":[ "start_date" ],
    "additionalProperties":true,
    "properties":{
      "start_date":{
        "type":"string",
        "title":"Start Date",
        "pattern":"^[0-9]{4}-[0-9]{2}-[0-9]{2}$",
        "description":"UTC date in the format 2020-09-17. Any data before this date will not be replicated.",
        "examples":[ "2021-05-17" ]
      },
      "account_ids":{
        "title":"Account IDs",
        "type":"array",
        "description":"Specify the Account IDs separated by space, to pull the data from. Leave empty, if you want to pull the data from all associated accounts.",
        "items":{
          "type":"integer"
        },
        "default":[]
      },
      "credentials":{
        "title":"Authorization Method",
        "type":"object",
        "properties":{
          "type":"object",
          "title":"OAuth2.0",
          "required":[
            "client_id",
            "client_secret",
            "refresh_token"
          ],
          "properties":{
            "auth_method":{
              "type":"string",
              "const":"oAuth2.0"
            },
            "client_id":{
              "type":"string",
              "title":"Client ID",
              "description":"The Client ID of the LinkedIn Ads developer application."
            },
            "client_secret":{
              "type":"string",
              "title":"Client Secret",
              "description":"The Client Secret the LinkedIn Ads developer application."
            },
            "refresh_token":{
              "type":"string",
              "title":"Refresh Token",
              "description":"The key to refresh the expired Access Token."
            }
          }
        }
      }
    }
  }
}`
