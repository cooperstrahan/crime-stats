package domain

type CrimeDecoder struct {
	DR_NO        string            `csv:"DR_NO"`
	DateRptd     string            `csv:"Date Rptd"`
	DATEOCC      string            `csv:"DATE OCC"`
	TIMEOCC      string            `csv:"TIME OCC"`
	AREA         string            `csv:"AREA"`
	AREANAME     string            `csv:"AREA NAME"`
	RptDistNo    string            `csv:"Rpt Dist No"`
	PartOneTwo   string            `csv:"Part 1-2"`
	CrmCd        string            `csv:"Crm Cd"`
	CrmCdDesc    string            `csv:"Crm Cd Desc"`
	Mocodes      string            `csv:"Mocodes"`
	VictAge      string            `csv:"Vict Age"`
	VictSex      string            `csv:"Vict Sex"`
	VictDescent  string            `csv:"Vict Descent"`
	PremisCd     string            `csv:"Premis Cd"`
	PremisDesc   string            `csv:"Premis Desc"`
	WeaponUsedCd string            `csv:"Weapon Used Cd"`
	WeaponDesc   string            `csv:"Weapon Desc"`
	Status       string            `csv:"Status"`
	StatusDesc   string            `csv:"Status Desc"`
	CrmCd1       string            `csv:"Crm Cd 1"`
	CrmCd2       string            `csv:"Crm Cd 2"`
	CrmCd3       string            `csv:"Crm Cd 3"`
	CrmCd4       string            `csv:"Crm Cd 4"`
	LOCATION     string            `csv:"LOCATION"`
	CrossStreet  string            `csv:"Cross Street"`
	LAT          string            `csv:"LAT"`
	LON          string            `csv:"LON"`
	OtherData    map[string]string `csv:"-"`
}
