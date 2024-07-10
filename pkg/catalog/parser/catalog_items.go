package parser

// Default Catalog Header
type CatalogHeader struct {
	CatalogID       string `xml:"CATALOG_ID,attr"`
	CatalogTypeCode string `xml:"CATALOG_TYPE_CODE,attr"`
	Name            string `xml:"NAME,attr"`
	IsAPI           string `xml:"IS_API,attr"`
	Description     string `xml:"DESCRIPTION,attr"`
	ReadOnly        string `xml:"READ_ONLY,attr"`
	Comments        string `xml:"COMMENTS,attr"`
	Author          string `xml:"AUTHOR,attr"`
	CreateDate      string `xml:"CREATE_DATE,attr"`
	CreateUserID    string `xml:"CREATE_USER_ID,attr"`
	CreateAppID     string `xml:"CREATE_APP_ID,attr"`
	UpdateDate      string `xml:"UPDATE_DATE,attr"`
	UpdateUserID    string `xml:"UPDATE_USER_ID,attr"`
	UpdateAppID     string `xml:"UPDATE_APP_ID,attr"`
}

// Api Drill Collar Catalog
type ApiDrillCollarCatalogItem struct {
	OdNominal             string `xml:"od_nominal,attr"`
	CompTypeCode          string `xml:"comp_type_code,attr"`
	CreateUserId          string `xml:"create_user_id,attr"`
	ApproximateWeight     string `xml:"approximate_weight,attr"`
	ClosedEndDisplacement string `xml:"closed_end_displacement,attr"`
	AverageJointLength    string `xml:"average_joint_length,attr"`
	IdBody                string `xml:"id_body,attr"`
	Connection            string `xml:"connection,attr"`
	IdNominal             string `xml:"id_nominal,attr"`
	MakeupTorque          string `xml:"makeup_torque,attr"`
	OdBody                string `xml:"od_body,attr"`
	CatalogItemId         string `xml:"catalog_item_id,attr"`
	LinearCapacity        string `xml:"linear_capacity,attr"`
	ApiIndicator          string `xml:"api_indicator,attr"`
	GradeId               string `xml:"grade_id,attr"`
	SectTypeCode          string `xml:"sect_type_code,attr"`
	CatalogId             string `xml:"catalog_id,attr"`
}

// Api Drill Pipe Catalog
type ApiDrillPipeCatalogItem struct {
    ServiceClass            string `xml:"service_class,attr"`
    ToolJointLength         string `xml:"tool_joint_length,attr"`
    CompTypeCode            string `xml:"comp_type_code,attr"`
    CreateUserId            string `xml:"create_user_id,attr"`
    NominalDiameter         string `xml:"nominal_diameter,attr"`
    NominalWeight           string `xml:"nominal_weight,attr"`
    ApproximateWeight       string `xml:"approximate_weight,attr"`
    IdConnection            string `xml:"id_connection,attr"`
    ClosedEndDisplacement   string `xml:"closed_end_displacement,attr"`
    AverageJointLength      string `xml:"average_joint_length,attr"`
    IdBody                  string `xml:"id_body,attr"`
    WallThicknessPercent    string `xml:"wall_thickness_percent,attr"`
    Connection              string `xml:"connection,attr"`
    MakeupTorque            string `xml:"makeup_torque,attr"`
    ConnectionTorsionalYield string `xml:"connection_torsional_yield,attr"`
    OdBody                  string `xml:"od_body,attr"`
    OdConnection            string `xml:"od_connection,attr"`
    CatalogItemId           string `xml:"catalog_item_id,attr"`
    LinearCapacity          string `xml:"linear_capacity,attr"`
    ApiIndicator            string `xml:"api_indicator,attr"`
    SectTypeCode            string `xml:"sect_type_code,attr"`
    GradeId                 string `xml:"grade_id,attr"`
    NominalWeightMeasure    string `xml:"nominal_weight_measure,attr"`
    CatalogId               string `xml:"catalog_id,attr"`
}

// Additional Catalog
type ApiCentralizerCatalogItem struct {
    StartingForce          string `xml:"starting_force,attr"`
    CasingDiameter         string `xml:"casing_diameter,attr"`
    PartNumber             string `xml:"part_number,attr"`
    CreateUserId           string `xml:"create_user_id,attr"`
    Type                   string `xml:"type,attr"`
    NominalDiameterMeasure string `xml:"nominal_diameter_measure,attr"`
    Description            string `xml:"description,attr"`
    RunningForce           string `xml:"running_force,attr"`
    RestoringForce         string `xml:"restoring_force,attr"`
    HoleDiameter           string `xml:"hole_diameter,attr"`
    Bows                   string `xml:"bows,attr"`
    NonFixed               string `xml:"non_fixed,attr"`
    MinimumDiameter        string `xml:"minimum_diameter,attr"`
    CatalogItemId          string `xml:"catalog_item_id,attr"`
    CatalogId              string `xml:"catalog_id,attr"`
}

// Adjustable Gauge Stabilizers Catalog
type ApiStabCatalogItem struct {
    EccStabBladeOd          string `xml:"ecc_stab_blade_od,attr"`
    CompTypeCode            string `xml:"comp_type_code,attr"`
    NominalSize             string `xml:"nominal_size,attr"`
    CreateUserId            string `xml:"create_user_id,attr"`
    ApproximateWeight       string `xml:"approximate_weight,attr"`
    ClosedEndDisplacement   string `xml:"closed_end_displacement,attr"`
    Description             string `xml:"description,attr"`
    Length                  string `xml:"length,attr"`
    IdBody                  string `xml:"id_body,attr"`
    Connection              string `xml:"connection,attr"`
    MakeupTorque            string `xml:"makeup_torque,attr"`
    OdBody                  string `xml:"od_body,attr"`
    CatalogItemId           string `xml:"catalog_item_id,attr"`
    LinearCapacity          string `xml:"linear_capacity,attr"`
    FishneckLength          string `xml:"fishneck_length,attr"`
    GradeId                 string `xml:"grade_id,attr"`
    SectTypeCode            string `xml:"sect_type_code,attr"`
    EccStabBladeLength      string `xml:"ecc_stab_blade_length,attr"`
    CatalogId               string `xml:"catalog_id,attr"`
}
