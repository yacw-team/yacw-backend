package utils

func DatabaseCheck() bool {

	table_name := map[string][]string{
		"chatconversation": []string{"id", "title", "uid", "modelid"},
		"chatmessage":      []string{"id", "content", "chatid", "show", "actor"},
		"game":             []string{"gameId", "description", "systemprompt", "name"},
		"gamemessage":      []string{"uid", "story", "chocie", "round", "gameId"},
		"personality":      []string{"id", "personalityname", "description", "prompts", "uid", "designer"},
		"prompt":           []string{"id", "promptname", "description", "prompts", "uid", "designer", "icon", "type"},
	}
	migrator := DB.Migrator()

	for key, value := range table_name {
		if !migrator.HasTable(key) {
			return false
		}
		for _, column := range value {
			if !migrator.HasColumn(key, column) {
				return false
			}
		}
	}

	return true
}
