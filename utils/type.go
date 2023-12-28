package utils

type Key string
type Premium string

const (
	KeyUserId        Key = "user_id"
	KeyCoordinatorTx Key = "coordinator_tx"

	PremiumSwipe    Premium = "swipe"
	PremiumVerified Premium = "verified"

	FormatYYYYMMDD = "2006-01-02" // YYYY-MM-DD

)
