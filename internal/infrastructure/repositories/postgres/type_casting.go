package postgres

import (
	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
)

// toGormOrganization converts a domain organization to a gorm organization.
func toGormOrganization(organization *entities.Organization) *models.Organization {
	org := &models.Organization{
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
	comp := &models.Company{
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
	newSite := &models.Site{
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
	newWell := &models.Well{
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
	newWellbore := &models.Wellbore{
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
	newField := &models.Field{
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
	newDesign := &models.Design{
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
		Name:        trajectory.Name,
		Description: trajectory.Description,
		Headers:     headers,
		Units:       units,
		Cases:       cases,
	}
}

// toDomainTrajectoryHeader maps the GORM TrajectoryHeader model to the domain TrajectoryHeader entity.
func toGormTrajectoryHeader(header *entities.TrajectoryHeader) *models.TrajectoryHeader {
	var headerID uuid.UUID
	if header.ID != "" {
		headerID, _ = uuid.Parse(header.ID)
	} else {
		headerID = uuid.Nil // Allow DB to generate if not present
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
	var unitID uuid.UUID
	if unit.ID != "" {
		unitID, _ = uuid.Parse(unit.ID)
	} else {
		unitID = uuid.Nil // Allow DB to generate if not present
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
		ID:              caseModel.ID.String(),
		CaseName:        caseModel.CaseName,
		CaseDescription: caseModel.CaseDescription,
		DrillDepth:      caseModel.DrillDepth,
		PipeSize:        caseModel.PipeSize,
		CreatedAt:       caseModel.CreatedAt,
	}

	for _, hole := range caseModel.Holes {
		domainHole := toDomainHole(&hole)
		newCase.Holes = append(newCase.Holes, &domainHole)
	}

	return &newCase
}

// toGormCase maps the domain Case entity to the GORM Case model.
func toGormCase(caseEntity *entities.Case) *models.Case {
	newCase := &models.Case{
		CaseName:        caseEntity.CaseName,
		CaseDescription: caseEntity.CaseDescription,
		DrillDepth:      caseEntity.DrillDepth,
		PipeSize:        caseEntity.PipeSize,
	}

	for _, hole := range caseEntity.Holes {
		gormHole := toGormHole(hole)
		newCase.Holes = append(newCase.Holes, *gormHole)
	}

	return newCase
}

// toDomainHole maps the GORM Hole model to the domain Hole entity.
func toDomainHole(holeModel *models.Hole) entities.Hole {
	return entities.Hole{
		ID:                        holeModel.ID.String(),
		CaseID:                    holeModel.CaseID.String(),
		CreatedAt:                 holeModel.CreatedAt,
		MDTop:                     holeModel.MDTop,
		MDBase:                    holeModel.MDBase,
		Length:                    holeModel.Length,
		ShoeMD:                    holeModel.ShoeMD,
		OD:                        holeModel.OD,
		CaisingInternalDiameter:   holeModel.CaisingInternalDiameter,
		DriftInternalDiameter:     holeModel.DriftInternalDiameter,
		EffectiveHoleDiameter:     holeModel.EffectiveHoleDiameter,
		Weight:                    holeModel.Weight,
		Grade:                     holeModel.Grade,
		MinYieldStrength:          holeModel.MinYieldStrength,
		BurstRating:               holeModel.BurstRating,
		CollapseRating:            holeModel.CollapseRating,
		FrictionFactorCasing:      holeModel.FrictionFactorCasing,
		LinearCapacityCasing:      holeModel.LinearCapacityCasing,
		DescriptionCasing:         holeModel.DescriptionCasing,
		ManufacturerCasing:        holeModel.ManufacturerCasing,
		ModelCasing:               holeModel.ModelCasing,
		OpenHoleMDTop:             holeModel.OpenHoleMDTop,
		OpenHoleMDBase:            holeModel.OpenHoleMDBase,
		OpenHoleLength:            holeModel.OpenHoleLength,
		OpenHoleInternalDiameter:  holeModel.OpenHoleInternalDiameter,
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
	}
}

// toGormHole maps the domain Hole entity to the GORM Hole model.
func toGormHole(hole *entities.Hole) *models.Hole {
	return &models.Hole{
		MDTop:                     hole.MDTop,
		MDBase:                    hole.MDBase,
		Length:                    hole.Length,
		ShoeMD:                    hole.ShoeMD,
		OD:                        hole.OD,
		CaisingInternalDiameter:   hole.CaisingInternalDiameter,
		DriftInternalDiameter:     hole.DriftInternalDiameter,
		EffectiveHoleDiameter:     hole.EffectiveHoleDiameter,
		Weight:                    hole.Weight,
		Grade:                     hole.Grade,
		MinYieldStrength:          hole.MinYieldStrength,
		BurstRating:               hole.BurstRating,
		CollapseRating:            hole.CollapseRating,
		FrictionFactorCasing:      hole.FrictionFactorCasing,
		LinearCapacityCasing:      hole.LinearCapacityCasing,
		DescriptionCasing:         hole.DescriptionCasing,
		ManufacturerCasing:        hole.ManufacturerCasing,
		ModelCasing:               hole.ModelCasing,
		OpenHoleMDTop:             hole.OpenHoleMDTop,
		OpenHoleMDBase:            hole.OpenHoleMDBase,
		OpenHoleLength:            hole.OpenHoleLength,
		OpenHoleInternalDiameter:  hole.OpenHoleInternalDiameter,
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
	}
}
