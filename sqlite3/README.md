gcc shell.c sqlite3.c -lpthread -ldl -lm -o sqlite3
mv sqlite3 /usr/bin/sqlite3
chmod +x /usr/bin/sqlite3
