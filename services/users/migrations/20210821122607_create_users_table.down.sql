DROP EXTENSION IF EXISTS "uuid-ossp";

DROP TABLE IF EXISTS users;
DROP INDEX IF EXISTS users_email_unique_index;

DROP TABLE IF EXISTS onboarding_state;
DROP INDEX IF EXISTS onboarding_state_user_unique_index;