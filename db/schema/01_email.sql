
CREATE TABLE service (
    service_id   SERIAL PRIMARY KEY NOT NULL,
    service_name TEXT   NOT NULL
);

CREATE TABLE email (
    email_id      SERIAL  PRIMARY KEY NOT NULL,
    email_to      TEXT    NOT NULL,
    email_from    TEXT    NOT NULL,
    email_subject TEXT    NOT NULL,
    email_content JSONB   NOT NULL,
    service_id    INTEGER NOT NULL REFERENCES service(service_id)
);

CREATE TABLE event_type (
    event_type_id          INTEGER PRIMARY KEY NOT NULL,
    event_type_name        TEXT    NOT NULL,
    event_type_description TEXT    NOT NULL
);

INSERT INTO event_type
    (event_type_id, event_type_name, event_type_description)
VALUES
    (10, 'rendered',     'The email was rendered'),
    (20, 'not_accepted', 'The email was not accepted by the service provider'),
    (30, 'accepted',     'The email was accepted by the service provider'),
    (40, 'temporary',    'Delivery of this email temporarily failed'),
    (50, 'delivered',    'The email was delivered'),
    (60, 'opened',       'The email was opened by the recipient'),
    (70, 'clicked',      'A link was clicked by the recipient'),
    (80, 'permanent',    'Delivery of this email permanently failed'),
    (90, 'unsubscribed', 'The recipient unsubscribed from the email');

CREATE TABLE event (
    event_id      SERIAL    PRIMARY KEY NOT NULL,
    email_id      INTEGER   NOT NULL REFERENCES email(email_id),
    event_type_id INTEGER   NOT NULL REFERENCES event_type(event_type_id) DEFAULT 10,
    event_content TEXT      NOT NULL,
    event_created TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);
