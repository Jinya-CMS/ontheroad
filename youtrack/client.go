package youtrack

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.jinya.de/ontheroad/database"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type Client struct {
	Project database.Project
}

func getEnumBundle(server string, query string) ([]EnumType, error) {
	reqUrl := fmt.Sprintf("%s%s?fields=%s", server, query, url.QueryEscape("values(id,description,ordinal,name)"))
	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("not found")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var enumBundle enumBundle
	err = json.Unmarshal(body, &enumBundle)
	if err != nil {
		return nil, err
	}

	items := enumBundle.Values[:]
	sort.Slice(items, func(i, j int) bool {
		return items[i].Ordinal < items[j].Ordinal
	})

	return items, nil
}

func (client *Client) GetVersions() ([]EnumType, error) {
	return getEnumBundle(client.Project.YouTrackServer, client.Project.VersionsQuery)
}

func (client *Client) GetSubsystems() ([]EnumType, error) {
	return getEnumBundle(client.Project.YouTrackServer, client.Project.SubsystemsQuery)
}

func (client *Client) GetTypes() ([]EnumType, error) {
	return getEnumBundle(client.Project.YouTrackServer, client.Project.TypesQuery)
}

func (client *Client) GetIssues(versions []string, subsystems []string, types []string, orderField string, orderDirection string) ([]Issue, error) {
	query := fmt.Sprintf("Project:%s", client.Project.Key)

	typesPrepared := make([]string, len(types))
	for idx, typ := range types {
		typesPrepared[idx] = fmt.Sprintf("type:{%s}", typ)
	}

	versionsPrepared := make([]string, len(versions))
	for idx, version := range versions {
		versionsPrepared[idx] = fmt.Sprintf("fix version:{%s}", version)
	}

	subsystemsPrepared := make([]string, len(subsystems))
	for idx, subsystem := range subsystems {
		subsystemsPrepared[idx] = fmt.Sprintf("subsystem:{%s}", subsystem)
	}

	versionQuery := strings.Join(versionsPrepared, " or ")
	typeQuery := strings.Join(typesPrepared, " or ")
	subsystemQuery := strings.Join(subsystemsPrepared, " or ")

	query = fmt.Sprintf("%s and (%s", query, strings.Join([]string{versionQuery, typeQuery, subsystemQuery}, ") and ("))

	if orderField != "" && orderDirection != "" {
		query = fmt.Sprintf("%s) and order by:{%s} %s", query, orderField, orderDirection)
	} else {
		query = fmt.Sprintf("%s)", query)
	}

	escapedQuery := url.QueryEscape(query)

	resp, err := http.Get(fmt.Sprintf(
		"%sapi/issues?query=%s&fields=idReadable,description,summary,customFields(id,projectCustomField(id,field(id,name)),value(isResolved,localizedName,name,text))",
		client.Project.YouTrackServer,
		escapedQuery))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var serverIssues []serverIssue
	err = json.Unmarshal(body, &serverIssues)
	if err != nil {
		return nil, err
	}

	issues := make([]Issue, len(serverIssues))
	for idx, serverIssue := range serverIssues {
		issues[idx] = serverIssue.convertServerIssue()
	}

	return issues, nil
}
