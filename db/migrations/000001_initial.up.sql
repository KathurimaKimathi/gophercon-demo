BEGIN;

CREATE TABLE "users_user" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "username" varchar(20) UNIQUE NOT NULL,
  "first_name" varchar(25) NOT NULL,
  "last_name" varchar(25) NOT NULL,
  "email" varchar(100) UNIQUE NOT NULL,
  "user_type" varchar(20) NOT NULL
);

CREATE TABLE "user_credentials" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "valid_from" timestamp NOT NULL,
  "valid_to" timestamp NOT NULL,
  "hashed_pin" text NOT NULL,
  "salt" text NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "business_profile" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp,
  "created_by" uuid,
  "updated_by" uuid,
  "name" varchar(255) UNIQUE NOT NULL,
  "logo" text UNIQUE NOT NULL,
  "active" boolean NOT NULL,
  "description" text NOT NULL,
  "intro_statement" text NOT NULL,
  "mission" text NOT NULL,
  "vision" text NOT NULL,
  "slogan" varchar(255),
  "phone_number" varchar(20) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "facebook" varchar(255) UNIQUE,
  "x" varchar(255) UNIQUE,
  "instagram" varchar(255) UNIQUE,
  "tiktok" varchar(255) UNIQUE,
  "whats_app" varchar(255) UNIQUE,
  "linkedin" varchar(255) UNIQUE,
  "postal_address" text UNIQUE NOT NULL,
  "city" varchar(100) NOT NULL,
  "country" varchar(100) NOT NULL,
  "building" varchar(100) NOT NULL,
  "floor_number" varchar(10)
);

CREATE TABLE "partner" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "heading" varchar(255) NOT NULL,
  "heading_support_text" text,
  "business_profile" uuid UNIQUE NOT NULL
);

CREATE TABLE "business_partner" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "name" varchar(255) UNIQUE NOT NULL,
  "logo" text NOT NULL,
  "partner_id" uuid NOT NULL
);

CREATE TABLE "base_service" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "heading" varchar(255) NOT NULL,
  "heading_support_text" text,
  "business_profile" uuid NOT NULL
);

CREATE TABLE "service" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "title" varchar(255) NOT NULL,
  "description" text NOT NULL,
  "base_service" uuid NOT NULL
);

CREATE TABLE "commitment" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "heading" varchar(255) NOT NULL,
  "heading_support_text" text,
  "business_profile" uuid NOT NULL
);

CREATE TABLE "why_us" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "title" varchar(255) NOT NULL,
  "description" text,
  "commitment" uuid NOT NULL
);

CREATE TABLE "story" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "heading" varchar(255) NOT NULL,
  "heading_support_text" text,
  "business_profile" uuid NOT NULL
);

CREATE TABLE "content" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "title" varchar(100) NOT NULL,
  "description" varchar(200) NOT NULL,
  "body" text NOT NULL,
  "hero_image" text NOT NULL,
  "gallery_images" text[],
  "category" varchar(50) NOT NULL,
  "tags" text[],
  "content_type" varchar(50) NOT NULL,
  "media" text,
  "price" float DEFAULT 0.00,
  "story_id" uuid NOT NULL
);

ALTER TABLE
    IF EXISTS "user_credentials"
    ADD CONSTRAINT "user_credentials_created_by_fkey" FOREIGN KEY ("created_by") REFERENCES "users_user" ("id");

ALTER TABLE
    IF EXISTS "content"
    ADD CONSTRAINT "content_created_by_fkey" FOREIGN KEY ("created_by") REFERENCES "users_user" ("id");

ALTER TABLE
    IF EXISTS "story"
    ADD CONSTRAINT "story_business_profile_fkey" FOREIGN KEY ("business_profile") REFERENCES "business_profile" ("id");

ALTER TABLE
    IF EXISTS "content"
    ADD CONSTRAINT "content_story_fkey" FOREIGN KEY ("story_id") REFERENCES "story" ("id");

ALTER TABLE
    IF EXISTS "base_service"
    ADD CONSTRAINT "base_service_business_profile_fkey" FOREIGN KEY ("business_profile") REFERENCES "business_profile" ("id");

ALTER TABLE
    IF EXISTS "service"
    ADD CONSTRAINT "service_base_service_fkey" FOREIGN KEY ("base_service") REFERENCES "base_service" ("id");

ALTER TABLE
    IF EXISTS "commitment"
    ADD CONSTRAINT "commitment_business_profile_fkey" FOREIGN KEY ("business_profile") REFERENCES "business_profile" ("id");

ALTER TABLE
    IF EXISTS "partner"
    ADD CONSTRAINT "partner_business_profile_fkey" FOREIGN KEY ("business_profile") REFERENCES "business_profile" ("id");

ALTER TABLE
    IF EXISTS "business_partner"
    ADD CONSTRAINT "business_partner_partner_fkey" FOREIGN KEY ("partner_id") REFERENCES "partner" ("id");

ALTER TABLE
    IF EXISTS "why_us"
    ADD CONSTRAINT "why_us_commitment_fkey" FOREIGN KEY ("commitment") REFERENCES "commitment" ("id");

COMMIT;