package config

import "time"

const TempFilePath = "../../internal/usecase/temp/tempcardapio"
const CardapioUrl = "https://www.restaurante.uff.br/cardapiomobile.xml"
const ReloadRate = 30 * time.Minute

// Security

const RequestLimit = 10
const BanDuration = 5 * time.Minute
const WindowDuration = 1 * time.Minute

// Database
