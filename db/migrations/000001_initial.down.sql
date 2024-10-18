BEGIN;

ALTER TABLE
    IF EXISTS "user_credentials"
    DROP CONSTRAINT IF EXISTS "user_credentials_created_by_fkey";

ALTER TABLE
    IF EXISTS "content"
    DROP CONSTRAINT IF EXISTS "content_created_by_fkey";

ALTER TABLE
    IF EXISTS "content"
    DROP CONSTRAINT IF EXISTS "content_business_profile_fkey";

ALTER TABLE
    IF EXISTS "service"
    DROP CONSTRAINT IF EXISTS "service_base_service_fkey";

ALTER TABLE
    IF EXISTS "base_service"
    DROP CONSTRAINT IF EXISTS "base_service_business_profile_fkey";

ALTER TABLE
    IF EXISTS "why_us"
    DROP CONSTRAINT IF EXISTS "why_us_commitment_fkey";

ALTER TABLE
    IF EXISTS "commitment"
    DROP CONSTRAINT IF EXISTS "commitment_business_profile_fkey";

ALTER TABLE
    IF EXISTS "business_partner"
    DROP CONSTRAINT IF EXISTS "business_partner_business_profile_fkey";

ALTER TABLE
    IF EXISTS "content"
    DROP CONSTRAINT IF EXISTS "content_story_fkey";

ALTER TABLE
    IF EXISTS "story"
    DROP CONSTRAINT IF EXISTS "story_business_profile_fkey";

ALTER TABLE
    IF EXISTS "partner"
    DROP CONSTRAINT IF EXISTS "partner_business_profile_fkey";

ALTER TABLE
    IF EXISTS "business_partner"
    DROP CONSTRAINT IF EXISTS "business_partner_partner_fkey";

DROP TABLE IF EXISTS "content";

DROP TABLE IF EXISTS "user_credentials";

DROP TABLE IF EXISTS "users_user";

DROP TABLE IF EXISTS "business_partner";

DROP TABLE IF EXISTS "service";

DROP TABLE IF EXISTS "base_service";

DROP TABLE IF EXISTS "why_us";

DROP TABLE IF EXISTS "commitment";

DROP TABLE IF EXISTS "story";

DROP TABLE IF EXISTS "partner";

DROP TABLE IF EXISTS "business_profile";

COMMIT;