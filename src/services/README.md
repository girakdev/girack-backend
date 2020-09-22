# Girack API

## Message

### GET /v1/message

#### Parameters

- `channel`

  - **Required**
  - _String_

- `count`

  - _Optional_
  - _int_

- `latest_datetime`

  - **Optional**
  - _datetime_

- `oldest_datetime`

  - **Optional**
  - _datetime_

#### Response

- `ok`

  - _boolean_

- `messages`

  - `Message`
  - _list_

### POST /v1/message

#### Parameters

- `channel`

  - **Required**
  - _String_

- `text`

  - **Required**
  - _int_

#### Response

- `ok`

  - _boolean_
