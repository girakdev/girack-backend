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

  - _Optional_
  - _datetime_

- `oldest_datetime`

  - _Optional_
  - _datetime_

#### Response

- `ok`

  - _boolean_

- `messages`

  - `id`

    - _int_

  - `user_id`

    - _int_

  - `timestamp`

    - _datetime_

  - `text`

    - _text_

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
