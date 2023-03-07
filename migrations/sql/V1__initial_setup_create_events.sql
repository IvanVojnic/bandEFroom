create table if not exists rooms
(
    id            uuid PRIMARY KEY,
    idUserCreator uuid         NOT NULL,
    place         VARCHAR(255) NOT NULL,
    date          TIMESTAMP    NOT NULL
);

create table if not exists invites
(
    id      uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    room_id uuid NOT NULL,
    status  INT  NOT NULL,
    FOREIGN KEY (room_id) REFERENCES rooms (id)
        ON DELETE CASCADE
);