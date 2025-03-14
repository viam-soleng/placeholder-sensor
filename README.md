# Module placeholder

This module provides a simple placeholder sensor that returns configured data without performing any actual sensing. It's useful for testing, development, or situations where you need to return static or pre-defined data.

## Model bill:placeholder:sensor

A configurable sensor that returns any data structure specified in its configuration. This sensor acts as a data pass-through, returning whatever data is provided in its configuration when readings are requested.

### Configuration
The following attribute template can be used to configure this model:

```json
{
  "readings": {
    "key1": <value1>,
    "key2": <value2>,
    ...
  }
}
```

#### Attributes

The following attributes are available for this model:

| Name       | Type                | Inclusion | Description                                     |
|------------|---------------------|-----------|------------------------------------------------|
| `readings` | map[string]interface{} | Required  | A map of key-value pairs to return as readings |

#### Example Configuration

```json
{
  "readings": {
    "location_details": ["Main Concourse"],
    "concession_options": ["Beer", "Wine", "Seltzers", "Cocktails"],
    "estimated_wait_time_min": "NONE",
    "location_name": "Section 123: Old Joe's Bar",
    "count_in_view": 0,
    "location_open": true
  }
}
```

### Readings

The sensor returns exactly what is specified in the `readings` configuration field. This allows for flexible data structures to be returned without modifying the sensor code.

Response:

```json
  {
    "location_details": ["Main Concourse"],
    "concession_options": ["Beer", "Wine", "Seltzers", "Cocktails"],
    "estimated_wait_time_min": "NONE",
    "location_name": "Section 123: Old Joe's Bar",
    "count_in_view": 0,
    "location_open": true
  }
```