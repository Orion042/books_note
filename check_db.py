import sqlite3

dbname = "./db/book_note.db"

conn = sqlite3.connect(dbname)

cur = conn.cursor()

print("=================")

cur.execute("select * from sqlite_master where type='table'")

for x in cur.fetchall():
    print(x)

print("=================")

cur.execute("SELECT * FROM book_infos")

names = list(map(lambda x: x[0], cur.description))

print(names)

for x in cur.fetchall():
    print(x)

print("=================\n")

cur.execute("SELECT * FROM tags")

names = list(map(lambda x: x[0], cur.description))

print(names)

for x in cur.fetchall():
    print(x)

conn.close()