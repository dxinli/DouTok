install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

up:
	goose mysql "root:a12bCd3_W45pUq6@tcp(localhost:3306)/doutok?parseTime=true" up

down:
	goose mysql "root:a12bCd3_W45pUq6@tcp(localhost:3306)/doutok?parseTime=true" down

create:
	goose mysql "root:a12bCd3_W45pUq6@tcp(localhost:3306)/doutok?parseTime=true" create ${name} sql
