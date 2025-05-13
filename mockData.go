package main

// Deployment repräsentiert ein Deployment-Objekt
type Deployment struct {
    Owner          string         `json:"owner"`
    Name           string         `json:"name"`
    Label          string         `json:"label"`
    Description    string         `json:"description"`
    Size           string         `json:"size"`
    Flow           string         `json:"flow"`
    Template       string         `json:"template"`
    Configurations []string       `json:"configurations"`
    Envs           []DeploymentEnv `json:"envs"`
}

// DeploymentEnv repräsentiert eine Deployment-Umgebung
type DeploymentEnv struct {
    Env          string            `json:"env"`
    Versions     DeploymentVersion `json:"versions"`
    State        DeploymentState   `json:"state"`
    Apis         []Api             `json:"apis"`
}

// DeploymentVersion repräsentiert die Versionen eines Deployments
type DeploymentVersion struct {
    Template       string            `json:"template"`
    Components     []Component       `json:"components"`
    Configurations []Configuration   `json:"configurations"`
}

// DeploymentState repräsentiert den Zustand eines Deployments
type DeploymentState struct {
    URL    string `json:"url"`
    Apis   int    `json:"apis"`
    Health string `json:"health"`
    Tests  string `json:"tests"`
}

// Api repräsentiert eine API in einem Deployment
type Api struct {
    Name      string   `json:"name"`
    Kinds     []string `json:"kinds"`
    Health    string   `json:"health"`
    Tests     string   `json:"tests"`
    TestRuns  []TestRun `json:"testRuns"`
}

// Component repräsentiert eine Komponente in einem Deployment
type Component struct {
    Name    string `json:"name"`
    Version string `json:"version"`
    Main    bool   `json:"main"`
}

// Configuration repräsentiert eine Konfiguration in einem Deployment
type Configuration struct {
    Name    string `json:"name"`
    Version string `json:"version"`
    Main    bool   `json:"main"`
}

// TestRun repräsentiert einen Testlauf für eine API
type TestRun struct {
    Name   string `json:"name"`
    Result string `json:"result"`
    Report string `json:"report"`
}

// ConfigSource repräsentiert eine Konfiguration
type ConfigSource struct {
    Owner       string `json:"owner"`
    Name        string `json:"name"`
    Label       string `json:"label"`
    Description string `json:"description"`
    Content     string `json:"content"`
    Store       string `json:"store"`
    Repository  string `json:"repository"`
}

// ConfigVersion repräsentiert die Versionen einer Konfiguration
type ConfigVersion struct {
    Owner    string   `json:"owner"`
    Name     string   `json:"name"`
    Versions []string `json:"versions"`
}

// Template repräsentiert ein Template
type Template struct {
    ID          string   `json:"id"`
    Label       string   `json:"label"`
    Description string   `json:"description"`
    Components  []string `json:"components"`
}

// TemplateVersion repräsentiert die Versionen eines Templates
type TemplateVersion struct {
    ID        string                       `json:"id"`
    Versions  []string                     `json:"versions"`
    Components map[string][]TemplateComponent `json:"components"`
}

type Settings struct {
    Owner     string         `json:"owner"`
    Version   string         `json:"version"`
    Mode      string         `json:"mode"`
    Templates SettingsConfig `json:"templates"`
    Flows     SettingsConfig `json:"flows"`
}

// SettingsConfig repräsentiert die Konfiguration für Templates oder Flows
type SettingsConfig struct {
    Default string `json:"_default"`
    Single  bool   `json:"single"`
}

// Mock-Daten für Settings
var settings = Settings{
    Owner:   "ii",
    Version: "0.9.0-mock",
    Mode:    "Expert",
    Templates: SettingsConfig{
        Default: "ldproxy",
        Single:  false,
    },
    Flows: SettingsConfig{
        Default: "full",
        Single:  false,
    },
}

// TemplateComponent repräsentiert eine Komponente in einer Template-Version
type TemplateComponent struct {
    Name    string `json:"name"`
    Version string `json:"version"`
}

// Mock-Daten für Templates
var templates = []Template{
    {
        ID:          "ldproxy",
        Label:       "ldproxy",
        Description: "Only ldproxy, nothing else.",
        Components:  []string{"ldproxy"},
    },
}

var templateVersions = []TemplateVersion{
    {
        ID:       "ldproxy",
        Versions: []string{"0.8.0", "0.9.0"},
        Components: map[string][]TemplateComponent{
            "0.8.0": {
                {Name: "ldproxy", Version: "4.2.4"},
            },
            "0.9.0": {
                {Name: "ldproxy", Version: "4.3.1"},
            },
        },
    },
}

// Mock-Daten für Konfigurationen
var configs = []ConfigSource{
    {
        Owner:       "ii",
        Name:        "demo-gitea",
        Label:       "Demo from Gitea",
        Description: "Demo configuration to showcase capabilities.",
        Content:     "All",
        Store:       "gitea-ldproxy",
        Repository:  "demo",
    },
    {
        Owner:       "ii",
        Name:        "aaa-inspire-gitea",
        Label:       "aaa-inspire from Gitea",
        Description: "AAA INSPIRE Dienste.",
        Content:     "All",
        Store:       "gitea-ldproxy",
        Repository:  "aaa-inspire",
    },
}

var configVersions = []ConfigVersion{
    {
        Owner:    "ii",
        Name:     "demo-gitea",
        Versions: []string{"main", "1.0.0", "1.0.1"},
    },
    {
        Owner:    "ii",
        Name:     "aaa-inspire-gitea",
        Versions: []string{"main", "1.0.0"},
    },
}

// Mock-Daten (ersetzt später durch echte Daten)
var deployments = []Deployment{
    {
        Owner:       "ii",
        Name:        "demo",
        Label:       "Demo",
        Description: "Demo deployment to showcase capabilities.",
        Size:        "S",
        Flow:        "full",
        Template:    "ldproxy",
        Configurations: []string{
        	"demo-gitea",
        },
        Envs: []DeploymentEnv{
            {
                Env: "DEV",
                Versions: DeploymentVersion{
                    Template: "0.9.0",
                    Components: []Component{
                        {Name: "ldproxy", Version: "4.3.1", Main: true},
                    },
                    Configurations: []Configuration{
                        {Name: "demo-gitea", Version: "main", Main: true},
                    },
                },
                State: DeploymentState{
                    URL:    "https://dev.ii.geoapihub.io/demo",
                    Apis:   5,
                    Health: "Unavailable",
                    Tests:  "None",
                },
                Apis: []Api{
                    {
                        Name:   "cshapes",
                        Kinds:  []string{"Features"},
                        Health: "Unavailable",
                        Tests:  "None",
                    },
                    {
                        Name:   "daraa",
                        Kinds:  []string{"Features", "Tiles"},
                        Health: "Unavailable",
                        Tests:  "None",
                    },
                    {
                        Name:   "strassen",
                        Kinds:  []string{"Features", "Tiles"},
                        Health: "Unavailable",
                        Tests:  "None",
                    },
                    {
                        Name:   "vineyards",
                        Kinds:  []string{"Features", "Tiles"},
                        Health: "Unavailable",
                        Tests:  "None",
                    },
                    {
                        Name:   "zoomstack",
                        Kinds:  []string{"Features", "Tiles"},
                        Health: "Unavailable",
                        Tests:  "None",
                    },
                },
            },
            {
                Env: "STAGE",
                Versions: DeploymentVersion{
                    Template: "0.9.0",
                    Components: []Component{
                        {Name: "ldproxy", Version: "4.3.1", Main: true},
                    },
                    Configurations: []Configuration{
                        {Name: "demo-gitea", Version: "1.0.0", Main: true},
                    },
                },
                State: DeploymentState{
                    URL:    "https://stg.ii.geoapihub.io/demo",
                    Apis:   5,
                    Health: "Limited",
                    Tests:  "Mixed",
                },
                Apis: []Api{
                    {
                        Name:   "cshapes",
                        Kinds:  []string{"Features"},
                        Health: "Available",
                        Tests:  "Passing",
                        TestRuns: []TestRun{
                            {Name: "compare", Result: "Passing", Report: "https://example.com/report"},
                            {Name: "compliance", Result: "Passing", Report: "https://example.com/report"},
                        },
                    },
                    {
                        Name:   "daraa",
                        Kinds:  []string{"Features", "Tiles"},
                        Health: "Available",
                        Tests:  "Passing",
                        TestRuns: []TestRun{
                            {Name: "compare", Result: "Passing", Report: "https://example.com/report"},
                            {Name: "compliance", Result: "Passing", Report: "https://example.com/report"},
                        },
                    },
                },
            },
        },
    },
}