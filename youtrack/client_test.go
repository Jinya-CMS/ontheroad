package youtrack

import (
	"go.jinya.de/ontheroad/database"
	"sort"
	"testing"
)

func TestClient_GetVersions(t *testing.T) {
	project := database.Project{
		YouTrackServer: "https://jinya.myjetbrains.com/youtrack",
		VersionsQuery:  "api/admin/customFieldSettings/bundles/version/71-2",
	}

	client := Client{project}
	versions, err := client.GetVersions()
	if err != nil {
		t.Error(err)
		return
	}

	if len(versions) == 0 {
		t.Error()
	}

	if !sort.SliceIsSorted(versions, func(i, j int) bool {
		return versions[i].Ordinal < versions[j].Ordinal
	}) {
		t.Error()
	}

	t.Logf("Versions: %#v", versions)
}

func TestClient_GetSubsystems(t *testing.T) {
	project := database.Project{
		YouTrackServer:  "https://jinya.myjetbrains.com/youtrack",
		SubsystemsQuery: "api/admin/customFieldSettings/bundles/ownedField/144-2",
	}

	client := Client{project}
	subsystems, err := client.GetSubsystems()
	if err != nil {
		t.Error(err)
		return
	}

	if len(subsystems) == 0 {
		t.Error()
	}

	if !sort.SliceIsSorted(subsystems, func(i, j int) bool {
		return subsystems[i].Ordinal < subsystems[j].Ordinal
	}) {
		t.Error()
	}

	t.Logf("Subsystems: %#v", subsystems)
}

func TestClient_GetTypes(t *testing.T) {
	project := database.Project{
		YouTrackServer: "https://jinya.myjetbrains.com/youtrack",
		TypesQuery:     "api/admin/customFieldSettings/bundles/enum/66-7",
	}

	client := Client{project}
	types, err := client.GetTypes()
	if err != nil {
		t.Error(err)
		return
	}

	if len(types) == 0 {
		t.Error()
	}

	if !sort.SliceIsSorted(types, func(i, j int) bool {
		return types[i].Ordinal < types[j].Ordinal
	}) {
		t.Error()
	}

	t.Logf("Types: %#v", types)
}

func TestClient_GetIssuesWithoutFilter(t *testing.T) {
	project := database.Project{
		YouTrackServer: "https://jinya.myjetbrains.com/youtrack",
		Key:            "JGCMS",
	}

	client := Client{project}
	issues, err := client.GetIssues(make([]string, 0), make([]string, 0), make([]string, 0), "", "")
	if err != nil {
		t.Error()
		return
	}

	if len(issues) == 0 {
		t.Error()
		return
	}

	t.Logf("Issue: %#v", issues[0])
	t.Logf("Type: %#v", issues[0].Type)
	t.Logf("Subsystems: %#v", issues[0].Subsystems)
	t.Logf("Fix Version: %#v", issues[0].FixVersion)
}

func TestClient_GetIssuesFilteredByType(t *testing.T) {
	project := database.Project{
		YouTrackServer: "https://jinya.myjetbrains.com/youtrack",
		TypesQuery:     "api/admin/customFieldSettings/bundles/enum/66-9",
		Key:            "JGCMS",
	}

	client := Client{project}
	types, err := client.GetTypes()
	if err != nil || len(types) == 0 {
		t.Error()
		return
	}

	issues, err := client.GetIssues(make([]string, 0), make([]string, 0), []string{types[0].Name}, "", "")
	if err != nil {
		t.Error()
		return
	}

	if len(issues) == 0 {
		t.Error()
		return
	}

	for _, issue := range issues {
		if issue.Type != types[0].Name {
			t.Error()
			return
		}
	}

	t.Logf("Issue: %#v", issues[0])
	t.Logf("Type: %#v", issues[0].Type)
	t.Logf("Subsystems: %#v", issues[0].Subsystems)
	t.Logf("Fix Version: %#v", issues[0].FixVersion)
}

func TestClient_GetIssuesFilteredByVersion(t *testing.T) {
	project := database.Project{
		YouTrackServer: "https://jinya.myjetbrains.com/youtrack",
		VersionsQuery:  "api/admin/customFieldSettings/bundles/version/71-2",
		Key:            "JGCMS",
	}

	client := Client{project}
	versions, err := client.GetVersions()
	if err != nil || len(versions) == 0 {
		t.Error()
		return
	}

	issues, err := client.GetIssues([]string{versions[0].Name}, make([]string, 0), make([]string, 0), "", "")
	if err != nil {
		t.Error()
		return
	}

	if len(issues) == 0 {
		t.Error()
		return
	}

	for _, issue := range issues {
		if issue.FixVersion != versions[0].Name {
			t.Error()
			return
		}
	}

	t.Logf("Issue: %#v", issues[0])
	t.Logf("Type: %#v", issues[0].Type)
	t.Logf("Subsystems: %#v", issues[0].Subsystems)
	t.Logf("Fix Version: %#v", issues[0].FixVersion)
}

func TestClient_GetIssuesFilteredBySubsystem(t *testing.T) {
	project := database.Project{
		YouTrackServer:  "https://jinya.myjetbrains.com/youtrack",
		SubsystemsQuery: "api/admin/customFieldSettings/bundles/ownedField/144-1",
		Key:             "JGCMS",
	}

	client := Client{project}
	subsystems, err := client.GetSubsystems()
	if err != nil || len(subsystems) == 0 {
		t.Error()
		return
	}

	issues, err := client.GetIssues(make([]string, 0), []string{subsystems[0].Name}, make([]string, 0), "", "")
	if err != nil {
		t.Error()
		return
	}

	if len(issues) == 0 {
		t.Error()
		return
	}

	contains := func(slice []string, value string) bool {
		for _, v := range slice {
			if v == value {
				return true
			}
		}
		return false
	}

	for _, issue := range issues {
		if !contains(issue.Subsystems, subsystems[0].Name) {
			t.Error()
			return
		}
	}

	t.Logf("Issue: %#v", issues[0])
	t.Logf("Type: %#v", issues[0].Type)
	t.Logf("Subsystems: %#v", issues[0].Subsystems)
	t.Logf("Fix Version: %#v", issues[0].FixVersion)
}
