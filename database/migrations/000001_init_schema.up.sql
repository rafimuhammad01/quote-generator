CREATE TABLE IF NOT EXISTS quote(
    "id" serial primary key,
    "number_of_people" int not null,
    "sentences" text not null
)