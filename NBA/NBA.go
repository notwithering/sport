// Package NBA provides many constants for interacting with NBA team names.
package nba

// LOCATION >>> LOCATION
const (
	Atlanta       string = "Atlanta"
	Boston        string = "Boston"
	Brooklyn      string = "Brooklyn"
	Charlotte     string = "Charlotte"
	Chicago       string = "Chicago"
	Cleveland     string = "Cleveland"
	Dallas        string = "Dallas"
	Denver        string = "Denver"
	Detroit       string = "Detroit"
	GoldenState   string = "Golden State"
	Houston       string = "Houston"
	Indiana       string = "Indiana"
	LosAngeles    string = "Los Angeles"
	Memphis       string = "Memphis"
	Miami         string = "Miami"
	Milwaukee     string = "Milwaukee"
	Minnesota     string = "Minnesota"
	NewOrleans    string = "New Orleans"
	NewYork       string = "New York"
	OklahomaCity  string = "Oklahoma City"
	Orlando       string = "Orlando"
	Philadelphia  string = "Philadelphia"
	Phoenix       string = "Phoenix"
	PortlandTrail string = "Portland Trail"
	Sacramento    string = "Sacramento"
	SanAntonio    string = "San Antonio"
	Toronto       string = "Toronto"
	Utah          string = "Utah"
	Washington    string = "Washington"
)

// TEAM NAME >>> TEAM NAME
const (
	Hawks        string = "Hawks"
	Celtics      string = "Celtics"
	Nets         string = "Nets"
	Hornets      string = "Hornets"
	Bulls        string = "Bulls"
	Cavaliers    string = "Cavaliers"
	Mavericks    string = "Mavericks"
	Nuggets      string = "Nuggets"
	Pistons      string = "Pistons"
	Warriors     string = "Warriors"
	Rockets      string = "Rockets"
	Pacers       string = "Pacers"
	Clippers     string = "Clippers"
	Lakers       string = "Lakers"
	Grizzlies    string = "Grizzlies"
	Heat         string = "Heat"
	Bucks        string = "Bucks"
	Timberwolves string = "Timberwolves"
	Pelicans     string = "Pelicans"
	Knicks       string = "Knicks"
	Thunder      string = "Thunder"
	Magic        string = "Magic"
	X76ers       string = "76ers"
	Suns         string = "Suns"
	Blazers      string = "Blazers"
	Kings        string = "Kings"
	Spurs        string = "Spurs"
	Raptors      string = "Raptors"
	Jazz         string = "Jazz"
	Wizards      string = "Wizards"
)

// FULL TEAM NAME >>> FULL TEAM NAME
const (
	AtlantaHawks          string = Atlanta + " " + Hawks
	BostonCeltics         string = Boston + " " + Celtics
	BrooklynNets          string = Brooklyn + " " + Nets
	CharlotteHornets      string = Charlotte + " " + Hornets
	ChicagoBulls          string = Chicago + " " + Bulls
	ClevelandCavaliers    string = Cleveland + " " + Cavaliers
	DallasMavericks       string = Dallas + " " + Mavericks
	DenverNuggets         string = Denver + " " + Nuggets
	DetroitPistons        string = Detroit + " " + Pistons
	GoldenStateWarriors   string = GoldenState + " " + Warriors
	HoustonRockets        string = Houston + " " + Rockets
	IndianaPacers         string = Indiana + " " + Pacers
	LosAngelesClippers    string = LosAngeles + " " + Clippers
	LosAngelesLakers      string = LosAngeles + " " + Lakers
	MemphisGrizzlies      string = Memphis + " " + Grizzlies
	MiamiHeat             string = Miami + " " + Heat
	MilwaukeeBucks        string = Milwaukee + " " + Bucks
	MinnesotaTimberwolves string = Minnesota + " " + Timberwolves
	NewOrleansPelicans    string = NewOrleans + " " + Pelicans
	NewYorkKnicks         string = NewYork + " " + Knicks
	OklahomaCityThunder   string = OklahomaCity + " " + Thunder
	OrlandoMagic          string = Orlando + " " + Magic
	Philadelphia76ers     string = Philadelphia + " " + X76ers
	PhoenixSuns           string = Phoenix + " " + Suns
	PortlandTrailBlazers  string = PortlandTrail + " " + Blazers
	SacramentoKings       string = Sacramento + " " + Kings
	SanAntonioSpurs       string = SanAntonio + " " + Spurs
	TorontoRaptors        string = Toronto + " " + Raptors
	UtahJazz              string = Utah + " " + Jazz
	WashingtonWizards     string = Washington + " " + Wizards
)

// TEAM NAME >>> LOCATION
const (
	HawksLocation        string = Atlanta
	CelticsLocation      string = Boston
	NetsLocation         string = Brooklyn
	HornetsLocation      string = Charlotte
	BullsLocation        string = Chicago
	CavaliersLocation    string = Cleveland
	MavericksLocation    string = Dallas
	NuggetsLocation      string = Denver
	PistonsLocation      string = Detroit
	WarriorsLocation     string = GoldenState
	RocketsLocation      string = Houston
	PacersLocation       string = Indiana
	ClippersLocation     string = LosAngeles
	LakersLocation       string = LosAngeles
	GrizzliesLocation    string = Memphis
	HeatLocation         string = Miami
	BucksLocation        string = Milwaukee
	TimberwolvesLocation string = Minnesota
	PelicansLocation     string = NewOrleans
	KnicksLocation       string = NewYork
	ThunderLocation      string = OklahomaCity
	MagicLocation        string = Orlando
	X76ersLocation       string = Philadelphia
	SunsLocation         string = Phoenix
	BlazersLocation      string = PortlandTrail
	KingsLocation        string = Sacramento
	SpursLocation        string = SanAntonio
	RaptorsLocation      string = Toronto
	JazzLocation         string = Utah
	WizardsLocation      string = Washington
)

// LOCATION >>> TEAM NAME
const (
	AtlantaTeam       string = Hawks
	BostonTeam        string = Celtics
	BrooklynTeam      string = Nets
	CharlotteTeam     string = Hornets
	ChicagoTeam       string = Bulls
	ClevelandTeam     string = Cavaliers
	DallasTeam        string = Mavericks
	DenverTeam        string = Nuggets
	DetroitTeam       string = Pistons
	GoldenStateTeam   string = Warriors
	HoustonTeam       string = Rockets
	IndianaTeam       string = Pacers
	LosAngelesCTeam   string = Clippers
	LosAngelesLTeam   string = Lakers
	MemphisTeam       string = Grizzlies
	MiamiTeam         string = Heat
	MilwaukeeTeam     string = Bucks
	MinnesotaTeam     string = Timberwolves
	NewOrleansTeam    string = Pelicans
	NewYorkTeam       string = Knicks
	OklahomaCityTeam  string = Thunder
	OrlandoTeam       string = Magic
	PhiladelphiaTeam  string = X76ers
	PhoenixTeam       string = Suns
	PortlandTrailTeam string = Blazers
	SacramentoTeam    string = Kings
	SanAntonioTeam    string = Spurs
	TorontoTeam       string = Raptors
	UtahTeam          string = Jazz
	WashingtonTeam    string = Wizards
)

// FULL TEAM NAME >>> LOCATION
const (
	AtlantaHawksLocation          string = Atlanta
	BostonCelticsLocation         string = Boston
	BrooklynNetsLocation          string = Brooklyn
	CharlotteHornetsLocation      string = Charlotte
	ChicagoBullsLocation          string = Chicago
	ClevelandCavaliersLocation    string = Cleveland
	DallasMavericksLocation       string = Dallas
	DenverNuggetsLocation         string = Denver
	DetroitPistonsLocation        string = Detroit
	GoldenStateWarriorsLocation   string = GoldenState
	HoustonRocketsLocation        string = Houston
	IndianaPacersLocation         string = Indiana
	LosAngelesClippersLocation    string = LosAngeles
	LosAngelesLakersLocation      string = LosAngeles
	MemphisGrizzliesLocation      string = Memphis
	MiamiHeatLocation             string = Miami
	MilwaukeeBucksLocation        string = Milwaukee
	MinnesotaTimberwolvesLocation string = Minnesota
	NewOrleansPelicansLocation    string = NewOrleans
	NewYorkKnicksLocation         string = NewYork
	OklahomaCityThunderLocation   string = OklahomaCity
	OrlandoMagicLocation          string = Orlando
	Philadelphia76ersLocation     string = Philadelphia
	PhoenixSunsLocation           string = Phoenix
	PortlandTrailBlazersLocation  string = PortlandTrail
	SacramentoKingsLocation       string = Sacramento
	SanAntonioSpursLocation       string = SanAntonio
	TorontoRaptorsLocation        string = Toronto
	UtahJazzLocation              string = Utah
	WashingtonWizardsLocation     string = Washington
)

// LOCATION >>> FULL TEAM NAME
const (
	AtlantaFullTeam       string = AtlantaHawks
	BostonFullTeam        string = BostonCeltics
	BrooklynFullTeam      string = BrooklynNets
	CharlotteFullTeam     string = CharlotteHornets
	ChicagoFullTeam       string = ChicagoBulls
	ClevelandFullTeam     string = ClevelandCavaliers
	DallasFullTeam        string = DallasMavericks
	DenverFullTeam        string = DenverNuggets
	DetroitFullTeam       string = DetroitPistons
	GoldenStateFullTeam   string = GoldenStateWarriors
	HoustonFullTeam       string = HoustonRockets
	IndianaFullTeam       string = IndianaPacers
	LosAngelesCFullTeam   string = LosAngelesClippers
	LosAngelesLFullTeam   string = LosAngelesLakers
	MemphisFullTeam       string = MemphisGrizzlies
	MiamiFullTeam         string = MiamiHeat
	MilwaukeeFullTeam     string = MilwaukeeBucks
	MinnesotaFullTeam     string = MinnesotaTimberwolves
	NewOrleansFullTeam    string = NewOrleansPelicans
	NewYorkFullTeam       string = NewYorkKnicks
	OklahomaCityFullTeam  string = OklahomaCityThunder
	OrlandoFullTeam       string = OrlandoMagic
	PhiladelphiaFullTeam  string = Philadelphia76ers
	PhoenixFullTeam       string = PhoenixSuns
	PortlandTrailFullTeam string = PortlandTrailBlazers
	SacramentoFullTeam    string = SacramentoKings
	SanAntonioFullTeam    string = SanAntonioSpurs
	TorontoFullTeam       string = TorontoRaptors
	UtahFullTeam          string = UtahJazz
	WashingtonFullTeam    string = WashingtonWizards
)
