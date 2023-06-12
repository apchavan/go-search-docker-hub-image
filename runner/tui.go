package runner

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Construct & return the application instance.
func GetTuiAppLayout() *tview.Application {
	// Create pointer object of `ImageSearchDetails` struct.
	imageSearchDetails := &ImageSearchDetails{}

	// Create application object.
	application := tview.NewApplication()

	// Create flex layout as primary layout.
	flexLayout := tview.NewFlex()
	flexLayout.SetBorder(true).SetTitle(GetApplicationName())

	// 1. Create TextView to show search results in right half.
	resultsTextView := tview.NewTextView()
	resultsTextView = resultsTextView.SetText("\nSearch Results will be shown here...")
	resultsTextView = resultsTextView.SetTextAlign(tview.AlignCenter)
	resultsTextView = resultsTextView.SetChangedFunc(func() {
		application.ForceDraw().Sync()
	})

	// 2. Create form for input in left half.
	inputForm := tview.NewForm()
	inputForm.SetBorder(true).SetTitle(" Provide Search Inputs ")

	// Add InputField to get image name.
	searchTermInputField := tview.NewInputField()
	searchTermInputField = searchTermInputField.SetLabel("Search Name :")
	searchTermInputField = searchTermInputField.SetPlaceholder(" (Mandatory Input) ")
	searchTermInputField = searchTermInputField.SetFieldWidth(0)
	inputForm = inputForm.AddFormItem(searchTermInputField)

	// Add InputField to get stars count value.
	starsCountValueInputField := tview.NewInputField()
	starsCountValueInputField = starsCountValueInputField.SetLabel("Minimum Stars :")
	starsCountValueInputField = starsCountValueInputField.SetPlaceholder(" (Optional Input) ")
	starsCountValueInputField = starsCountValueInputField.SetFieldWidth(0)
	inputForm = inputForm.AddFormItem(starsCountValueInputField)

	// Add Dropdown to get input for automated builds.
	buildType_DropDown := tview.NewDropDown()
	buildType_DropDown = buildType_DropDown.SetLabel("Build Type :")
	buildType_DropDown = buildType_DropDown.SetOptions(
		[]string{
			" Default ",
			" Automated Only ",
			" Manual Only ",
		},
		func(option string, optionIndex int) {
			// Capture & set option if selected other than 'Default'.
			if optionIndex != 0 {
				imageSearchDetails.AutomatedOrManualBuildType = optionIndex
			} else {
				imageSearchDetails.AutomatedOrManualBuildType = 0
			}
		},
	)
	buildType_DropDown = buildType_DropDown.SetCurrentOption(0)
	buildType_DropDown = buildType_DropDown.SetFieldWidth(0)
	inputForm = inputForm.AddFormItem(buildType_DropDown)

	// Add Dropdown to get input for official builds.
	buildStatus_DropDown := tview.NewDropDown()
	buildStatus_DropDown = buildStatus_DropDown.SetLabel("Build Status :")
	buildStatus_DropDown = buildStatus_DropDown.SetOptions(
		[]string{
			" Default ",
			" Official Only ",
			" Community Only ",
		},
		func(option string, optionIndex int) {
			// Capture & set option if selected other than 'Default'.
			if optionIndex != 0 {
				imageSearchDetails.OfficialOrCommunityBuildStatus = optionIndex
			} else {
				imageSearchDetails.OfficialOrCommunityBuildStatus = 0
			}
		},
	)
	buildStatus_DropDown = buildStatus_DropDown.SetCurrentOption(0)
	buildStatus_DropDown = buildStatus_DropDown.SetFieldWidth(0)
	inputForm = inputForm.AddFormItem(buildStatus_DropDown)

	// Add Dropdown to get input for limit to search records.
	limitOptions := []string{" 25 - Default "}
	for limitValue := 10; limitValue <= 100; limitValue += 10 {
		limitOptions = append(limitOptions, fmt.Sprintf("%d", limitValue))
	}

	limit_DropDown := tview.NewDropDown()
	limit_DropDown = limit_DropDown.SetLabel("Search Limit :")
	limit_DropDown = limit_DropDown.SetOptions(
		limitOptions,
		func(option string, optionIndex int) {
			// Capture & set option if selected other than 'Default'.
			if optionIndex != 0 {
				limitValue, _ := strconv.Atoi(strings.TrimSpace(option))
				imageSearchDetails.SearchLimit = limitValue
			} else {
				imageSearchDetails.SearchLimit = 0
			}
		},
	)
	limit_DropDown = limit_DropDown.SetCurrentOption(0)
	limit_DropDown = limit_DropDown.SetFieldWidth(0)
	inputForm = inputForm.AddFormItem(limit_DropDown)

	// Add Button to perform search operation.
	searchResponseStringChannel := make(chan string)
	inputForm = inputForm.AddButton(
		"Search",
		func() {
			// Check if Star count input is provided.
			inputMinStars := strings.TrimSpace(starsCountValueInputField.GetText())
			if minimumStarsValue, err := strconv.Atoi(inputMinStars); err == nil {
				imageSearchDetails.MinimumStars = minimumStarsValue
			} else {
				imageSearchDetails.MinimumStars = 0
			}

			// Perform search on Docker Hub if image name is non-empty.
			searchTerm := strings.TrimSpace(searchTermInputField.GetText())
			searchTerm = strings.ToLower(searchTerm)

			if searchTerm != "" {
				imageSearchDetails.ImageName = searchTerm
				go SearchDockerHub(
					imageSearchDetails,
					searchResponseStringChannel,
				)
				resultsTextView.SetText(<-searchResponseStringChannel)
				resultsTextView = resultsTextView.ScrollToBeginning()
			} else {
				resultsTextView.SetText("\n(❌) No image name provided for search...")
			}
		},
	).SetButtonBackgroundColor(tcell.ColorPurple).SetButtonsAlign(tview.AlignRight)

	// Add button to quit from application.
	inputForm = inputForm.AddButton(
		"Quit",
		func() {
			application.Stop()
		},
	).SetButtonBackgroundColor(tcell.ColorPurple).SetButtonsAlign(tview.AlignRight)

	// Add both forms to the `flexLayout`
	flexLayout = flexLayout.AddItem(
		tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(inputForm, 0, 1, true).
			AddItem(resultsTextView, 0, 3, false),
		0, 1, true)

	// Run the application with mouse support enabled
	if err := application.SetRoot(flexLayout, true).SetFocus(flexLayout).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	return application
}
