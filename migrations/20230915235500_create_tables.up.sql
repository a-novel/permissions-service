CREATE TABLE IF NOT EXISTS users_permissions (
    id uuid PRIMARY KEY NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,

    validated_account BOOLEAN NOT NULL DEFAULT FALSE,
    admin_access BOOLEAN NOT NULL DEFAULT FALSE
);
