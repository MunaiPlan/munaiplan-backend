-- Company names are unique within an organization, not across tenants.
DROP INDEX public.idx_companies_name;
CREATE UNIQUE INDEX idx_companies_name ON public.companies USING btree (organization_id, name) WHERE (deleted_at IS NULL);

-- WellPlan hole sections report the casing inner diameter (mm).
ALTER TABLE public.caisings ADD COLUMN inner_diameter numeric;

-- One row per imported WellPlan case: provenance plus the full parsed report, including
-- WellPlan's own Torque & Drag and hydraulics results used as reference values.
CREATE TABLE public.case_imports (
    id uuid DEFAULT public.uuid_generate_v4() PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    organization_id uuid NOT NULL REFERENCES public.organizations(id) ON DELETE CASCADE,
    case_id uuid NOT NULL REFERENCES public.cases(id) ON DELETE CASCADE,
    created_by uuid REFERENCES public.users(id) ON DELETE SET NULL,
    format text NOT NULL,
    fingerprint text NOT NULL,
    report jsonb NOT NULL
);
CREATE UNIQUE INDEX idx_case_imports_fingerprint ON public.case_imports (organization_id, fingerprint);
CREATE INDEX idx_case_imports_case_id ON public.case_imports (case_id);
