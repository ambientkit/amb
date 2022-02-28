package internal

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/c-bata/go-prompt"
)

const (
	paramCreatePluginFolder     = "-folder"
	paramCreatePluginTemplate   = "-template"
	defaultCreatePluginTemplate = "mvp"
	defaultCreatePluginFolder   = "plugin/myplugin"
)

// assets represents the embedded templates.
//go:embed testdata/templates/*.go
var assets embed.FS

// CmdCreatePlugin represents a command object.
type CmdCreatePlugin struct {
	CmdBase
}

// Command returns the initial command.
func (c *CmdCreatePlugin) Command() string {
	return "createplugin"
}

// Suggestion returns the suggestion for the initial command.
func (c *CmdCreatePlugin) Suggestion() prompt.Suggest {
	return prompt.Suggest{Text: c.Command(), Description: "Create Ambient plugin..."}
}

// ArgumentSuggestions returns a smart suggestion group that includes validation.
func (c *CmdCreatePlugin) ArgumentSuggestions() SmartSuggestGroup {
	return SmartSuggestGroup{
		{Suggest: prompt.Suggest{Text: paramCreatePluginFolder, Description: fmt.Sprintf("Folder for the plugin (default: %v)", defaultCreatePluginFolder)}, Required: false},
		{Suggest: prompt.Suggest{Text: paramCreatePluginTemplate, Description: fmt.Sprintf("Plugin template (default: %v | options: full)", defaultCreatePluginTemplate)}, Required: false},
	}
}

// Executer executes the command.
func (c *CmdCreatePlugin) Executer(args []string) {
	// Get folder name.
	folderName, err := c.Param(args, paramCreatePluginFolder)
	if err != nil {
		folderName = defaultCreatePluginFolder
	}

	// Determine if folder already exists.
	if _, err := os.Stat(folderName); !os.IsNotExist(err) {
		log.Error("amb: folder already exists: %v", folderName)
		return
	}

	// Get template name.
	templateName, err := c.Param(args, paramCreatePluginTemplate)
	if err != nil {
		templateName = defaultCreatePluginTemplate
	}

	// Make plugin folder.
	err = os.MkdirAll(folderName, 0755)
	if err != nil {
		log.Error("amb: couldn't create folder (%v): %v", folderName, err.Error())
		return
	}

	// Use the root folder.
	fsys, err := fs.Sub(assets, ".")
	if err != nil {
		log.Error("amb: couldn't load assets: %v", err.Error())
		return
	}

	// Open the file.
	templatePath := path.Join("testdata", "templates", templateName+".go")
	f, err := fsys.Open(templatePath)
	if err != nil {
		log.Error("amb: couldn't load assets: %v", err.Error())
		return
	}
	defer f.Close()

	// Get the contents.
	pluginTemplate, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error("amb: couldn't find template (%v): %v", templatePath, err.Error())
		return
	}

	// Replace the variables.
	out := strings.ReplaceAll(string(pluginTemplate), "var1", path.Base(folderName))

	// Write out the plugin file.
	finalOutput := path.Join(folderName, path.Base(folderName)+".go")
	err = os.WriteFile(finalOutput, []byte(out), 0644)
	if err != nil {
		log.Error("amb: couldn't write template (%v): %v", finalOutput, err.Error())
		return
	}

	log.Info("amb: created plugin successfully: %v", finalOutput)
}

// Completer returns a list of suggestions based on the user input.
func (c *CmdCreatePlugin) Completer(d prompt.Document, args []string) []prompt.Suggest {
	// Don't show any suggestions if type types: --parameter SPACE
	prevCursor := d.GetWordBeforeCursorWithSpace()
	if strings.HasPrefix(prevCursor, "-") && strings.HasSuffix(prevCursor, " ") {
		return nil
	}

	// Remove duplicates from autocomplete if they've already been typed in.
	list := filterAlreadyUsed(c.ArgumentSuggestions().ToSuggest(), args)

	// Only show autocomplete when the word matches.
	list = prompt.FilterHasPrefix(list, d.GetWordBeforeCursor(), true)

	return list
}
