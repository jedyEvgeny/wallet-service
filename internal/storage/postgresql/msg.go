package storage

const (
	msgPing              = "Пинг к БД %s выполнен за %v"
	msgMigrationsNotNeed = "нет изменений схемы БД. Миграции не требуются"
	msgMigrationsDone    = "Миграции применены"
)

const (
	errCreateDB           = "ошибка при создании БД: %v"
	errCantOpen           = "ошибка при открытии БД %s: %w"
	errPing               = "ошибка пинга БД %s: %w"
	errDriver             = "не удалось загрузить драйвер: %w"
	isntExistMigrations   = "нет каталога с миграциями %s: %w"
	errInitMigrations     = "ошибка инициализации миграции: %w"
	errExecMigrations     = "ошибка применения миграций: %w"
	errLaunchMigrations   = "ошибка обновления схемы БД %s: %w"
	errParseNotActiveConn = "ошибка при парсинге времени ожидания неактивного соединения: %s"
	errParseLifeConn      = "ошибка при парсинге времени жизни соединения: %s"
)
