package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"tipprunde/internal/model"
	"tipprunde/internal/util"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

var ergebnisseCmd = &cobra.Command{
	Use:   "ergebnisse",
	Short: "Kurzbeschreibung",
	Long:  `Lange Beschreibung.`,
	Run: func(cmd *cobra.Command, args []string) {

		matchUrls := []string{
			//täglich
			"https://www.weltfussball.de/alle_spiele/2-bundesliga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/3-liga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/africa-cup-2025-marokko_2/",
			"https://www.weltfussball.de/alle_spiele/wm-quali-afrika-2023-2025/",
			"https://www.weltfussball.de/alle_spiele/asian-cup-qual-2027/",
			"https://www.weltfussball.de/alle_spiele/wm-quali-asien-2023-2025/",
			"https://www.weltfussball.de/alle_spiele/blr-cempionat-2025/",
			"https://www.weltfussball.de/alle_spiele/bih-premier-liga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/bul-parva-liga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/wm-quali-concacaf-2024-2025/",
			"https://www.weltfussball.de/alle_spiele/conference-league-qual-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/den-superliga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/eng-premier-league-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/est-meistriliiga-2025/",
			"https://www.weltfussball.de/alle_spiele/europa-league-qual-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/wm-quali-europa-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/fin-veikkausliiga-2025/",
			"https://www.weltfussball.de/alle_spiele/fra-ligue-1-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/franz-beckenbauer-supercup-2025/",
			"https://www.weltfussball.de/alle_spiele/fro-effodeildin-2025/",
			"https://www.weltfussball.de/alle_spiele/geo-erovnuli-liga-2025/",
			"https://www.weltfussball.de/alle_spiele/gold-cup-2025/",
			"https://www.weltfussball.de/alle_spiele/irl-premier-division-2025/",
			"https://www.weltfussball.de/alle_spiele/isl-urvalsdeild-2025/",
			"https://www.weltfussball.de/alle_spiele/isr-ligat-haal-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/ita-serie-a-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/kaz-premier-liga-2025/",
			"https://www.weltfussball.de/alle_spiele/lat-virsliga-2025/",
			"https://www.weltfussball.de/alle_spiele/ltu-a-lyga-2025/",
			"https://www.weltfussball.de/alle_spiele/mlt-premier-league-2025-2026-opening/",
			"https://www.weltfussball.de/alle_spiele/mda-divizia-nationala-2025-2026-phase-i/",
			"https://www.weltfussball.de/alle_spiele/ned-eredivisie-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/nir-premier-league-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/nor-eliteserien-2025/",
			"https://www.weltfussball.de/alle_spiele/wm-quali-play-offs-2026/",
			"https://www.weltfussball.de/alle_spiele/pol-ekstraklasa-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/rou-liga-1-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/rus-premier-liga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/sco-premiership-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/swe-allsvenskan-2025/",
			"https://www.weltfussball.de/alle_spiele/sui-super-league-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/wm-quali-suedamerika-2023-2025/",
			"https://www.weltfussball.de/alle_spiele/srb-super-liga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/svk-super-liga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/svn-prvaliga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/esp-primera-division-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/aut-bundesliga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/tur-sueperlig-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/cze-1-fotbalova-liga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/ukr-premyer-liga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/hun-nb-i-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/klub-wm-2025/",

			// minütlich
			"https://www.weltfussball.de/alle_spiele/bundesliga-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/dfb-pokal-2025-2026/",
			"https://www.weltfussball.de/alle_spiele/wm-2026-kanada-mexiko-usa_2/",
		}

		allMatches := []model.Match{}

		for _, matchURL := range matchUrls {

			fmt.Println("------")
			fmt.Printf("Fetching match data from: %s\n", matchURL)

			randomDelay := time.Duration(rand.Intn(3000)+2000) * time.Millisecond
			fmt.Printf("Waiting for %v before fetching the URL...\n", randomDelay)
			time.Sleep(randomDelay)

			matches, err := parseMatchURL(matchURL)
			if err != nil {
				fmt.Printf("Error parsing Weltfussball data from %s: %v\n", matchURL, err)
				continue
			}
			allMatches = append(allMatches, matches...)
		}

		saveMatchesToCSV(allMatches, "C:\\temp\\weltfussball.csv")

		fmt.Println("Fertig.")
	},
}

func parseMatchURL(url string) ([]model.Match, error) {
	matches := []model.Match{}

	contest := strings.TrimPrefix(
		strings.TrimSuffix(url, "/"),
		"https://www.weltfussball.de/alle_spiele/")

	reader, err := util.GetURLContent(url)
	if err != nil {
		log.Fatalf("Error fetching URL content: %v", err)
	}
	defer reader.Close()

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("Error loading the HTML document: %v", err)
	}

	currentGameday := ""
	lastKnownDate := (*time.Time)(nil)

	doc.Find("table.standard_tabelle tr").Each(func(i int, s *goquery.Selection) {
		// check if it's a gameday header
		if s.Find("th[colspan='7']").Length() > 0 {
			currentGameday = strings.TrimSpace(s.Find("th[colspan='7'] a").Text())
			fmt.Printf("Spieltag: %s\n", currentGameday)
			return
		}

		// Check if it's a match data row (has at least 6 td elements)
		tds := s.Find("td")
		if tds.Length() < 6 {
			return
		}

		match := &model.Match{
			GameDay: currentGameday,
			Contest: contest,
		}

		// Column 1: Date
		dateStr := strings.TrimSpace(tds.Eq(0).Find("a").Text())
		if dateStr != "" {
			parsedDate, parseErr := time.Parse("02.01.2006", dateStr) // DD.MM.YYYY
			if parseErr != nil {
				log.Printf("Warning: Could not parse match date '%s': %v. Setting date to nil.", dateStr, parseErr)
				match.Date = nil
				lastKnownDate = nil
			} else {
				match.Date = &parsedDate
				lastKnownDate = &parsedDate
			}
		} else {
			// Date is missing, use the last known date
			match.Date = lastKnownDate
		}

		// Column 2: Time
		match.Time = strings.TrimSpace(tds.Eq(1).Text())

		// Column 3: Home Team
		homeTeamCell := tds.Eq(2)
		homeTeamName := strings.TrimSpace(homeTeamCell.Find("a").Text())
		if homeTeamName == "" {
			homeTeamName = strings.TrimSpace(homeTeamCell.Text())
		}
		match.HomeTeam = homeTeamName

		// Column 5: Away Team
		awayTeamCell := tds.Eq(4)
		// Try to get text from <a> first, if not found, get text directly from <td>
		awayTeamName := strings.TrimSpace(awayTeamCell.Find("a").Text())
		if awayTeamName == "" {
			awayTeamName = strings.TrimSpace(awayTeamCell.Text())
		}
		match.AwayTeam = awayTeamName

		// Column 6: Score
		scoreText := strings.TrimSpace(tds.Eq(5).Find("a").Text())
		scoreRegexp := regexp.MustCompile(`^(\d+:\d+)`) // Matches "2:3" from "2:3 (0:2)"
		scoreMatches := scoreRegexp.FindStringSubmatch(scoreText)
		if len(scoreMatches) > 1 {
			match.Score = scoreMatches[1]
		} else {
			match.Score = ""
		}

		printMatchSummary(match)

		if match.HomeTeam == "" || match.AwayTeam == "" {
			log.Printf("Warning: Skipping incomplete match row due to missing team names (Gameday: %s, Date: %v). Row HTML: %s", match.GameDay, match.Date, s.Text())
			return
		}

		matches = append(matches, *match)
	})

	return matches, nil
}

func saveMatchesToCSV(matches []model.Match, filePath string) error {
	fmt.Printf("Attempting to write %d matches to CSV: %s\n", len(matches), filePath)
	if err := util.WriteStructsToCSV(matches, filePath); err != nil {
		return fmt.Errorf("failed to write players to CSV: %w", err)
	}
	return nil
}

func printMatchSummary(match *model.Match) {
	fmt.Printf("  Match: %s vs %s %s\n",
		match.HomeTeam, match.AwayTeam, match.Score)

}
