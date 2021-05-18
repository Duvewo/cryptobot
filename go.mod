module github.com/Duvewo/cryptobot

go 1.16

require (
	github.com/gorilla/websocket v1.4.2
	github.com/jmoiron/sqlx v1.3.3
	github.com/lib/pq v1.2.0
	gopkg.in/tucnak/telebot.v3 v3.0.0-00010101000000-000000000000
)

replace gopkg.in/tucnak/telebot.v3 => ./vendor/gopkg.in/tucnak/telebot.v3
