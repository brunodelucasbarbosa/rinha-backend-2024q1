run:
	go run main.go

up:
	docker-compose up --build -d

down:
	docker-compose down

t:
	powershell -ExecutionPolicy ByPass -File "c:\Users\brunobarbosa\Desktop\Bruno\codigos\api_rinha\executar-teste-local.ps1"