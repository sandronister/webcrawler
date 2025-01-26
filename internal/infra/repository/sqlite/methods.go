package sqlite

func (m *Model) Insert(url, content string) error {
	_, err := m.db.Exec("INSERT INTO pages (url, content) VALUES (?, ?)", url, content)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) Create() error {
	_, err := m.db.Exec(`CREATE TABLE IF NOT EXISTS pages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT NOT NULL,
		content TEXT NOT NULL,
		time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
	);`)

	if err != nil {
		m.logger.Error(err.Error())
		return err
	}

	_, err = m.db.Exec(`CREATE TABLE IF NOT EXISTS domains (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		domain TEXT NOT NULL,
		time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
	);`)

	if err != nil {
		m.logger.Error(err.Error())
		return err
	}

	return nil
}
