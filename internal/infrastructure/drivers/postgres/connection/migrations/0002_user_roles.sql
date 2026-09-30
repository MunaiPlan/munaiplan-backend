-- Accounts are provisioned by administrators (B2B). A role distinguishes them.
ALTER TABLE public.users ADD COLUMN role text NOT NULL DEFAULT 'user';
ALTER TABLE public.users ADD CONSTRAINT users_role_check CHECK (role IN ('user', 'admin'));
