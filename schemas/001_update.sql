CREATE TABLE IF NOT EXISTS Node (
    NodeId INTEGER PRIMARY KEY AUTOINCREMENT,
    Address   varchar(20) NOT NULL UNIQUE,
	PublicKey varchar(512) NOT NULL,
	Name nvarchar(512) NOT NULL
);