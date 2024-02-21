CREATE TABLE "user"(
    "id" SERIAL NOT NULL PRIMARY KEY,
    "full_name" VARCHAR NOT NULL,
    "nick_name" VARCHAR NOT NULL,
    "photo" VARCHAR NOT NULL,
    "birthday" VARCHAR NOT NULL,
    "location" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);