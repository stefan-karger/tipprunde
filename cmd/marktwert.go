package cmd

import (
	"fmt"
	"log"
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
		parseURL("https://www.transfermarkt.de/fc-bayern-munchen/kader/verein/27/saison_id/2025/plus/1")
	},
}

func init() {
	rootCmd.AddCommand(marktwertCmd)
}

func parseURL(url string) ([]model.Player, error) {
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

	doc.Find("table.items>tbody>tr").Each(func(i int, s *goquery.Selection) {
		parsePlayerRow(s)
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
	log.Printf("Attempting to write %d players to CSV: %s", len(players), filePath)
	if err := util.WriteStructsToCSV(players, filePath); err != nil {
		return fmt.Errorf("failed to write players to CSV: %w", err)
	}
	return nil
}

func printPlayer(player *model.Player) {
	joinedAtStr := "N/A"
	if player.JoinedAt != nil {
		joinedAtStr = player.JoinedAt.Format("2006-01-02")
	}
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("  Position: %s\n", player.Position)
	fmt.Printf("  Birthday: %s\n", player.Birthday.Format("2006-01-02"))
	fmt.Printf("  Height: %s\n", player.Height)
	fmt.Printf("  Foot: %s\n", player.Foot)
	fmt.Printf("  Joined At: %s\n", joinedAtStr)
	fmt.Printf("  Market Value: %d €\n", player.MarketValue)
	fmt.Printf("  Injury Status: %s\n", player.InjuryStatus)
	fmt.Println("---")
}
