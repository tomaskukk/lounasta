package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/tomaskukk/lounasta/api"

	"golang.org/x/exp/rand"
)

var foodEmojis = []string{"🍕", "🍔", "🍣", "🍤", "🍜", "🥗", "🍝", "🍱", "🍛", "🥩", "🍖", "🥙", "🌮", "🍲", "🍚", "🍤", "🥪"}

func PrintRestaurants(restaurants []api.Restaurant, food string) {
	fmt.Println("\nLunch Restaurants Near You:")
	fmt.Println("========================================")

	for _, restaurant := range restaurants {
		emoji := getRandomEmoji()
		highlightedName := highlightRestaurant(restaurant.Name)
		fmt.Printf("%s %s (Distance: %s)\n", emoji, highlightedName, restaurant.Distance)
		fmt.Println("----------------------------------------")

		for _, dish := range restaurant.Dishes {
			if food != "" {
				// Apply highlighting to the food term
				dish = highlightMatchedFood(dish, food)
			}
			dishName, dietaryInfo := parseDish(dish)
			fmt.Printf("  • %-40s %s\n", dishName, formatDietaryInfo(dietaryInfo))
		}

		fmt.Println("========================================\n")
	}
}

// getRandomEmoji returns a random emoji from the foodEmojis slice.
func getRandomEmoji() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	return foodEmojis[rand.Intn(len(foodEmojis))]
}

// highlightText wraps the text with ANSI codes for bold and yellow color.
func highlightRestaurant(text string) string {
	return fmt.Sprintf("\033[1;33m%s\033[0m", text) // Bold Yellow
}

func highlightMatchedFood(text, term string) string {
	highlightStart := "\033[1;34m"
	highlightEnd := "\033[0m"

	lowerText := strings.ToLower(text)
	lowerTerm := strings.ToLower(term)

	// Replace all occurrences of the term with the highlighted version
	highlightedText := strings.ReplaceAll(lowerText, lowerTerm, highlightStart+lowerTerm+highlightEnd)

	return highlightedText
}

// parseDish separates the dish name from dietary information.
func parseDish(dish string) (string, string) {
	// Split the dish into the main name and dietary information by searching for "   " pattern
	parts := strings.Split(dish, "   ")
	dishName := parts[0]
	dietaryInfo := ""
	if len(parts) > 1 {
		dietaryInfo = parts[1]
	}

	dietaryInfo = strings.TrimSpace(dietaryInfo)

	return dishName, dietaryInfo
}

// formatDietaryInfo formats dietary restrictions in a compact manner.
func formatDietaryInfo(dietaryInfo string) string {
	if dietaryInfo == "" {
		return ""
	}
	return fmt.Sprintf("(%s)", dietaryInfo)
}

func parseCliArgs() {

}
