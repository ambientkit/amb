package internal

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
)

const (
	paramCreateAppFolder     = "-folder"
	paramCreateAppTemplate   = "-template"
	defaultCreateAppTemplate = "https://github.com/ambientkit/ambient-template"
	defaultCreateAppFolder   = "ambapp"
)

// CmdCreateApp represents a command object.
type CmdCreateApp struct {
	CmdBase
}

// Command returns the initial command.
func (c *CmdCreateApp) Command() string {
	return "createapp"
}

// Suggestion returns the suggestion for the initial command.
func (c *CmdCreateApp) Suggestion() prompt.Suggest {
	return prompt.Suggest{Text: c.Command(), Description: "Create Ambient app..."}
}

// ArgumentSuggestions returns a smart suggestion group that includes validation.
func (c *CmdCreateApp) ArgumentSuggestions() SmartSuggestGroup {
	return SmartSuggestGroup{
		{Suggest: prompt.Suggest{Text: paramCreateAppFolder, Description: fmt.Sprintf("Folder to create the project (default: %v)", defaultCreateAppFolder)}, Required: false},
		{Suggest: prompt.Suggest{Text: paramCreateAppTemplate, Description: fmt.Sprintf("Template project to git clone (default: %v)", defaultCreateAppTemplate)}, Required: false},
	}
}

// Executer executes the command.
func (c *CmdCreateApp) Executer(args []string) {
	// Get folder name.
	folderName, err := c.Param(args, paramCreateAppFolder)
	if err != nil {
		folderName = defaultCreateAppFolder
	}

	// Determine if folder already exists.
	if _, err := os.Stat(folderName); !os.IsNotExist(err) {
		log.Error("folder already exists: %v", folderName)
		return
	}

	// Get template name.
	templateName, err := c.Param(args, paramCreateAppTemplate)
	if err != nil {
		templateName = defaultCreateAppTemplate
	}

	// Perform git clone on the template.
	log.Info("creating new project from template: %v", templateName)
	gitArgs := []string{"clone", "--depth=1", "--branch=main", templateName, folderName}
	cmd := exec.Command("git", gitArgs...)
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	err = cmd.Run()
	if err != nil {
		log.Error("couldn't create project (git %v): %v %v", strings.Join(gitArgs, " "), err.Error(), stdErr.String())
		return
	}

	// Remove .git folder.
	gitFolder := filepath.Join(folderName, ".git")
	err = os.RemoveAll(gitFolder)
	if err != nil {
		log.Error("couldn't remove .git folder: %v", err.Error())
	}

	// Make bin folder.
	binFolder := filepath.Join(folderName, "bin")
	err = os.Mkdir(binFolder, 0755)
	if err != nil {
		log.Error("couldn't create bin folder: %v", err.Error())
	}

	log.Info("removing folder: %v", gitFolder)
	log.Info("created project successfully in folder: %v", folderName)
}

// Completer returns a list of suggestions based on the user input.
func (c *CmdCreateApp) Completer(d prompt.Document, args []string) []prompt.Suggest {
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
