CREATE TABLE IF NOT EXISTS Node (
    NodeId BLOB PRIMARY KEY DEFAULT (randomblob(16)),
    Address   varchar(20) NOT NULL UNIQUE,
	PublicKey BLOB NOT NULL,
	Name nvarchar(512) NOT NULL
);