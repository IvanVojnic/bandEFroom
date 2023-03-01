create table if not exists rooms(
    id uuid primary key not null,
    idUserCreator uuid not null,
    place varchar(255) not null,
    date timestamp not null
);

create table if not exists invites(
    id uuid primary key not null,
    user_id uuid not null,
    room_id uuid not null,
    status int not null,
    FOREIGN KEY (room_id) REFERENCES rooms(id)
      ON DELETE CASCADE
);