package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Announcement struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Author string `json:"author"`
	Priority string `json:"priority"`
	Channel string `json:"channel"`
	Pinned int `json:"pinned"`
	ExpiresAt string `json:"expires_at"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"announcements.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS announcements(id TEXT PRIMARY KEY,title TEXT NOT NULL,body TEXT DEFAULT '',author TEXT DEFAULT '',priority TEXT DEFAULT 'normal',channel TEXT DEFAULT 'general',pinned INTEGER DEFAULT 0,expires_at TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Announcement)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO announcements(id,title,body,author,priority,channel,pinned,expires_at,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.Title,e.Body,e.Author,e.Priority,e.Channel,e.Pinned,e.ExpiresAt,e.CreatedAt);return err}
func(d *DB)Get(id string)*Announcement{var e Announcement;if d.db.QueryRow(`SELECT id,title,body,author,priority,channel,pinned,expires_at,created_at FROM announcements WHERE id=?`,id).Scan(&e.ID,&e.Title,&e.Body,&e.Author,&e.Priority,&e.Channel,&e.Pinned,&e.ExpiresAt,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Announcement{rows,_:=d.db.Query(`SELECT id,title,body,author,priority,channel,pinned,expires_at,created_at FROM announcements ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Announcement;for rows.Next(){var e Announcement;rows.Scan(&e.ID,&e.Title,&e.Body,&e.Author,&e.Priority,&e.Channel,&e.Pinned,&e.ExpiresAt,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM announcements WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM announcements`).Scan(&n);return n}
