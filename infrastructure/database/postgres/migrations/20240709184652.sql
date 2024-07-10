table "cases" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "case_name" {
    null = true
    type = text
  }
  column "number" {
    null = true
    type = bigint
  }
  column "case_description" {
    null = true
    type = text
  }
  column "drill_depth" {
    null = true
    type = numeric
  }
  column "pipe_size" {
    null = true
    type = numeric
  }
  column "design_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_designs_cases" {
    columns     = [column.design_id]
    ref_columns = [table.designs.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_cases_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "companies" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "user_id" {
    null = false
    type = uuid
  }
  column "name" {
    null = false
    type = character_varying(255)
  }
  column "division" {
    null = true
    type = text
  }
  column "group" {
    null = true
    type = text
  }
  column "representative" {
    null = true
    type = text
  }
  column "address" {
    null = true
    type = text
  }
  column "phone" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_users_companies" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_companies_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "designs" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "plan_name" {
    null = true
    type = text
  }
  column "stage" {
    null = true
    type = text
  }
  column "version" {
    null = true
    type = text
  }
  column "actual_date" {
    null = true
    type = timestamptz
  }
  column "wellbore_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_wellbores_designs" {
    columns     = [column.wellbore_id]
    ref_columns = [table.wellbores.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_designs_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "fields" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "company_id" {
    null = false
    type = uuid
  }
  column "name" {
    null = false
    type = character_varying(255)
  }
  column "description" {
    null = true
    type = text
  }
  column "reduction_level" {
    null = true
    type = text
  }
  column "active_field_unit" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_companies_fields" {
    columns     = [column.company_id]
    ref_columns = [table.companies.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_fields_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "sites" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "field_id" {
    null = false
    type = uuid
  }
  column "name" {
    null = false
    type = character_varying(255)
  }
  column "area" {
    null = true
    type = numeric
  }
  column "block" {
    null = true
    type = text
  }
  column "azimuth" {
    null = true
    type = numeric
  }
  column "country" {
    null = true
    type = text
  }
  column "state" {
    null = true
    type = text
  }
  column "region" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_fields_sites" {
    columns     = [column.field_id]
    ref_columns = [table.fields.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_sites_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "trajectories" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "name" {
    null = true
    type = character_varying(255)
  }
  column "description" {
    null = true
    type = text
  }
  column "design_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_designs_trajectories" {
    columns     = [column.design_id]
    ref_columns = [table.designs.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_trajectories_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "trajectory_headers" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "trajectory_id" {
    null = false
    type = uuid
  }
  column "customer" {
    null = true
    type = text
  }
  column "project" {
    null = true
    type = text
  }
  column "profile_type" {
    null = true
    type = text
  }
  column "field" {
    null = true
    type = text
  }
  column "your_ref" {
    null = true
    type = text
  }
  column "structure" {
    null = true
    type = text
  }
  column "job_number" {
    null = true
    type = text
  }
  column "wellhead" {
    null = true
    type = text
  }
  column "kelly_bushing_elev" {
    null = true
    type = numeric
  }
  column "profile" {
    null = true
    type = text
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_trajectories_headers" {
    columns     = [column.trajectory_id]
    ref_columns = [table.trajectories.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
table "trajectory_units" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "trajectory_id" {
    null = false
    type = uuid
  }
  column "md" {
    null = true
    type = numeric
  }
  column "incl" {
    null = true
    type = numeric
  }
  column "azim" {
    null = true
    type = numeric
  }
  column "sub_sea" {
    null = true
    type = numeric
  }
  column "tvd" {
    null = true
    type = numeric
  }
  column "local_n_coord" {
    null = true
    type = numeric
  }
  column "local_e_coord" {
    null = true
    type = numeric
  }
  column "global_n_coord" {
    null = true
    type = numeric
  }
  column "global_e_coord" {
    null = true
    type = numeric
  }
  column "dogleg" {
    null = true
    type = numeric
  }
  column "vertical_section" {
    null = true
    type = numeric
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_trajectories_units" {
    columns     = [column.trajectory_id]
    ref_columns = [table.trajectories.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "name" {
    null = false
    type = text
  }
  column "surname" {
    null = false
    type = text
  }
  column "email" {
    null = false
    type = character_varying(255)
  }
  column "password" {
    null = false
    type = character_varying(70)
  }
  column "phone" {
    null = true
    type = character_varying(20)
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_users_deleted_at" {
    columns = [column.deleted_at]
  }
  unique "uni_users_email" {
    columns = [column.email]
  }
}
table "wellbores" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "name" {
    null = false
    type = character_varying(255)
  }
  column "bottom_hole_location" {
    null = true
    type = text
  }
  column "wellbore_depth" {
    null = true
    type = numeric
  }
  column "average_hook_load" {
    null = true
    type = numeric
  }
  column "riser_pressure" {
    null = true
    type = numeric
  }
  column "average_inlet_flow" {
    null = true
    type = numeric
  }
  column "average_column_rotation_frequency" {
    null = true
    type = numeric
  }
  column "maximum_column_rotation_frequency" {
    null = true
    type = numeric
  }
  column "average_weight_on_bit" {
    null = true
    type = numeric
  }
  column "maximum_weight_on_bit" {
    null = true
    type = numeric
  }
  column "average_torque" {
    null = true
    type = numeric
  }
  column "maximum_torque" {
    null = true
    type = numeric
  }
  column "down_static_friction" {
    null = true
    type = numeric
  }
  column "depth_interval" {
    null = true
    type = numeric
  }
  column "well_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_wells_wellbores" {
    columns     = [column.well_id]
    ref_columns = [table.wells.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_wellbores_deleted_at" {
    columns = [column.deleted_at]
  }
}
table "wells" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("public.uuid_generate_v4()")
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  column "site_id" {
    null = false
    type = uuid
  }
  column "name" {
    null = false
    type = character_varying(255)
  }
  column "description" {
    null = true
    type = text
  }
  column "location" {
    null = true
    type = text
  }
  column "universal_well_identifier" {
    null = true
    type = text
  }
  column "type" {
    null = true
    type = text
  }
  column "well_number" {
    null = true
    type = text
  }
  column "working_group" {
    null = true
    type = text
  }
  column "active_well_unit" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_sites_wells" {
    columns     = [column.site_id]
    ref_columns = [table.sites.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_wells_deleted_at" {
    columns = [column.deleted_at]
  }
}
schema "public" {
  comment = "standard public schema"
}
