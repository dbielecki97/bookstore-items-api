package domain

type Item struct {
	ID                string      `json:"id,omitempty"`
	Seller            int64       `json:"seller,omitempty"`
	Title             string      `json:"title,omitempty"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures,omitempty"`
	Video             string      `json:"video,omitempty"`
	Price             float32     `json:"price,omitempty"`
	AvailableQuantity int         `json:"available_quantity,omitempty"`
	SoldQuantity      int         `json:"sold_quantity,omitempty"`
	Status            string      `json:"status,omitempty"`
}

type Description struct {
	PlainText string `json:"plain_text,omitempty"`
	HTML      string `json:"html,omitempty"`
}

type Picture struct {
	ID  int64  `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}
