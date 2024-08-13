-- -------------------------------------------------------------
-- TablePlus 5.1.0(468)
--
-- https://tableplus.com/
--
-- Database: sahakolay
-- Generation Time: 2024-08-13 20:06:45.4950
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS cities_id_seq;

-- Table Definition
CREATE TABLE "public"."cities" (
    "id" int2 NOT NULL DEFAULT nextval('cities_id_seq'::regclass),
    "country_id" int2 NOT NULL,
    "name" varchar NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS counties_id_seq;

-- Table Definition
CREATE TABLE "public"."counties" (
    "id" int4 NOT NULL DEFAULT nextval('counties_id_seq'::regclass),
    "city_id" int2 NOT NULL,
    "name" varchar NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS countries_id_seq;

-- Table Definition
CREATE TABLE "public"."countries" (
    "id" int4 NOT NULL DEFAULT nextval('countries_id_seq'::regclass),
    "name" varchar NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS towns_id_seq;

-- Table Definition
CREATE TABLE "public"."towns" (
    "id" int4 NOT NULL DEFAULT nextval('towns_id_seq'::regclass),
    "county_id" int4 NOT NULL,
    "name" varchar NOT NULL,
    "latitude" float8,
    "longitude" float8,
    "postal_code" varchar,
    PRIMARY KEY ("id")
);

ALTER TABLE "public"."cities" ADD FOREIGN KEY ("country_id") REFERENCES "public"."countries"("id");
ALTER TABLE "public"."counties" ADD FOREIGN KEY ("city_id") REFERENCES "public"."cities"("id");
ALTER TABLE "public"."towns" ADD FOREIGN KEY ("county_id") REFERENCES "public"."counties"("id");
