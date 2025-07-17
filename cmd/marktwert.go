package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"

	"tipprunde/internal/model"
	"tipprunde/internal/util"
)

var marktwertCmd = &cobra.Command{
	Use:   "marktwert",
	Short: "Kurzbeschreibung",
	Long:  `Lange Beschreibung.`,
	Run: func(cmd *cobra.Command, args []string) {

		clubUrls := []string{
			"https://www.transfermarkt.de/fc-bayern-munchen/kader/verein/27/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/bayer-04-leverkusen/kader/verein/15/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/hertha-bsc/kader/verein/44/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/fc-schalke-04/kader/verein/33/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/tsg-1899-hoffenheim/kader/verein/533/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sv-werder-bremen/kader/verein/86/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/1-fsv-mainz-05/kader/verein/39/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sc-freiburg/kader/verein/60/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/rb-leipzig/kader/verein/23826/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/eintracht-frankfurt/kader/verein/24/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/borussia-dortmund/kader/verein/16/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/1-fc-union-berlin/kader/verein/89/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/1-fc-koln/kader/verein/3/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/borussia-monchengladbach/kader/verein/18/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/vfl-wolfsburg/kader/verein/82/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/vfl-bochum/kader/verein/80/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/fc-augsburg/kader/verein/167/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/vfb-stuttgart/kader/verein/79/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/1-fc-kaiserslautern/kader/verein/2/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/spvgg-greuther-furth/kader/verein/65/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/fc-st-pauli/kader/verein/35/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sc-paderborn-07/kader/verein/127/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/fc-hansa-rostock/kader/verein/30/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sv-sandhausen/kader/verein/254/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/ssv-jahn-regensburg/startseite/verein/109",
			"https://www.transfermarkt.de/1-fc-magdeburg/kader/verein/187/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/eintracht-braunschweig/kader/verein/23/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/hannover-96/kader/verein/42/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/holstein-kiel/kader/verein/269/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/1-fc-nurnberg/kader/verein/4/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/karlsruher-sc/kader/verein/48/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/hertha-bsc/kader/verein/44/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/arminia-bielefeld/kader/verein/10/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sv-darmstadt-98/kader/verein/105/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/fortuna-dusseldorf/kader/verein/38/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/hamburger-sv/kader/verein/41/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/vfl-osnabruck/kader/verein/81/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/vfb-oldenburg/kader/verein/166/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/fc-ingolstadt-04/kader/verein/4795/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sv-waldhof-mannheim/kader/verein/85/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/rot-weiss-essen/kader/verein/56/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sg-dynamo-dresden/kader/verein/129/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/1-fc-saarbrucken/kader/verein/1/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sc-freiburg-ii/kader/verein/245/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/fsv-zwickau/kader/verein/275/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sv-wehen-wiesbaden/kader/verein/108/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/msv-duisburg/kader/verein/52/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sv-meppen/kader/verein/247/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/spvgg-bayreuth/kader/verein/752/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/fc-viktoria-koln/kader/verein/1622/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sv-07-elversberg/kader/verein/64/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/tsv-1860-munchen/kader/verein/72/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/sc-verl/kader/verein/93/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/fc-erzgebirge-aue/kader/verein/94/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/hallescher-fc/kader/verein/440/saison_id/2025/plus/1",
			"https://www.transfermarkt.de/borussia-dortmund-ii/kader/verein/17/saison_id/2025/plus/1",
		}

		allPlayers := []model.Player{}

		for _, clubURL := range clubUrls {

			fmt.Println("------")
			fmt.Printf("Fetching market values and player data from: %s\n", clubURL)

			randomDelay := time.Duration(rand.Intn(3000)+2000) * time.Millisecond
			fmt.Printf("Waiting for %v before fetching the URL...\n", randomDelay)
			time.Sleep(randomDelay)

			players, err := parsePlayersURL(clubURL)
			if err != nil {
				log.Printf("Error parsing Transfermarkt data from %s: %v", clubURL, err)
				continue
			}
			allPlayers = append(allPlayers, players...)
		}

		savePlayersToCSV(allPlayers, "C:\\temp\\transfermarkt.csv")

		fmt.Println("Fertig.")
	},
}

func parsePlayersURL(url string) ([]model.Player, error) {
	players := []model.Player{}

	reader, err := util.GetURLContent(url)
	if err != nil {
		log.Fatalf("Error fetching URL content: %v", err)
	}
	defer reader.Close()

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("Error loading the HTML document: %v", err)
	}

	clubName := strings.TrimSpace(doc.Find(".data-header__headline-wrapper").First().Contents().Not("span").Text())
	fmt.Printf("  Club: %s\n", clubName)

	doc.Find("table.items>tbody>tr").Each(func(i int, s *goquery.Selection) {
		player, err := parsePlayerRow(s)
		if err != nil {
			log.Printf("Warning: Could not parse player row %d: %v\n", i+1, err)
			return
		}
		player.Club = clubName

		players = append(players, *player)
	})

	return players, nil
}

func parsePlayerRow(s *goquery.Selection) (*model.Player, error) {
	player := &model.Player{}

	// Column 2: Player Name, Position, Injury Status
	playerCol := s.Find("td.posrela")
	player.Name = strings.TrimSpace(playerCol.Find(".hauptlink a").First().Text())
	player.Position = strings.TrimSpace(playerCol.Find("td").Last().Text())

	injurySpan := s.Find("span.verletzt-table")
	if injuryStatus, exists := injurySpan.Attr("title"); exists {
		player.InjuryStatus = strings.TrimSpace(injuryStatus)
	}

	// Column 3: Birthday (08.08.2003 (21))
	birthdayStr := strings.TrimSpace(s.Find("td.zentriert").Eq(1).Text())
	if birthdayStr != "" {
		re := regexp.MustCompile(`^(\d{2}\.\d{2}\.\d{4})`)
		match := re.FindStringSubmatch(birthdayStr)
		if len(match) > 1 {
			parsedDate, err := time.Parse("02.01.2006", match[1])
			if err != nil {
				log.Printf("Warning: Could not parse birthday '%s': %v", match[1], err)
			} else {
				player.Birthday = parsedDate
			}
		}
	}

	// Column 5: Height (1,89m)
	player.Height = strings.TrimSpace(s.Find("td.zentriert").Eq(3).Text())

	// Column 6: Foot (beidfüßig, rechts, links)
	player.Foot = strings.TrimSpace(s.Find("td.zentriert").Eq(4).Text())

	// Column 7: Joined At (27.01.2025)
	joinedAtStr := strings.TrimSpace(s.Find("td.zentriert").Eq(5).Text())
	if joinedAtStr != "" && joinedAtStr != "-" { // Check for non-empty and not just a dash
		parsedDate, err := time.Parse("02.01.2006", joinedAtStr)
		if err != nil {
			log.Printf("Warning: Could not parse joined date '%s': %v", joinedAtStr, err)
		} else {
			player.JoinedAt = &parsedDate // Assign pointer to parsed date
		}
	}

	// Last Column: Market Value (12,00 Mio. €, 500 Tsd. €, -)
	marketValueAnchor := s.Find("td.rechts.hauptlink a")
	marketValueStr := strings.TrimSpace(marketValueAnchor.Text())
	player.MarketValue = parseMarketValue(marketValueStr)

	// show some player info to the console
	printPlayer(player)

	if player.Name == "" {
		return nil, fmt.Errorf("Spielername nicht gefunden.")
	}
	return player, nil
}

func parseMarketValue(mv string) int {
	mv = strings.ToLower(mv)
	mv = strings.ReplaceAll(mv, " ", "")  // Remove spaces
	mv = strings.ReplaceAll(mv, ".", "")  // Remove thousand separators (e.g., "12.000.000")
	mv = strings.ReplaceAll(mv, "€", "")  // Remove Euro symbol
	mv = strings.ReplaceAll(mv, ",", ".") // Replace comma with dot for float conversion

	if mv == "-" || mv == "" {
		return 0
	}

	var value float64
	var multiplier float64 = 1.0

	if strings.Contains(mv, "mio") {
		multiplier = 1_000_000
		mv = strings.ReplaceAll(mv, "mio", "")
	} else if strings.Contains(mv, "tsd") {
		multiplier = 1_000
		mv = strings.ReplaceAll(mv, "tsd", "")
	}

	parsedValue, err := strconv.ParseFloat(mv, 64)
	if err != nil {
		log.Printf("Warning: Could not parse market value '%s': %v", mv, err)
		return 0
	}

	value = parsedValue * multiplier
	return int(value)
}

func savePlayersToCSV(players []model.Player, filePath string) error {
	fmt.Printf("Attempting to write %d players to CSV: %s\n", len(players), filePath)
	if err := util.WriteStructsToCSV(players, filePath); err != nil {
		return fmt.Errorf("failed to write players to CSV: %w", err)
	}
	return nil
}

func printPlayer(player *model.Player) {
	injuryStatusInfo := ""
	if player.InjuryStatus != "" {
		injuryStatusInfo = fmt.Sprintf(" [Injury: %s]", player.InjuryStatus)
	}
	fmt.Printf("  Player: %s (MV: %d €%s)\n", player.Name, player.MarketValue, injuryStatusInfo)
}
