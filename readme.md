# Ticketer

An event/show ticket management system
Users sign up and buy ticket for events
Organisations can sign up and create events
Register users can buy ticket for themselves and others that also have an acount

## Database

mongo

## Api endpoints

#### User endpoint

create user `````{domain}/user/`````

get user `````{domain}/user/id`````

put user `````{domain}/user/id`````

put bookmark `````{domain}/user/bookmark/id/event_id`````

delete user `````{domain}/user/id`````

expected json when creating account:

```json
{
  "email": "randomEmail",
  "password": "randomPassword",
  "username": "randomFullName"
}
```

full parameters:

```json
{
  "ID": "666ba20485ea766421b1d151",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "email": "Susanna6@hotmail.com",
  "password": "WVHfRdvoqnmB5Xy",
  "username": "Vernon Franecki",
  "profile_image_url": "http://fgretrrtrtttttttttttttttttttttt.com/640/480",
  "cover_image_url": "",
  "tickets": "null array",
  "events_interested_in": "null array"
}
```

#### Organisation endpoint

create organisation `````{domain}/org/`````

get organisation `````{domain}/org/id`````

put organisation `````{domain}/org/id`````

delete organisation `````{domain}/org/id`````

expected json when creating organisation:

```json
{
  "organisation_name": "randomCompanyName",
  "organisation_email": "randomEmail",
  "organisation_password": "randomPassword",
  "organisation_address": "randomStreetAddress",
  "organisation_description": "randomLoremSentences"
}
```

full parameters:

```json
{
  "ID": "66760e6eebed1ebfe25f615d",
  "organisation_name": "Ullrich, Wisoky and Morar",
  "organisation_email": "Hazel_Wuckert12@gmail.com",
  "organisation_password": "2r7_qOWq9IrWXb7",
  "organisation_address": "45322 Moen Estate",
  "organisation_profile_image_url": "",
  "organisation_cover_image_url": "",
  "organisation_description": "Maxime vel labore qui illo. Nesciunt quasi impedit quas. Eum consequuntur dolorem quasi ut voluptatem porro omnis sit rem.",
  "events": "null array",
  "CreatedAt": "0001-01-01T00:00:00Z"
}
```

#### Event endpoint

create event `````{domain}/event/`````

get event `````{domain}/event/id`````

put event `````{domain}/event/id`````

search event `````{domain}/event?keyword=yourkeyword&limit=30&page=1`````

buy ticket to event `````{domain}/event?event_id=yourevent_id&user_id=example_user_id&buy_for_id=exmaple_user_id`````

delete event `````{domain}/event/id`````

expected json when creating organisation:

```json
{
  "number_of_ticket_printed": 5,
  "number_of_ticket_sold": 0,
  "number_of_ticket_available": 5,
  "even_name": "randomProductName",
  "event_description": "randomLoremSentences",
  "event_address": "randomStreetAddress",
  "event_date": "2024-06-01 22:04:14.159383 +00:00",
  "ticket_start_sales_date": "2024-06-01 22:04:14.159383 +00:00",
  "ticket_end_sales_date": "2024-06-01 22:04:14.159383 +00:00",
  "organisation_id": "666d08c38c32fdd09bec890d"
}
```

full parameters:

```json
{
  "ID": "667a3f17b1448b15964b0538",
  "number_of_ticket_printed": 0,
  "number_of_ticket_sold": 0,
  "number_of_ticket_available": 0,
  "even_name": "Fantastic Fresh Chicken",
  "event_description": "",
  "event_cover_image": "",
  "event_address": "",
  "event_date": "",
  "ticket_start_sales_date": "",
  "ticket_end_sales_date": "",
  "ticket_ids": "null array",
  "bought_ticket_ids": "null array",
  "organisation_id": ""
}
```

#### Ticket endpoint

get ticket `````{domain}/ticket/id`````
full parameters:
```json
{
    "ID": "66776b0f9ff92e6d13c5232f",
    "event_id": "66776b0f9ff92e6d13c5232d",
    "ticket_id": "9u$zRBM$Xvj",
    "buy_date": "",
    "bought_by": "",
    "bought_for": ""
}
```


