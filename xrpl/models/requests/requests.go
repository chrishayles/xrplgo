package requests

import "fmt"

type RequestMethod string

const (
	// account methods
	ACCOUNT_CHANNELS   RequestMethod = "account_channels"
	ACCOUNT_CURRENCIES RequestMethod = "account_currencies"
	ACCOUNT_INFO       RequestMethod = "account_info"
	ACCOUNT_LINES      RequestMethod = "account_lines"
	ACCOUNT_OBJECTS    RequestMethod = "account_objects"
	ACCOUNT_OFFERS     RequestMethod = "account_offers"
	ACCOUNT_TX         RequestMethod = "account_tx"
	GATEWAY_BALANCES   RequestMethod = "gateway_balances"
	NO_RIPPLE_CHECK    RequestMethod = "noripple_check"

	// transaction methods
	SIGN               RequestMethod = "sign"
	SIGN_FOR           RequestMethod = "sign_for"
	SUBMIT             RequestMethod = "submit"
	SUBMIT_MULTISIGNED RequestMethod = "submit_multisigned"
	TRANSACTION_ENTRY  RequestMethod = "transaction_entry"
	TX                 RequestMethod = "tx"

	// channel methods
	CHANNEL_AUTHORIZE RequestMethod = "channel_authorize"
	CHANNEL_VERIFY    RequestMethod = "channel_verify"

	// path methods
	BOOK_OFFERS        RequestMethod = "book_offers"
	DEPOSIT_AUTHORIZED RequestMethod = "deposit_authorized"
	PATH_FIND          RequestMethod = "path_find"
	RIPPLE_PATH_FIND   RequestMethod = "ripple_path_find"

	// ledger methods
	LEDGER         RequestMethod = "ledger"
	LEDGER_CLOSED  RequestMethod = "ledger_closed"
	LEDGER_CURRENT RequestMethod = "ledger_current"
	LEDGER_DATA    RequestMethod = "ledger_data"
	LEDGER_ENTRY   RequestMethod = "ledger_entry"

	// subscription methods
	SUBSCRIBE   RequestMethod = "subscribe"
	UNSUBSCRIBE RequestMethod = "unsubscribe"

	// server info methods
	FEE          RequestMethod = "fee"
	MANIFEST     RequestMethod = "manifest"
	SERVER_INFO  RequestMethod = "server_info"
	SERVER_STATE RequestMethod = "server_state"

	// utility methods
	PING   RequestMethod = "ping"
	RANDOM RequestMethod = "random"
)

func Test(rm RequestMethod) {
	fmt.Println(string(rm))
}
