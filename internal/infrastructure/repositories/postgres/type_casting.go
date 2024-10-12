package postgres

import (
	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/types"
)

// toGormOrganization converts a domain organization to a gorm organization.
func toGormOrganization(organization *entities.Organization) *models.Organization {
	orgID, err := validateGormId(organization.ID)
	if err != nil {
		return nil
	}

	org := &models.Organization{
		ID:      orgID,
		Name:    organization.Name,
		Email:   organization.Email,
		Phone:   organization.Phone,
		Address: organization.Address,
	}

	for _, company := range organization.Companies {
		gormCompany := toGormCompany(company)
		org.Companies = append(org.Companies, *gormCompany)
	}
	return org
}

// toDomainOrganization converts a gorm organization to a domain organization.
func toDomainOrganization(organization *models.Organization) *entities.Organization {
	org := entities.Organization{
		ID:      organization.ID.String(),
		Name:    organization.Name,
		Email:   organization.Email,
		Phone:   organization.Phone,
		Address: organization.Address,
	}

	for _, company := range organization.Companies {
		domainCompany := toDomainCompany(&company)
		org.Companies = append(org.Companies, domainCompany)
	}
	return &org
}

// toDomainCompany maps the GORM Company model to the domain Company entity.
func toDomainCompany(companyModel *models.Company) *entities.Company {
	company := entities.Company{
		ID:             companyModel.ID.String(),
		Name:           companyModel.Name,
		Division:       companyModel.Division,
		Group:          companyModel.Group,
		Representative: companyModel.Representative,
		Address:        companyModel.Address,
		Phone:          companyModel.Phone,
	}

	for _, field := range companyModel.Fields {
		domainField := toDomainField(&field)
		company.Fields = append(company.Fields, domainField)
	}
	return &company
}

// toGormCompany maps the domain Company entity to the GORM Company model.
func toGormCompany(company *entities.Company) *models.Company {
	companyID, err := validateGormId(company.ID)
	if err != nil {
		return nil
	}

	comp := &models.Company{
		ID:             companyID,
		Name:           company.Name,
		Division:       company.Division,
		Group:          company.Group,
		Representative: company.Representative,
		Address:        company.Address,
		Phone:          company.Phone,
	}

	for _, field := range company.Fields {
		gormField := toGormField(field)
		comp.Fields = append(comp.Fields, *gormField)
	}
	return comp
}

// toDomainSite maps the GORM Site model to the domain Site entity.
func toDomainSite(siteModel *models.Site) *entities.Site {
	site := entities.Site{
		ID:      siteModel.ID.String(),
		Name:    siteModel.Name,
		Area:    siteModel.Area,
		Block:   siteModel.Block,
		Azimuth: siteModel.Azimuth,
		Country: siteModel.Country,
		State:   siteModel.State,
		Region:  siteModel.Region,
	}

	for _, well := range siteModel.Wells {
		domainWell := toDomainWell(&well)
		site.Wells = append(site.Wells, domainWell)
	}
	return &site
}

// toGormSite maps the domain Site entity to the GORM Site model.
func toGormSite(site *entities.Site) *models.Site {
	siteID, err := validateGormId(site.ID)
	if err != nil {
		return nil
	}

	newSite := &models.Site{
		ID:      siteID,
		Name:    site.Name,
		Area:    site.Area,
		Block:   site.Block,
		Azimuth: site.Azimuth,
		Country: site.Country,
		State:   site.State,
		Region:  site.Region,
	}

	for _, well := range site.Wells {
		gormWell := toGormWell(well)
		newSite.Wells = append(newSite.Wells, *gormWell)
	}

	return newSite
}

// toDomainWell maps the GORM Well model to the domain Well entity.
func toDomainWell(wellModel *models.Well) *entities.Well {
	well := entities.Well{
		ID:                      wellModel.ID.String(),
		Name:                    wellModel.Name,
		Description:             wellModel.Description,
		Location:                wellModel.Location,
		UniversalWellIdentifier: wellModel.UniversalWellIdentifier,
		Type:                    wellModel.Type,
		WellNumber:              wellModel.WellNumber,
		WorkingGroup:            wellModel.WorkingGroup,
		ActiveWellUnit:          wellModel.ActiveWellUnit,
	}

	for _, wellbore := range wellModel.Wellbores {
		domainWellbore := toDomainWellbore(&wellbore)
		well.Wellbores = append(well.Wellbores, domainWellbore)
	}
	return &well
}

// toGormWell maps the domain Well entity to the GORM Well model.
func toGormWell(well *entities.Well) *models.Well {
	wellID, err := validateGormId(well.ID)
	if err != nil {
		return nil
	}

	newWell := &models.Well{
		ID:                      wellID,
		Name:                    well.Name,
		Description:             well.Description,
		Location:                well.Location,
		UniversalWellIdentifier: well.UniversalWellIdentifier,
		Type:                    well.Type,
		WellNumber:              well.WellNumber,
		WorkingGroup:            well.WorkingGroup,
		ActiveWellUnit:          well.ActiveWellUnit,
	}

	for _, wellbore := range well.Wellbores {
		gormWellbore := toGormWellbore(wellbore)
		newWell.Wellbores = append(newWell.Wellbores, *gormWellbore)
	}

	return newWell
}

// toDomainWellbore maps the GORM Wellbore model to the domain Wellbore entity.
func toDomainWellbore(wellboreModel *models.Wellbore) *entities.Wellbore {
	wellbore := entities.Wellbore{
		ID:                             wellboreModel.ID.String(),
		Name:                           wellboreModel.Name,
		BottomHoleLocation:             wellboreModel.BottomHoleLocation,
		WellboreDepth:                  wellboreModel.WellboreDepth,
		AverageHookLoad:                wellboreModel.AverageHookLoad,
		RiserPressure:                  wellboreModel.RiserPressure,
		AverageInletFlow:               wellboreModel.AverageInletFlow,
		AverageColumnRotationFrequency: wellboreModel.AverageColumnRotationFrequency,
		MaximumColumnRotationFrequency: wellboreModel.MaximumColumnRotationFrequency,
		AverageWeightOnBit:             wellboreModel.AverageWeightOnBit,
		MaximumWeightOnBit:             wellboreModel.MaximumWeightOnBit,
		AverageTorque:                  wellboreModel.AverageTorque,
		MaximumTorque:                  wellboreModel.MaximumTorque,
		DownStaticFriction:             wellboreModel.DownStaticFriction,
		DepthInterval:                  wellboreModel.DepthInterval,
	}

	for _, design := range wellboreModel.Designs {
		domainDesign := toDomainDesign(&design)
		wellbore.Designs = append(wellbore.Designs, domainDesign)
	}
	return &wellbore
}

// toGormWellbore maps the domain Wellbore entity to the GORM Wellbore model.
func toGormWellbore(wellbore *entities.Wellbore) *models.Wellbore {
	wellboreID, err := validateGormId(wellbore.ID)
	if err != nil {
		return nil
	}

	newWellbore := &models.Wellbore{
		ID:                             wellboreID,
		Name:                           wellbore.Name,
		BottomHoleLocation:             wellbore.BottomHoleLocation,
		WellboreDepth:                  wellbore.WellboreDepth,
		AverageHookLoad:                wellbore.AverageHookLoad,
		RiserPressure:                  wellbore.RiserPressure,
		AverageInletFlow:               wellbore.AverageInletFlow,
		AverageColumnRotationFrequency: wellbore.AverageColumnRotationFrequency,
		MaximumColumnRotationFrequency: wellbore.MaximumColumnRotationFrequency,
		AverageWeightOnBit:             wellbore.AverageWeightOnBit,
		MaximumWeightOnBit:             wellbore.MaximumWeightOnBit,
		AverageTorque:                  wellbore.AverageTorque,
		MaximumTorque:                  wellbore.MaximumTorque,
		DownStaticFriction:             wellbore.DownStaticFriction,
		DepthInterval:                  wellbore.DepthInterval,
	}

	for _, design := range wellbore.Designs {
		gormDesign := toGormDesign(design)
		newWellbore.Designs = append(newWellbore.Designs, *gormDesign)
	}

	return newWellbore
}

// toDomainField maps the GORM Field model to the domain Field entity.
func toDomainField(fieldModel *models.Field) *entities.Field {
	field := entities.Field{
		ID:              fieldModel.ID.String(),
		Name:            fieldModel.Name,
		Description:     fieldModel.Description,
		ReductionLevel:  fieldModel.ReductionLevel,
		ActiveFieldUnit: fieldModel.ActiveFieldUnit,
	}

	for _, site := range fieldModel.Sites {
		domainSite := toDomainSite(&site)
		field.Sites = append(field.Sites, domainSite)
	}
	return &field
}

// toGormField maps the domain Field entity to the GORM Field model.
func toGormField(field *entities.Field) *models.Field {
	fieldID, err := validateGormId(field.ID)
	if err != nil {
		return nil
	}

	newField := &models.Field{
		ID:              fieldID,
		Name:            field.Name,
		Description:     field.Description,
		ReductionLevel:  field.ReductionLevel,
		ActiveFieldUnit: field.ActiveFieldUnit,
	}

	for _, site := range field.Sites {
		gormSite := toGormSite(site)
		newField.Sites = append(newField.Sites, *gormSite)
	}

	return newField
}

// toDomainDesign maps the GORM Design model to the domain Design entity.
func toDomainDesign(designModel *models.Design) *entities.Design {
	design := entities.Design{
		ID:         designModel.ID.String(),
		PlanName:   designModel.PlanName,
		Stage:      designModel.Stage,
		Version:    designModel.Version,
		ActualDate: designModel.ActualDate,
	}

	for _, trajectory := range designModel.Trajectories {
		domainTrajectory := toDomainTrajectory(&trajectory)
		design.Trajectories = append(design.Trajectories, domainTrajectory)
	}
	return &design
}

// toGormDesign maps the domain Design entity to the GORM Design model.
func toGormDesign(design *entities.Design) *models.Design {
	designID, err := validateGormId(design.ID)
	if err != nil {
		return nil
	}

	newDesign := &models.Design{
		ID:         designID,
		PlanName:   design.PlanName,
		Stage:      design.Stage,
		Version:    design.Version,
		ActualDate: design.ActualDate,
	}

	for _, trajectory := range design.Trajectories {
		gormTrajectory := toGormTrajectory(trajectory)
		newDesign.Trajectories = append(newDesign.Trajectories, *gormTrajectory)
	}

	return newDesign
}

// toDomainTrajectory maps the GORM Trajectory model to the domain Trajectory entity.
func toDomainTrajectory(trajectoryModel *models.Trajectory) *entities.Trajectory {
	headers := make([]*entities.TrajectoryHeader, len(trajectoryModel.Headers))
	units := make([]*entities.TrajectoryUnit, len(trajectoryModel.Units))
	cases := make([]*entities.Case, len(trajectoryModel.Cases))

	for i, h := range trajectoryModel.Headers {
		headers[i] = toDomainTrajectoryHeader(&h)
	}

	for i, u := range trajectoryModel.Units {
		units[i] = toDomainTrajectoryUnit(&u)
	}

	for i, c := range trajectoryModel.Cases {
		cases[i] = toDomainCase(&c)
	}

	return &entities.Trajectory{
		ID:          trajectoryModel.ID.String(),
		Name:        trajectoryModel.Name,
		Description: trajectoryModel.Description,
		Headers:     headers,
		Units:       units,
		Cases:       cases,
		CreatedAt:   trajectoryModel.CreatedAt,
	}
}

// toGormTrajectory maps the domain Trajectory entity to the GORM Trajectory model.
func toGormTrajectory(trajectory *entities.Trajectory) *models.Trajectory {
	trajectoryID, err := validateGormId(trajectory.ID)
	if err != nil {
		return nil
	}

	headers := make([]models.TrajectoryHeader, len(trajectory.Headers))
	units := make([]models.TrajectoryUnit, len(trajectory.Units))
	cases := make([]models.Case, len(trajectory.Cases))
	for i, h := range trajectory.Headers {
		headers[i] = *toGormTrajectoryHeader(h)
	}

	for i, u := range trajectory.Units {
		units[i] = *toGormTrajectoryUnit(u)
	}

	for i, c := range trajectory.Cases {
		cases[i] = *toGormCase(c)
	}

	return &models.Trajectory{
		ID:          trajectoryID,
		Name:        trajectory.Name,
		Description: trajectory.Description,
		Headers:     headers,
		Units:       units,
		Cases:       cases,
	}
}

// toDomainTrajectoryHeader maps the GORM TrajectoryHeader model to the domain TrajectoryHeader entity.
func toGormTrajectoryHeader(header *entities.TrajectoryHeader) *models.TrajectoryHeader {
	headerID, err := validateGormId(header.ID)
	if err != nil {
		return nil
	}

	return &models.TrajectoryHeader{
		ID:               headerID,
		Customer:         header.Customer,
		Project:          header.Project,
		ProfileType:      header.ProfileType,
		Field:            header.Field,
		YourRef:          header.YourRef,
		Structure:        header.Structure,
		JobNumber:        header.JobNumber,
		Wellhead:         header.Wellhead,
		KellyBushingElev: header.KellyBushingElev,
		Profile:          header.Profile,
	}
}

// toDomainTrajectoryHeader maps the GORM TrajectoryHeader model to the domain TrajectoryHeader entity.
func toDomainTrajectoryHeader(headerModel *models.TrajectoryHeader) *entities.TrajectoryHeader {
	return &entities.TrajectoryHeader{
		ID:               headerModel.ID.String(),
		Customer:         headerModel.Customer,
		Project:          headerModel.Project,
		ProfileType:      headerModel.ProfileType,
		Field:            headerModel.Field,
		YourRef:          headerModel.YourRef,
		Structure:        headerModel.Structure,
		JobNumber:        headerModel.JobNumber,
		Wellhead:         headerModel.Wellhead,
		KellyBushingElev: headerModel.KellyBushingElev,
		Profile:          headerModel.Profile,
		CreatedAt:        headerModel.CreatedAt,
	}
}

// toGormTrajectoryUnit maps the domain TrajectoryUnit entity to the GORM TrajectoryUnit model.
func toGormTrajectoryUnit(unit *entities.TrajectoryUnit) *models.TrajectoryUnit {
	unitID, err := validateGormId(unit.ID)
	if err != nil {
		return nil
	}

	return &models.TrajectoryUnit{
		ID:              unitID,
		MD:              unit.MD,
		Incl:            unit.Incl,
		Azim:            unit.Azim,
		SubSea:          unit.SubSea,
		TVD:             unit.TVD,
		LocalNCoord:     unit.LocalNCoord,
		LocalECoord:     unit.LocalECoord,
		GlobalNCoord:    unit.GlobalNCoord,
		GlobalECoord:    unit.GlobalECoord,
		Dogleg:          unit.Dogleg,
		VerticalSection: unit.VerticalSection,
	}
}

// toDomainTrajectoryUnit maps the GORM TrajectoryUnit model to the domain TrajectoryUnit entity.
func toDomainTrajectoryUnit(unitModel *models.TrajectoryUnit) *entities.TrajectoryUnit {
	return &entities.TrajectoryUnit{
		ID:              unitModel.ID.String(),
		MD:              unitModel.MD,
		Incl:            unitModel.Incl,
		Azim:            unitModel.Azim,
		SubSea:          unitModel.SubSea,
		TVD:             unitModel.TVD,
		LocalNCoord:     unitModel.LocalNCoord,
		LocalECoord:     unitModel.LocalECoord,
		GlobalNCoord:    unitModel.GlobalNCoord,
		GlobalECoord:    unitModel.GlobalECoord,
		Dogleg:          unitModel.Dogleg,
		VerticalSection: unitModel.VerticalSection,
		CreatedAt:       unitModel.CreatedAt,
	}
}

// toDomainCase maps the GORM Case model to the domain Case entity.
func toDomainCase(caseModel *models.Case) *entities.Case {
	newCase := entities.Case{
		ID:                   caseModel.ID.String(),
		CaseName:             caseModel.CaseName,
		CaseDescription:      caseModel.CaseDescription,
		DrillDepth:           caseModel.DrillDepth,
		PipeSize:             caseModel.PipeSize,
		CreatedAt:            caseModel.CreatedAt,
		Holes:                make([]*entities.Hole, len(caseModel.Holes)),
		Fluids:               make([]*entities.Fluid, len(caseModel.Fluids)),
		Strings:              make([]*entities.String, len(caseModel.Strings)),
		PorePressures:        make([]*entities.PorePressure, len(caseModel.PorePressures)),
		PressureDataProfiles: make([]*entities.PressureDataProfile, len(caseModel.PressureDataProfiles)),
		FractureGradients:    make([]*entities.FractureGradient, len(caseModel.FractureGradients)),
		Rigs:                 make([]*entities.Rig, len(caseModel.Rigs)),
	}

	for _, hole := range caseModel.Holes {
		domainHole := toDomainHole(&hole)
		newCase.Holes = append(newCase.Holes, domainHole)
	}

	for _, fluid := range caseModel.Fluids {
		domainFluid := toDomainFluid(&fluid)
		newCase.Fluids = append(newCase.Fluids, domainFluid)
	}

	for _, str := range caseModel.Strings {
		domainString := toDomainString(&str)
		newCase.Strings = append(newCase.Strings, domainString)
	}

	for _, pp := range caseModel.PorePressures {
		domainPP := toDomainPorePressure(&pp)
		newCase.PorePressures = append(newCase.PorePressures, domainPP)
	}

	for _, pdp := range caseModel.PressureDataProfiles {
		domainPDP := toDomainPressureDataProfile(&pdp)
		newCase.PressureDataProfiles = append(newCase.PressureDataProfiles, domainPDP)
	}

	for _, fg := range caseModel.FractureGradients {
		domainFG := toDomainFractureGradient(&fg)
		newCase.FractureGradients = append(newCase.FractureGradients, domainFG)
	}

	for _, rig := range caseModel.Rigs {
		domainRig := toDomainRig(&rig)
		newCase.Rigs = append(newCase.Rigs, domainRig)
	}

	return &newCase
}

// toGormCase maps the domain Case entity to the GORM Case model.
func toGormCase(caseEntity *entities.Case) *models.Case {
	caseID, err := validateGormId(caseEntity.ID)
	if err != nil {
		return nil
	}

	newCase := &models.Case{
		ID:              caseID,
		CaseName:        caseEntity.CaseName,
		CaseDescription: caseEntity.CaseDescription,
		DrillDepth:      caseEntity.DrillDepth,
		PipeSize:        caseEntity.PipeSize,
		Holes:           make([]models.Hole, len(caseEntity.Holes)),
	}

	for _, hole := range caseEntity.Holes {
		gormHole := toGormHole(hole)
		newCase.Holes = append(newCase.Holes, *gormHole)
	}

	return newCase
}

// toDomainHole maps the GORM Hole model to the domain Hole entity.
func toDomainHole(holeModel *models.Hole) *entities.Hole {
	newHole := &entities.Hole{
		ID:                        holeModel.ID.String(),
		CreatedAt:                 holeModel.CreatedAt,
		OpenHoleMDTop:             holeModel.OpenHoleMDTop,
		OpenHoleMDBase:            holeModel.OpenHoleMDBase,
		OpenHoleLength:            holeModel.OpenHoleLength,
		EffectiveDiameter:         holeModel.EffectiveDiameter,
		FrictionFactorOpenHole:    holeModel.FrictionFactorOpenHole,
		LinearCapacityOpenHole:    holeModel.LinearCapacityOpenHole,
		VolumeExcess:              holeModel.VolumeExcess,
		DescriptionOpenHole:       holeModel.DescriptionOpenHole,
		TrippingInCasing:          holeModel.TrippingInCasing,
		TrippingOutCasing:         holeModel.TrippingOutCasing,
		RotatingOnBottomCasing:    holeModel.RotatingOnBottomCasing,
		SlideDrillingCasing:       holeModel.SlideDrillingCasing,
		BackReamingCasing:         holeModel.BackReamingCasing,
		RotatingOffBottomCasing:   holeModel.RotatingOffBottomCasing,
		TrippingInOpenHole:        holeModel.TrippingInOpenHole,
		TrippingOutOpenHole:       holeModel.TrippingOutOpenHole,
		RotatingOnBottomOpenHole:  holeModel.RotatingOnBottomOpenHole,
		SlideDrillingOpenHole:     holeModel.SlideDrillingOpenHole,
		BackReamingOpenHole:       holeModel.BackReamingOpenHole,
		RotatingOffBottomOpenHole: holeModel.RotatingOffBottomOpenHole,
		Caisings:                  make([]*entities.Caising, len(holeModel.Caisings)),
	}

	for _, caising := range holeModel.Caisings {
		domainCaising := toDomainCaising(&caising)
		newHole.Caisings = append(newHole.Caisings, domainCaising)
	}

	return newHole
}

// toGormHole maps the domain Hole entity to the GORM Hole model.
func toGormHole(hole *entities.Hole) *models.Hole {
	holeID, err := validateGormId(hole.ID)
	if err != nil {
		return nil
	}

	newHole := &models.Hole{
		ID:                        holeID,
		OpenHoleMDTop:             hole.OpenHoleMDTop,
		OpenHoleMDBase:            hole.OpenHoleMDBase,
		OpenHoleLength:            hole.OpenHoleLength,
		EffectiveDiameter:         hole.EffectiveDiameter,
		FrictionFactorOpenHole:    hole.FrictionFactorOpenHole,
		LinearCapacityOpenHole:    hole.LinearCapacityOpenHole,
		VolumeExcess:              hole.VolumeExcess,
		DescriptionOpenHole:       hole.DescriptionOpenHole,
		TrippingInCasing:          hole.TrippingInCasing,
		TrippingOutCasing:         hole.TrippingOutCasing,
		RotatingOnBottomCasing:    hole.RotatingOnBottomCasing,
		SlideDrillingCasing:       hole.SlideDrillingCasing,
		BackReamingCasing:         hole.BackReamingCasing,
		RotatingOffBottomCasing:   hole.RotatingOffBottomCasing,
		TrippingInOpenHole:        hole.TrippingInOpenHole,
		TrippingOutOpenHole:       hole.TrippingOutOpenHole,
		RotatingOnBottomOpenHole:  hole.RotatingOnBottomOpenHole,
		SlideDrillingOpenHole:     hole.SlideDrillingOpenHole,
		BackReamingOpenHole:       hole.BackReamingOpenHole,
		RotatingOffBottomOpenHole: hole.RotatingOffBottomOpenHole,
		Caisings:                  make([]models.Caising, len(hole.Caisings)),
	}

	for _, caising := range hole.Caisings {
		gormCaising := toGormCaising(caising)
		newHole.Caisings = append(newHole.Caisings, *gormCaising)
	}

	return newHole
}

// toDomainCaising converts a gorm caising to a domain caising.
func toDomainCaising(casingModel *models.Caising) *entities.Caising {
	return &entities.Caising{
		ID:                    casingModel.ID.String(),
		MDTop:                 casingModel.MDTop,
		MDBase:                casingModel.MDBase,
		Length:                casingModel.Length,
		ShoeMD:                casingModel.ShoeMD,
		OD:                    casingModel.OD,
		VD:                    casingModel.VD,
		DriftID:               casingModel.DriftID,
		EffectiveHoleDiameter: casingModel.EffectiveHoleDiameter,
		Weight:                casingModel.Weight,
		Grade:                 casingModel.Grade,
		MinYieldStrength:      casingModel.MinYieldStrength,
		BurstRating:           casingModel.BurstRating,
		CollapseRating:        casingModel.CollapseRating,
		FrictionFactorCaising: casingModel.FrictionFactorCaising,
		LinearCapacityCaising: casingModel.LinearCapacityCaising,
		DescriptionCaising:    casingModel.DescriptionCaising,
		ManufacturerCaising:   casingModel.ManufacturerCaising,
		ModelCaising:          casingModel.ModelCaising,
	}
}

// toGormCaising converts a domain caising to a gorm caising.
func toGormCaising(caising *entities.Caising) *models.Caising {
	caisingID, err := validateGormId(caising.ID)
	if err != nil {
		return nil
	}

	return &models.Caising{
		ID:                    caisingID,
		MDTop:                 caising.MDTop,
		MDBase:                caising.MDBase,
		Length:                caising.Length,
		ShoeMD:                caising.ShoeMD,
		OD:                    caising.OD,
		VD:                    caising.VD,
		DriftID:               caising.DriftID,
		EffectiveHoleDiameter: caising.EffectiveHoleDiameter,
		Weight:                caising.Weight,
		Grade:                 caising.Grade,
		MinYieldStrength:      caising.MinYieldStrength,
		BurstRating:           caising.BurstRating,
		CollapseRating:        caising.CollapseRating,
		FrictionFactorCaising: caising.FrictionFactorCaising,
		LinearCapacityCaising: caising.LinearCapacityCaising,
		DescriptionCaising:    caising.DescriptionCaising,
		ManufacturerCaising:   caising.ManufacturerCaising,
		ModelCaising:          caising.ModelCaising,
	}
}

// toDomainFluid converts a gorm fluid to a domain fluid.
func toDomainFluid(fluidModel *models.Fluid) *entities.Fluid {
	return &entities.Fluid{
		ID:            fluidModel.ID.String(),
		Name:          fluidModel.Name,
		Description:   fluidModel.Description,
		Density:       fluidModel.Density,
		BaseFluid:     toDomainFluidType(&fluidModel.BaseFluid),
		FluidBaseType: toDomainFluidType(&fluidModel.FluidBaseType),
	}
}

// toDomainFluidType converts a gorm fluid type to a domain fluid type.
func toDomainFluidType(fluidTypeModel *models.FluidType) *entities.FluidType {
	return &entities.FluidType{
		ID:   fluidTypeModel.ID.String(),
		Name: fluidTypeModel.Name,
	}
}

// toGormFluid converts a domain fluid to a gorm fluid.
func toGormFluid(fluid *entities.Fluid) *models.Fluid {
	fluidID, err := validateGormId(fluid.ID)
	if err != nil {
		return nil
	}
	return &models.Fluid{
		ID:            fluidID,
		Name:          fluid.Name,
		Description:   fluid.Description,
		Density:       fluid.Density,
		BaseFluid:     *toGormFluidType(fluid.BaseFluid),
		FluidBaseType: *toGormFluidType(fluid.FluidBaseType),
	}
}

// toGormFluidType converts a domain fluid type to a gorm fluid type.
func toGormFluidType(fluidType *entities.FluidType) *models.FluidType {
	fluidTypeID, err := validateGormId(fluidType.ID)
	if err != nil {
		return nil
	}

	return &models.FluidType{
		ID:   fluidTypeID,
		Name: fluidType.Name,
	}
}

// toDomainRig maps the GORM Rig model to the domain Rig entity.
func toDomainRig(rigModel *models.Rig) *entities.Rig {
	return &entities.Rig{
		ID:                                rigModel.ID.String(),
		CreatedAt:                         rigModel.CreatedAt,
		BlockRating:                       rigModel.BlockRating,
		TorqueRating:                      rigModel.TorqueRating,
		RatedWorkingPressure:              rigModel.RatedWorkingPressure,
		BopPressureRating:                 rigModel.BopPressureRating,
		SurfacePressureLoss:               rigModel.SurfacePressureLoss,
		StandpipeLength:                   rigModel.StandpipeLength,
		StandpipeInternalDiameter:         rigModel.StandpipeInternalDiameter,
		HoseLength:                        rigModel.HoseLength,
		HoseInternalDiameter:              rigModel.HoseInternalDiameter,
		SwivelLength:                      rigModel.SwivelLength,
		SwivelInternalDiameter:            rigModel.SwivelInternalDiameter,
		KellyLength:                       rigModel.KellyLength,
		KellyInternalDiameter:             rigModel.KellyInternalDiameter,
		PumpDischargeLineLength:           rigModel.PumpDischargeLineLength,
		PumpDischargeLineInternalDiameter: rigModel.PumpDischargeLineInternalDiameter,
		TopDriveStackupLength:             rigModel.TopDriveStackupLength,
		TopDriveStackupInternalDiameter:   rigModel.TopDriveStackupInternalDiameter,
	}
}

// toGormRig maps the domain Rig entity to the GORM Rig model.
func toGormRig(rig *entities.Rig) *models.Rig {
	rigID, err := validateGormId(rig.ID)
	if err != nil {
		return nil
	}

	return &models.Rig{
		ID:                                rigID,
		BlockRating:                       rig.BlockRating,
		TorqueRating:                      rig.TorqueRating,
		RatedWorkingPressure:              rig.RatedWorkingPressure,
		BopPressureRating:                 rig.BopPressureRating,
		SurfacePressureLoss:               rig.SurfacePressureLoss,
		StandpipeLength:                   rig.StandpipeLength,
		StandpipeInternalDiameter:         rig.StandpipeInternalDiameter,
		HoseLength:                        rig.HoseLength,
		HoseInternalDiameter:              rig.HoseInternalDiameter,
		SwivelLength:                      rig.SwivelLength,
		SwivelInternalDiameter:            rig.SwivelInternalDiameter,
		KellyLength:                       rig.KellyLength,
		KellyInternalDiameter:             rig.KellyInternalDiameter,
		PumpDischargeLineLength:           rig.PumpDischargeLineLength,
		PumpDischargeLineInternalDiameter: rig.PumpDischargeLineInternalDiameter,
		TopDriveStackupLength:             rig.TopDriveStackupLength,
		TopDriveStackupInternalDiameter:   rig.TopDriveStackupInternalDiameter,
	}
}

// toDomainPorePressure maps the GORM PorePressure model to the domain PorePressure entity.
func toDomainPorePressure(ppModel *models.PorePressure) *entities.PorePressure {
	return &entities.PorePressure{
		ID:        ppModel.ID.String(),
		TVD:       ppModel.TVD,
		Pressure:  ppModel.Pressure,
		EMW:       ppModel.EMW,
		CreatedAt: ppModel.CreatedAt,
	}
}

// toGormPorePressure maps the domain PorePressure entity to the GORM PorePressure model.
func toGormPorePressure(pp *entities.PorePressure) *models.PorePressure {
	pressureID, err := validateGormId(pp.ID)
	if err != nil {
		return nil
	}

	return &models.PorePressure{
		ID:       pressureID,
		TVD:      pp.TVD,
		Pressure: pp.Pressure,
		EMW:      pp.EMW,
	}
}

// toDomainPressureDataProfile maps the GORM PressureDataProfile model to the domain PressureDataProfile entity.
func toDomainPressureDataProfile(profileModel *models.PressureDataProfile) *entities.PressureDataProfile {
	return &entities.PressureDataProfile{
		ID:        profileModel.ID.String(),
		TVD:       profileModel.TVD,
		Pressure:  profileModel.Pressure,
		EMW:       profileModel.EMW,
		CreatedAt: profileModel.CreatedAt,
	}
}

// toGormPressureDataProfile maps the domain PressureDataProfile entity to the GORM PressureDataProfile model.
func toGormPressureDataProfile(profile *entities.PressureDataProfile) *models.PressureDataProfile {
	ppDataProfileID, err := validateGormId(profile.ID)
	if err != nil {
		return nil
	}

	return &models.PressureDataProfile{
		ID:       ppDataProfileID,
		TVD:      profile.TVD,
		Pressure: profile.Pressure,
		EMW:      profile.EMW,
	}
}

// toDomainFractureGradient maps the GORM FractureGradient model to the domain FractureGradient entity.
func toDomainFractureGradient(model *models.FractureGradient) *entities.FractureGradient {
	return &entities.FractureGradient{
		ID:                   model.ID.String(),
		TemperatureAtSurface: model.TemperatureAtSurface,
		TemperatureAtWellTVD: model.TemperatureAtWellTVD,
		TemperatureGradient:  model.TemperatureGradient,
		WellTVD:              model.WellTVD,
		CreatedAt:            model.CreatedAt,
	}
}

// toGormFractureGradient maps the domain FractureGradient entity to the GORM FractureGradient model.
func toGormFractureGradient(entity *entities.FractureGradient) *models.FractureGradient {
	gradientID, err := validateGormId(entity.ID)
	if err != nil {
		return nil
	}

	return &models.FractureGradient{
		ID:                   uuid.MustParse(entity.ID),
		CaseID:               gradientID,
		TemperatureAtSurface: entity.TemperatureAtSurface,
		TemperatureAtWellTVD: entity.TemperatureAtWellTVD,
		TemperatureGradient:  entity.TemperatureGradient,
		WellTVD:              entity.WellTVD,
		CreatedAt:            entity.CreatedAt,
	}
}

// toDomainString maps the GORM String model to the domain String entity.
func toDomainString(gormString *models.String) *entities.String {
	sections := make([]*entities.Section, len(gormString.Sections))
	for i, sec := range gormString.Sections {
		sections[i] = toDomainSection(&sec)
	}

	return &entities.String{
		ID:        gormString.ID.String(),
		Name:      gormString.Name,
		Depth:     gormString.Depth,
		CreatedAt: gormString.CreatedAt,
		Sections:  sections,
	}
}

// toGormString maps the domain String entity to the GORM String model.
func toGormString(stringEntity *entities.String) *models.String {
	caseUUID, err := validateGormId(stringEntity.ID)
	if err != nil {
		return nil
	}
	sections := make([]models.Section, len(stringEntity.Sections))
	for i, sec := range stringEntity.Sections {
		sections[i] = *toGormSection(sec)
	}

	return &models.String{
		ID:        caseUUID,
		Name:      stringEntity.Name,
		Depth:     stringEntity.Depth,
		CreatedAt: stringEntity.CreatedAt,
		Sections:  sections,
	}
}

// toDomainSection maps the GORM Section model to the domain Section entity.
func toDomainSection(gormSection *models.Section) *entities.Section {
	return &entities.Section{
		ID:                  gormSection.ID.String(),
		Description:         gormSection.Description,
		Manufacturer:        gormSection.Manufacturer,
		Type:                gormSection.Type,
		BodyOD:              gormSection.BodyOD,
		BodyID:              gormSection.BodyID,
		AvgJointLength:      gormSection.AvgJointLength,
		StabilizerLength:    gormSection.StabilizerLength,
		StabilizerOD:        gormSection.StabilizerOD,
		StabilizerID:        gormSection.StabilizerID,
		Weight:              gormSection.Weight,
		Material:            gormSection.Material,
		Grade:               gormSection.Grade,
		Class:               gormSection.Class,
		FrictionCoefficient: gormSection.FrictionCoefficient,
		MinYieldStrength:    gormSection.MinYieldStrength,
		CreatedAt:           gormSection.CreatedAt,
	}
}

// toGormSection maps the domain Section entity to the GORM Section model.
func toGormSection(sectionEntity *entities.Section) *models.Section {
	sectionUUID, err := validateGormId(sectionEntity.ID)
	if err != nil {
		return nil
	}
	return &models.Section{
		ID:                  sectionUUID,
		Description:         sectionEntity.Description,
		Manufacturer:        sectionEntity.Manufacturer,
		Type:                sectionEntity.Type,
		BodyOD:              sectionEntity.BodyOD,
		BodyID:              sectionEntity.BodyID,
		AvgJointLength:      sectionEntity.AvgJointLength,
		StabilizerLength:    sectionEntity.StabilizerLength,
		StabilizerOD:        sectionEntity.StabilizerOD,
		StabilizerID:        sectionEntity.StabilizerID,
		Weight:              sectionEntity.Weight,
		Material:            sectionEntity.Material,
		Grade:               sectionEntity.Grade,
		Class:               sectionEntity.Class,
		FrictionCoefficient: sectionEntity.FrictionCoefficient,
		MinYieldStrength:    sectionEntity.MinYieldStrength,
		CreatedAt:           sectionEntity.CreatedAt,
	}
}

// validateGormId validates the GORM ID.
func validateGormId(id string) (uuid.UUID, error) {
	if id == "" {
		return uuid.Nil, nil
	}
	newId, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, types.ErrInvalidUUID
	}
	return newId, nil
}
