package youtrack

type serverIssue struct {
	Summary      string        `json:"summary"`
	Description  *string       `json:"description"`
	ID           string        `json:"idReadable"`
	CustomFields []customField `json:"customFields"`
}

func (serverIssue *serverIssue) getDescription() string {
	if serverIssue.Description != nil {
		return *serverIssue.Description
	}

	return ""
}

func (serverIssue *serverIssue) isResolved() bool {
	for _, field := range serverIssue.CustomFields {
		if field.ProjectCustomField.Field.Name == "Stage" {
			if data, ok := field.Value.(map[string]interface{}); ok {
				isResolvedField := data["isResolved"]

				if isResolved, ok := isResolvedField.(bool); ok {
					return isResolved
				}
			}
		}
	}

	return false
}

func (serverIssue *serverIssue) getPriority() string {
	for _, field := range serverIssue.CustomFields {
		if field.ProjectCustomField.Field.Name == "Priority" {
			if data, ok := field.Value.(map[string]interface{}); ok {
				localizedNameField := data["localizedName"]

				if localizedName, ok := localizedNameField.(string); ok {
					return localizedName
				}

				nameField := data["name"]
				if name, ok := nameField.(string); ok {
					return name
				}
			}
		}
	}

	return ""
}

func (serverIssue *serverIssue) getStage() string {
	for _, field := range serverIssue.CustomFields {
		if field.ProjectCustomField.Field.Name == "Stage" {
			if data, ok := field.Value.(map[string]interface{}); ok {
				localizedNameField := data["localizedName"]

				if localizedName, ok := localizedNameField.(string); ok {
					return localizedName
				}

				nameField := data["name"]
				if name, ok := nameField.(string); ok {
					return name
				}
			}
		}
	}

	return ""
}

func (serverIssue *serverIssue) getFixVersion() string {
	for _, field := range serverIssue.CustomFields {
		if field.ProjectCustomField.Field.Name == "Fix versions" {
			if data, ok := field.Value.(map[string]interface{}); ok {
				localizedNameField := data["localizedName"]

				if localizedName, ok := localizedNameField.(string); ok {
					return localizedName
				}

				nameField := data["name"]
				if name, ok := nameField.(string); ok {
					return name
				}
			}
		}
	}

	return ""
}

func (serverIssue *serverIssue) getType() string {
	for _, field := range serverIssue.CustomFields {
		if field.ProjectCustomField.Field.Name == "Type" {
			if data, ok := field.Value.(map[string]interface{}); ok {
				localizedNameField := data["localizedName"]

				if localizedName, ok := localizedNameField.(string); ok {
					return localizedName
				}

				nameField := data["name"]
				if name, ok := nameField.(string); ok {
					return name
				}
			}
		}
	}

	return ""
}

func (serverIssue *serverIssue) getSubsystems() []string {
	subsystems := make([]string, 0)
	for _, field := range serverIssue.CustomFields {
		if field.ProjectCustomField.Field.Name == "Subsystem" {
			if arrayData, ok := field.Value.([]interface{}); ok {
				if len(arrayData) > 0 {
					if data, ok := arrayData[0].(map[string]interface{}); ok {
						localizedNameField := data["localizedName"]

						if localizedName, ok := localizedNameField.(string); ok {
							subsystems = append(subsystems, localizedName)
						} else {
							nameField := data["name"]
							if name, ok := nameField.(string); ok {
								subsystems = append(subsystems, name)
							}
						}
					}
				}
			}
		}
	}

	return subsystems
}

func (serverIssue *serverIssue) convertServerIssue() Issue {
	issue := Issue{
		Summary:     serverIssue.Summary,
		Description: serverIssue.getDescription(),
		ID:          serverIssue.ID,
		IsResolved:  serverIssue.isResolved(),
		Priority:    serverIssue.getPriority(),
		Stage:       serverIssue.getStage(),
		FixVersion:  serverIssue.getFixVersion(),
		Type:        serverIssue.getType(),
		Subsystems:  serverIssue.getSubsystems(),
	}

	return issue
}
