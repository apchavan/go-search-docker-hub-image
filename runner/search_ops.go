package runner

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Structure that holds necessary details regarding image to search from Docker Hub.
type ImageSearchDetails struct {
	ImageName                      string // Stores the name of image.
	MinimumStars                   int    // Stores the minimum number of stars to consider while searching.
	AutomatedOrManualBuildType     int    // Stores whether Default(0), Automated(1) or Manual(2) build type for search.
	OfficialOrCommunityBuildStatus int    // Stores whether Default(0), Official(1) or Community(2) build status for search.
	SearchLimit                    int    // Stores the limit of records to return while searching.
}

// Perform search in Docker Hub using provided parameters & return result to caller.
//
// Implementation Reference 1 : https://docs.docker.com/engine/api/sdk/
//
// Implementation Reference 2 : https://docs.docker.com/engine/api/sdk/examples/
//
// Command line reference for 'docker search' : https://docs.docker.com/engine/reference/commandline/search/
func SearchDockerHub(
	imageDetails *ImageSearchDetails,
	searchResponseStringChannel chan string,
) {

	// Create new Docker API client.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// Set `ImageSearchOptions` based on available values from inputs.
	imageSearchOptions := types.ImageSearchOptions{}
	initialKeyValue_FilterPairs := make([]filters.KeyValuePair, 0)

	// Set filtering for minimum star count if specified.
	if imageDetails.MinimumStars > 0 {
		initialKeyValue_FilterPairs = append(
			initialKeyValue_FilterPairs,
			filters.KeyValuePair{
				Key:   "stars",
				Value: fmt.Sprintf("%d", imageDetails.MinimumStars),
			},
		)
	}

	// Set status of whether automated/manual build type is specified.
	if imageDetails.AutomatedOrManualBuildType == 1 {
		initialKeyValue_FilterPairs = append(
			initialKeyValue_FilterPairs,
			filters.KeyValuePair{
				Key:   "is-automated",
				Value: "true",
			},
		)
	} else if imageDetails.AutomatedOrManualBuildType == 2 {
		initialKeyValue_FilterPairs = append(
			initialKeyValue_FilterPairs,
			filters.KeyValuePair{
				Key:   "is-automated",
				Value: "false",
			},
		)
	}

	// Set status of whether official/community build status is specified.
	if imageDetails.OfficialOrCommunityBuildStatus == 1 {
		initialKeyValue_FilterPairs = append(
			initialKeyValue_FilterPairs,
			filters.KeyValuePair{
				Key:   "is-official",
				Value: "true",
			},
		)
	} else if imageDetails.OfficialOrCommunityBuildStatus == 2 {
		initialKeyValue_FilterPairs = append(
			initialKeyValue_FilterPairs,
			filters.KeyValuePair{
				Key:   "is-official",
				Value: "false",
			},
		)
	}

	// If any filter present, then set the `Filters` in `imageSearchOptions` struct.
	if len(initialKeyValue_FilterPairs) > 0 {
		imageSearchOptions.Filters = filters.NewArgs(initialKeyValue_FilterPairs...)
	}

	// Set if search limit is specified.
	if imageDetails.SearchLimit > 0 {
		imageSearchOptions.Limit = imageDetails.SearchLimit
	}

	// Search for docker image.
	results, err := cli.ImageSearch(
		context.Background(),
		imageDetails.ImageName,
		imageSearchOptions,
	)
	if err != nil {
		line1 := fmt.Sprintf("\n\n (❌) Problem while searching for '%s' image...\n", imageDetails.ImageName)
		line2 := "1. Make sure Docker Daemon is running.\n"
		line3 := "2. Make sure system is connected to internet.\n"
		line4 := fmt.Sprintf("\nError : %v\n", err)
		searchResponseStringChannel <- (line1 + line2 + line3 + line4)
		return
	}

	// Handle case if no results found.
	if len(results) == 0 {
		searchResponseStringChannel <- fmt.Sprintf("\n\n (⚠️) No results found for '%s' image...\n", imageDetails.ImageName)
		return
	}

	// Format the result into table.
	prettyTableWriter := table.NewWriter()
	prettyTableWriter.AppendHeader(
		table.Row{"Name", "Stars", "Official?", "Automated?", "Description"},
	)

	for _, result := range results {
		isOfficialIcon := "No"
		isAutomatedIcon := "No"

		if result.IsOfficial {
			isOfficialIcon = "Yes"
		}
		if result.IsAutomated {
			isAutomatedIcon = "Yes"
		}

		prettyTableWriter.AppendRow(
			table.Row{
				result.Name,
				result.StarCount,
				isOfficialIcon,
				isAutomatedIcon,
				result.Description,
			},
		)
	}

	prettyTableWriter.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: false, WidthMax: 27},
		{Number: 2, AutoMerge: false, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Number: 3, AutoMerge: false, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 16},
		{Number: 4, AutoMerge: false, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 16},
		{Number: 5, AutoMerge: false, WidthMax: 38},
	})

	// Make sure the table is sorted depending on Stars count.
	prettyTableWriter.SortBy(
		[]table.SortBy{
			{Name: "Stars", Mode: table.DscNumeric},
		},
	)

	prettyTableWriter.SetStyle(table.StyleRounded)
	prettyTableWriter.SetAutoIndex(true)
	prettyTableWriter.Style().Options.SeparateRows = true
	searchResponseStringChannel <- prettyTableWriter.Render()
}
