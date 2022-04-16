package store

func (m *Manager) StartBackgroundTasks() (err error) {

	return
}

/*

	views, _, err := ghClient.Repositories.ListTrafficViews(ctx, "xarantolus", "filtrite", &github.TrafficBreakdownOptions{
		Per: "day",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", views)

	paths, _, err := ghClient.Repositories.ListTrafficPaths(ctx, "xarantolus", "filtrite")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", paths)

	refs, _, err := ghClient.Repositories.ListTrafficReferrers(ctx, "xarantolus", "filtrite")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", refs)

	clones, _, err := ghClient.Repositories.ListTrafficClones(ctx, "xarantolus", "filtrite", &github.TrafficBreakdownOptions{
		Per: "day",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", clones)
*/
