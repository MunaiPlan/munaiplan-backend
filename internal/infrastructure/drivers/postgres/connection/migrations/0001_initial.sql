-- Version 0001: PostgreSQL 16.9 schema snapshot generated from current GORM models, setup.sql, and indexes.sql.
-- Apply only to an empty disposable database. Do not edit after release; add a later migration.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: caisings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.caisings (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    hole_id uuid NOT NULL,
    md_top numeric,
    md_base numeric,
    length numeric,
    shoe_md numeric,
    od numeric,
    vd numeric,
    drift_id numeric,
    effective_hole_diameter numeric,
    weight numeric,
    grade text,
    min_yield_strength numeric,
    burst_rating numeric,
    collapse_rating numeric,
    friction_factor_caising numeric,
    linear_capacity_caising numeric,
    description_caising text,
    manufacturer_caising text,
    model_caising text
);


--
-- Name: cases; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.cases (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    case_name text,
    case_description text,
    drill_depth numeric,
    pipe_size numeric,
    is_complete boolean DEFAULT false NOT NULL,
    trajectory_id uuid NOT NULL
);


--
-- Name: companies; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.companies (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    organization_id uuid NOT NULL,
    name character varying(255) NOT NULL,
    division text,
    "group" text,
    representative text,
    address text,
    phone text
);


--
-- Name: designs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.designs (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    wellbore_id uuid NOT NULL,
    plan_name text,
    stage text,
    version text,
    actual_date timestamp with time zone
);


--
-- Name: fields; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.fields (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    company_id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    reduction_level text,
    active_field_unit text
);


--
-- Name: fluid_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.fluid_types (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL
);


--
-- Name: fluids; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.fluids (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    case_id uuid NOT NULL,
    name text NOT NULL,
    description text,
    density numeric NOT NULL,
    fluid_base_type_id uuid NOT NULL,
    base_fluid_id uuid NOT NULL
);


--
-- Name: fracture_gradients; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.fracture_gradients (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    case_id uuid NOT NULL,
    temperature_at_surface numeric NOT NULL,
    temperature_at_well_tvd numeric NOT NULL,
    temperature_gradient numeric NOT NULL,
    well_tvd numeric NOT NULL
);


--
-- Name: holes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.holes (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    case_id uuid NOT NULL,
    open_hole_md_top numeric,
    open_hole_md_base numeric,
    open_hole_length numeric,
    open_hole_vd numeric,
    effective_diameter numeric,
    friction_factor_open_hole numeric,
    linear_capacity_open_hole numeric,
    volume_excess numeric,
    description_open_hole text,
    tripping_in_casing numeric,
    tripping_out_casing numeric,
    rotating_on_bottom_casing numeric,
    slide_drilling_casing numeric,
    back_reaming_casing numeric,
    rotating_off_bottom_casing numeric,
    tripping_in_open_hole numeric,
    tripping_out_open_hole numeric,
    rotating_on_bottom_open_hole numeric,
    slide_drilling_open_hole numeric,
    back_reaming_open_hole numeric,
    rotating_off_bottom_open_hole numeric
);


--
-- Name: library_sections; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.library_sections (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    description text,
    manufacturer text,
    type text NOT NULL,
    body_od numeric NOT NULL,
    body_id numeric NOT NULL,
    avg_joint_length numeric,
    stabilizer_length numeric,
    stabilizer_od numeric,
    stabilizer_id numeric,
    weight numeric,
    material text,
    grade text,
    class bigint,
    friction_coefficient numeric,
    min_yield_strength numeric,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: organizations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.organizations (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    email character varying(255) NOT NULL,
    phone character varying(20),
    address text
);


--
-- Name: pore_pressures; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.pore_pressures (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    case_id uuid NOT NULL,
    tvd numeric NOT NULL,
    pressure numeric NOT NULL,
    emw numeric NOT NULL
);


--
-- Name: rigs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.rigs (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    case_id uuid NOT NULL,
    block_rating numeric,
    torque_rating numeric,
    rated_working_pressure numeric NOT NULL,
    bop_pressure_rating numeric NOT NULL,
    surface_pressure_loss numeric DEFAULT 0 NOT NULL,
    standpipe_length numeric,
    standpipe_internal_diameter numeric,
    hose_length numeric,
    hose_internal_diameter numeric,
    swivel_length numeric,
    swivel_internal_diameter numeric,
    kelly_length numeric,
    kelly_internal_diameter numeric,
    pump_discharge_line_length numeric,
    pump_discharge_line_internal_diameter numeric,
    top_drive_stackup_length numeric,
    top_drive_stackup_internal_diameter numeric
);


--
-- Name: sections; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sections (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    string_id uuid NOT NULL,
    description text,
    manufacturer text,
    type text NOT NULL,
    body_md numeric NOT NULL,
    body_length numeric NOT NULL,
    body_od numeric NOT NULL,
    body_id numeric NOT NULL,
    avg_joint_length numeric,
    stabilizer_length numeric,
    stabilizer_od numeric,
    stabilizer_id numeric,
    weight numeric,
    material text,
    grade text,
    class bigint,
    friction_coefficient numeric,
    min_yield_strength numeric,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: sites; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sites (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    field_id uuid NOT NULL,
    name character varying(255) NOT NULL,
    area numeric,
    block text,
    azimuth numeric,
    country text,
    state text,
    region text
);


--
-- Name: strings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.strings (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    case_id uuid NOT NULL,
    name text NOT NULL,
    depth numeric NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: trajectories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.trajectories (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    design_id uuid NOT NULL,
    name character varying(255),
    description text
);


--
-- Name: trajectory_headers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.trajectory_headers (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    trajectory_id uuid NOT NULL,
    customer text,
    project text,
    profile_type text,
    field text,
    your_ref text,
    structure text,
    job_number text,
    wellhead text,
    kelly_bushing_elev numeric,
    profile text
);


--
-- Name: trajectory_units; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.trajectory_units (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    trajectory_id uuid NOT NULL,
    md numeric,
    incl numeric,
    azim numeric,
    sub_sea numeric,
    tvd numeric,
    local_n_coord numeric,
    local_e_coord numeric,
    global_n_coord numeric,
    global_e_coord numeric,
    dogleg numeric,
    vertical_section numeric
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    organization_id uuid NOT NULL,
    name text NOT NULL,
    surname text NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(70) NOT NULL,
    phone character varying(20)
);


--
-- Name: wellbores; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.wellbores (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    well_id uuid NOT NULL,
    name character varying(255) NOT NULL,
    bottom_hole_location text,
    wellbore_depth numeric,
    average_hook_load numeric,
    riser_pressure numeric,
    average_inlet_flow numeric,
    average_column_rotation_frequency numeric,
    maximum_column_rotation_frequency numeric,
    average_weight_on_bit numeric,
    maximum_weight_on_bit numeric,
    average_torque numeric,
    maximum_torque numeric,
    down_static_friction numeric,
    depth_interval numeric
);


--
-- Name: wells; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.wells (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    site_id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    location text,
    universal_well_identifier text,
    type text,
    well_number text,
    working_group text,
    active_well_unit text
);


--
-- Name: caisings caisings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.caisings
    ADD CONSTRAINT caisings_pkey PRIMARY KEY (id);


--
-- Name: cases cases_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cases
    ADD CONSTRAINT cases_pkey PRIMARY KEY (id);


--
-- Name: companies companies_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (id);


--
-- Name: designs designs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.designs
    ADD CONSTRAINT designs_pkey PRIMARY KEY (id);


--
-- Name: fields fields_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fields
    ADD CONSTRAINT fields_pkey PRIMARY KEY (id);


--
-- Name: fluid_types fluid_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fluid_types
    ADD CONSTRAINT fluid_types_pkey PRIMARY KEY (id);


--
-- Name: fluids fluids_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fluids
    ADD CONSTRAINT fluids_pkey PRIMARY KEY (id);


--
-- Name: fracture_gradients fracture_gradients_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fracture_gradients
    ADD CONSTRAINT fracture_gradients_pkey PRIMARY KEY (id);


--
-- Name: holes holes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.holes
    ADD CONSTRAINT holes_pkey PRIMARY KEY (id);


--
-- Name: library_sections library_sections_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.library_sections
    ADD CONSTRAINT library_sections_pkey PRIMARY KEY (id);


--
-- Name: organizations organizations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.organizations
    ADD CONSTRAINT organizations_pkey PRIMARY KEY (id);


--
-- Name: pore_pressures pore_pressures_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.pore_pressures
    ADD CONSTRAINT pore_pressures_pkey PRIMARY KEY (id);


--
-- Name: rigs rigs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rigs
    ADD CONSTRAINT rigs_pkey PRIMARY KEY (id);


--
-- Name: sections sections_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sections
    ADD CONSTRAINT sections_pkey PRIMARY KEY (id);


--
-- Name: sites sites_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sites
    ADD CONSTRAINT sites_pkey PRIMARY KEY (id);


--
-- Name: strings strings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.strings
    ADD CONSTRAINT strings_pkey PRIMARY KEY (id);


--
-- Name: trajectories trajectories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.trajectories
    ADD CONSTRAINT trajectories_pkey PRIMARY KEY (id);


--
-- Name: trajectory_headers trajectory_headers_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.trajectory_headers
    ADD CONSTRAINT trajectory_headers_pkey PRIMARY KEY (id);


--
-- Name: trajectory_units trajectory_units_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.trajectory_units
    ADD CONSTRAINT trajectory_units_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: wellbores wellbores_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wellbores
    ADD CONSTRAINT wellbores_pkey PRIMARY KEY (id);


--
-- Name: wells wells_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wells
    ADD CONSTRAINT wells_pkey PRIMARY KEY (id);


--
-- Name: idx_caisings_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_caisings_deleted_at ON public.caisings USING btree (deleted_at);


--
-- Name: idx_cases_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_cases_deleted_at ON public.cases USING btree (deleted_at);


--
-- Name: idx_companies_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_companies_deleted_at ON public.companies USING btree (deleted_at);


--
-- Name: idx_companies_name; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_companies_name ON public.companies USING btree (name) WHERE (deleted_at IS NULL);


--
-- Name: idx_designs_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_designs_deleted_at ON public.designs USING btree (deleted_at);


--
-- Name: idx_fields_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_fields_deleted_at ON public.fields USING btree (deleted_at);


--
-- Name: idx_fluid_types_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_fluid_types_deleted_at ON public.fluid_types USING btree (deleted_at);


--
-- Name: idx_fluids_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_fluids_deleted_at ON public.fluids USING btree (deleted_at);


--
-- Name: idx_fracture_gradients_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_fracture_gradients_deleted_at ON public.fracture_gradients USING btree (deleted_at);


--
-- Name: idx_holes_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_holes_deleted_at ON public.holes USING btree (deleted_at);


--
-- Name: idx_library_sections_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_library_sections_deleted_at ON public.library_sections USING btree (deleted_at);


--
-- Name: idx_organizations_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_organizations_deleted_at ON public.organizations USING btree (deleted_at);


--
-- Name: idx_organizations_email; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_organizations_email ON public.organizations USING btree (email) WHERE (deleted_at IS NULL);


--
-- Name: idx_pore_pressures_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_pore_pressures_deleted_at ON public.pore_pressures USING btree (deleted_at);


--
-- Name: idx_rigs_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_rigs_deleted_at ON public.rigs USING btree (deleted_at);


--
-- Name: idx_sections_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sections_deleted_at ON public.sections USING btree (deleted_at);


--
-- Name: idx_sites_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_sites_deleted_at ON public.sites USING btree (deleted_at);


--
-- Name: idx_strings_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_strings_deleted_at ON public.strings USING btree (deleted_at);


--
-- Name: idx_trajectories_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_trajectories_deleted_at ON public.trajectories USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email) WHERE (deleted_at IS NULL);


--
-- Name: idx_wellbores_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_wellbores_deleted_at ON public.wellbores USING btree (deleted_at);


--
-- Name: idx_wells_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_wells_deleted_at ON public.wells USING btree (deleted_at);


--
-- Name: fluids fk_cases_fluids; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fluids
    ADD CONSTRAINT fk_cases_fluids FOREIGN KEY (case_id) REFERENCES public.cases(id) ON DELETE CASCADE;


--
-- Name: fracture_gradients fk_cases_fracture_gradients; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fracture_gradients
    ADD CONSTRAINT fk_cases_fracture_gradients FOREIGN KEY (case_id) REFERENCES public.cases(id) ON DELETE CASCADE;


--
-- Name: holes fk_cases_holes; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.holes
    ADD CONSTRAINT fk_cases_holes FOREIGN KEY (case_id) REFERENCES public.cases(id) ON DELETE CASCADE;


--
-- Name: pore_pressures fk_cases_pore_pressures; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.pore_pressures
    ADD CONSTRAINT fk_cases_pore_pressures FOREIGN KEY (case_id) REFERENCES public.cases(id) ON DELETE CASCADE;


--
-- Name: rigs fk_cases_rigs; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rigs
    ADD CONSTRAINT fk_cases_rigs FOREIGN KEY (case_id) REFERENCES public.cases(id) ON DELETE CASCADE;


--
-- Name: strings fk_cases_strings; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.strings
    ADD CONSTRAINT fk_cases_strings FOREIGN KEY (case_id) REFERENCES public.cases(id) ON DELETE CASCADE;


--
-- Name: fields fk_companies_fields; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fields
    ADD CONSTRAINT fk_companies_fields FOREIGN KEY (company_id) REFERENCES public.companies(id) ON DELETE CASCADE;


--
-- Name: trajectories fk_designs_trajectories; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.trajectories
    ADD CONSTRAINT fk_designs_trajectories FOREIGN KEY (design_id) REFERENCES public.designs(id) ON DELETE CASCADE;


--
-- Name: sites fk_fields_sites; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sites
    ADD CONSTRAINT fk_fields_sites FOREIGN KEY (field_id) REFERENCES public.fields(id) ON DELETE CASCADE;


--
-- Name: fluids fk_fluids_base_fluid; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fluids
    ADD CONSTRAINT fk_fluids_base_fluid FOREIGN KEY (base_fluid_id) REFERENCES public.fluid_types(id) ON DELETE CASCADE;


--
-- Name: fluids fk_fluids_fluid_base_type; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fluids
    ADD CONSTRAINT fk_fluids_fluid_base_type FOREIGN KEY (fluid_base_type_id) REFERENCES public.fluid_types(id) ON DELETE CASCADE;


--
-- Name: caisings fk_holes_caisings; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.caisings
    ADD CONSTRAINT fk_holes_caisings FOREIGN KEY (hole_id) REFERENCES public.holes(id) ON DELETE CASCADE;


--
-- Name: companies fk_organizations_companies; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT fk_organizations_companies FOREIGN KEY (organization_id) REFERENCES public.organizations(id) ON DELETE CASCADE;


--
-- Name: users fk_organizations_users; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_organizations_users FOREIGN KEY (organization_id) REFERENCES public.organizations(id) ON DELETE CASCADE;


--
-- Name: wells fk_sites_wells; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wells
    ADD CONSTRAINT fk_sites_wells FOREIGN KEY (site_id) REFERENCES public.sites(id) ON DELETE CASCADE;


--
-- Name: sections fk_strings_sections; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sections
    ADD CONSTRAINT fk_strings_sections FOREIGN KEY (string_id) REFERENCES public.strings(id) ON DELETE CASCADE;


--
-- Name: cases fk_trajectories_cases; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cases
    ADD CONSTRAINT fk_trajectories_cases FOREIGN KEY (trajectory_id) REFERENCES public.trajectories(id) ON DELETE CASCADE;


--
-- Name: trajectory_headers fk_trajectories_headers; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.trajectory_headers
    ADD CONSTRAINT fk_trajectories_headers FOREIGN KEY (trajectory_id) REFERENCES public.trajectories(id) ON DELETE CASCADE;


--
-- Name: trajectory_units fk_trajectories_units; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.trajectory_units
    ADD CONSTRAINT fk_trajectories_units FOREIGN KEY (trajectory_id) REFERENCES public.trajectories(id) ON DELETE CASCADE;


--
-- Name: designs fk_wellbores_designs; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.designs
    ADD CONSTRAINT fk_wellbores_designs FOREIGN KEY (wellbore_id) REFERENCES public.wellbores(id) ON DELETE CASCADE;


--
-- Name: wellbores fk_wells_wellbores; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wellbores
    ADD CONSTRAINT fk_wells_wellbores FOREIGN KEY (well_id) REFERENCES public.wells(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

