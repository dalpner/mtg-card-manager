# mtg-card-manager
This project is about to build a database about the cardbase you own.



## first database design:

```mermaid
erDiagram
  CARD {
    uint id
    string card_name
    uint language_id
    bool foil
    uint set_id
    int set_number
    int quanitiy
  }
  SET {
    uint id
    string set_name
    int card_numbers
  }

  LANGUAGE {
    uint id
    string language
  }

  CARD ||--|{ SET : set_id
  CARD ||--|{ LANGUAGE : language_id
```