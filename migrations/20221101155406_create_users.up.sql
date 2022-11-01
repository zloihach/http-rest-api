CREATE TABLE users
(
    id    bigserial PRIMARY KEY NOT NULL,
    email varchar(255) not null unique,
    encrypted_password varchar(255) not null
)