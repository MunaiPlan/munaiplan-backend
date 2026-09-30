-- The baseline had no indexes on foreign keys, so every child lookup and the navigation
-- tree scanned whole tables. Index each parent reference (case_imports.case_id exists).

CREATE INDEX IF NOT EXISTS idx_caisings_hole_id ON public.caisings USING btree (hole_id);
CREATE INDEX IF NOT EXISTS idx_cases_trajectory_id ON public.cases USING btree (trajectory_id);
CREATE INDEX IF NOT EXISTS idx_companies_organization_id ON public.companies USING btree (organization_id);
CREATE INDEX IF NOT EXISTS idx_designs_wellbore_id ON public.designs USING btree (wellbore_id);
CREATE INDEX IF NOT EXISTS idx_fields_company_id ON public.fields USING btree (company_id);
CREATE INDEX IF NOT EXISTS idx_fluids_case_id ON public.fluids USING btree (case_id);
CREATE INDEX IF NOT EXISTS idx_fracture_gradients_case_id ON public.fracture_gradients USING btree (case_id);
CREATE INDEX IF NOT EXISTS idx_holes_case_id ON public.holes USING btree (case_id);
CREATE INDEX IF NOT EXISTS idx_pore_pressures_case_id ON public.pore_pressures USING btree (case_id);
CREATE INDEX IF NOT EXISTS idx_rigs_case_id ON public.rigs USING btree (case_id);
CREATE INDEX IF NOT EXISTS idx_sections_string_id ON public.sections USING btree (string_id);
CREATE INDEX IF NOT EXISTS idx_sites_field_id ON public.sites USING btree (field_id);
CREATE INDEX IF NOT EXISTS idx_strings_case_id ON public.strings USING btree (case_id);
CREATE INDEX IF NOT EXISTS idx_trajectories_design_id ON public.trajectories USING btree (design_id);
CREATE INDEX IF NOT EXISTS idx_trajectory_headers_trajectory_id ON public.trajectory_headers USING btree (trajectory_id);
CREATE INDEX IF NOT EXISTS idx_trajectory_units_trajectory_id ON public.trajectory_units USING btree (trajectory_id);
CREATE INDEX IF NOT EXISTS idx_users_organization_id ON public.users USING btree (organization_id);
CREATE INDEX IF NOT EXISTS idx_wellbores_well_id ON public.wellbores USING btree (well_id);
CREATE INDEX IF NOT EXISTS idx_wells_site_id ON public.wells USING btree (site_id);
CREATE INDEX IF NOT EXISTS idx_case_imports_organization_id ON public.case_imports USING btree (organization_id);
CREATE INDEX IF NOT EXISTS idx_trajectory_units_trajectory_md ON public.trajectory_units USING btree (trajectory_id, md);

-- Refresh planner statistics so the new indexes are used immediately.
ANALYZE;
