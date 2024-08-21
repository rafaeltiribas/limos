package config

import (
	"time"
)

const TempFilePath = "../../internal/usecase/temp/tempcardapio"
const CardapioUrl = "https://www.restaurante.uff.br/cardapiomobile.xml"
const ReloadRate = 30 * time.Minute

const EnvPath = "../../.env"

// Security

const RequestLimit = 10
const BanDuration = 5 * time.Minute
const WindowDuration = 1 * time.Minute

// Database
const DbDriver = "postgres"
const DbName = "cardapiouff"
const DbUser = "your_user"
const DbPassword = "your_password"
const DbHost = "localhost"
const DbPort = "5432"
