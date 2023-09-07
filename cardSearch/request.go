package cardSearch

import (
	"log"
	"mtg-card-manager/cardList"

	"github.com/MagicTheGathering/mtg-sdk-go"
)

func searchCard(m model) {
	cards, err := mtg.NewQuery().Where(mtg.CardName, m.inputs[0].Value()).All()
	if err != nil {
		log.Fatal(err)
	}
	var cardNameList []string
	for _, card := range cards {
		cardNameList = append(cardNameList, card.Name)
	}
	// for _, item := range cardNameList {
	// 	log.Fatal(item)
	// }
	cardList.ShowList(cardNameList)
	// p := tea.NewProgram(cardList.InitialModel())
	// l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	// l.Title = "What do you want for dinner?"
	// l.SetShowStatusBar(false)
	// l.SetFilteringEnabled(false)
	// l.Styles.Title = titleStyle
	// l.Styles.PaginationStyle = paginationStyle
	// l.Styles.HelpStyle = helpStyle
	//
	// m := model{list: l}
	//
	// if _, err := tea.NewProgram(m).Run(); err != nil {
	// 	fmt.Println("Error running program:", err)
	// 	os.Exit(1)
	// }
}
