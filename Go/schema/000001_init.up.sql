BEGIN;

CREATE TABLE "User"
(
    "ID"       BigSerial              NOT NULL PRIMARY KEY,
    "Username" Character varying(50)  NOT NULL,
    "Salt"     bytea                  NOT NULL,
    "Hash"     Character varying(200) NOT NULL
) WITH (autovacuum_enabled = true);

CREATE TABLE "Shopping_List"
(
    "ID"          BigSerial                                         NOT NULL PRIMARY KEY,
    "User_id"     BigInt REFERENCES "User" ("ID") ON DELETE CASCADE NOT NULL,
    "Title"       Character varying(100)                            NOT NULL,
    "Description" Text
) WITH (autovacuum_enabled = true);

CREATE TABLE "Shopping_Item"
(
    "ID"          BigSerial                                                  NOT NULL PRIMARY KEY,
    "List_id"     BigInt REFERENCES "Shopping_List" ("ID") ON DELETE CASCADE NOT NULL,
    "Title"       Character varying(100)                                     NOT NULL,
    "Description" Text,
    "Checked"     Boolean                                                    NOT NULL DEFAULT false
) WITH (autovacuum_enabled = true);

COMMIT;