DROP TABLE "users";
DROP TABLE "sessions";
DROP TABLE "subjects";
DROP TABLE "ratings";
DROP TABLE "addresses";

CREATE TABLE "users" (
    "id" INTEGER,
    "uuid" TEXT NOT NULL UNIQUE,
    "first_name" TEXT,
    "last_name" TEXT,
    "email" TEXT NOT NULL UNIQUE,
    "password" TEXT NOT NULL,
    "created_at" DATETIME NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "sessions"(
    "id" INTEGER,
    "uuid" TEXT NOT NULL UNIQUE,
    "email" TEXT NOT NULL UNIQUE,
    "created_at" DATETIME NOT NULL,
    "user_id" INTEGER,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES "users"("id")
);

CREATE TABLE "subjects"(
    "id" INTEGER,
    "title" TEXT NOT NULL UNIQUE,
    PRIMARY KEY ("id")
);

CREATE TABLE "ratings"(
    "rating" INTEGER CHECK("rating">=0),
    "rater_id" INTEGER,
    "rated_id" INTEGER,
    PRIMARY KEY ("rater_id", "rated_id"),
    FOREIGN KEY ("rater_id") REFERENCES "users"("id"),
    FOREIGN KEY ("rated_id") REFERENCES "users"("id")
);

CREATE TABLE "addresses"(
    "id" INTEGER,
    "line_1" TEXT,
    "line_2" TEXT,
    "landmark" TEXT,
    "city" TEXT,
    "state" TEXT,
    "pin" INTEGER,
    "country" TEXT,
    "latitude" REAL,
    "longitude" REAL,
    "user_id" INTEGER,
    PRIMARY KEY ("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id")
);