package generate

type SchemaType string

const (
	SchemaTypeSocialMedia   SchemaType = "social_media"
	SchemaTypeFintech       SchemaType = "fintech"
	SchemaTypeSakila        SchemaType = "sakila"
	SchemaTypeStackOverflow SchemaType = "stackoverflow"
	SchemaTypeLogictics     SchemaType = "logistics"
	SchemaTypeWebshop       SchemaType = "webshop"
)

var socialMedia = []string{
	"_users", "_followers", "_posts", "_comments", "_likes", "_shares", "_groups", "_pages", "_events", "_messages", "_notifications", "_settings", "_profile", "_media", "_tags", "_locations", "_checkins", "_friends", "_family", "_colleagues", "_classmates", "_neighbors", "_acquaintances", "_strangers", "_blocked", "_muted", "_restricted", "_unrestricted", "_banned", "_deleted", "_archived", "_hidden", "_suspended", "_verified", "_unverified", "_pending", "_approved", "_rejected", "_reported", "_flagged", "_spam",
}

var fintech = []string{
	"_users", "_accounts", "_transactions", "_payments", "_loans", "_deposits", "_withdrawals", "_transfers", "_balances", "_statements", "_invoices", "_bills", "_purchases", "_sales", "_orders", "_carts", "_checkout", "_shipping", "_delivery", "_returns", "_refunds", "_disputes", "_claims", "_fraud", "_scams", "_phishing", "_hacking", "_security", "_privacy", "_encryption", "_decryption", "_authentication", "_authorization", "_verification", "_validation", "_confirmation", "_rejection", "_approval", "_reporting", "_flagging", "_spam",
}

var sakila = []string{
	"_actors", "_categories", "_films", "_languages", "_customers", "_stores", "_staff", "_inventory", "_rentals", "_payments", "_cities", "_countries", "_addresses", "_districts", "_postal_codes", "_phone_numbers", "_emails", "_websites", "_social_media", "_profiles", "_logins", "_passwords", "_security_questions", "_answers", "_verification", "_confirmation", "_rejection", "_approval", "_reporting", "_flagging", "_spam",
}

var stackOverflow = []string{
	"_users", "_questions", "_answers", "_comments", "_votes", "_badges", "_tags", "_categories", "_reputation", "_privileges", "_settings", "_profile", "_notifications", "_messages", "_inbox", "_outbox", "_drafts", "_trash", "_spam", "_flagged", "_reported", "_blocked", "_muted", "_restricted", "_unrestricted", "_banned", "_deleted", "_archived", "_hidden", "_suspended", "_verified", "_unverified", "_pending", "_approved", "_rejected", "_reported", "_flagged", "_spam",
}

var logistics = []string{
	"_users", "_customers", "_suppliers", "_vendors", "_partners", "_employees", "_contractors", "_drivers", "_vehicles", "_warehouses", "_inventory", "_products", "_orders", "_shipments", "_deliveries", "_returns", "_refunds", "_invoices", "_bills", "_purchases", "_sales", "_transactions", "_payments", "_loans", "_deposits", "_withdrawals", "_transfers", "_balances", "_statements", "_invoices", "_bills", "_purchases", "_sales", "_orders", "_carts", "_checkout", "_shipping", "_delivery", "_returns", "_refunds", "_disputes", "_claims", "_fraud", "_scams", "_phishing", "_hacking", "_security", "_privacy", "_encryption", "_decryption", "_authentication", "_authorization", "_verification", "_validation", "_confirmation", "_rejection", "_approval", "_reporting", "_flagging", "_spam",
}

var webshop = []string{
	"_cart", "_checkout", "_shipping", "_delivery", "_returns", "_refunds", "_disputes", "_claims", "_fraud", "_scams", "_phishing", "_hacking", "_security", "_privacy", "_encryption", "_decryption", "_authentication", "_authorization", "_verification", "_validation", "_confirmation", "_rejection", "_approval", "_reporting", "_flagging", "_spam",
}
