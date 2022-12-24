import sqlite3

dbname = "./db/book_note.db"

conn = sqlite3.connect(dbname)

cur = conn.cursor()

#cur.execute("select * from sqlite_master where type='table'")

cur.execute("SELECT * FROM book_infos")

for x in cur.fetchall():
    print(x)

conn.close()