package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-announcements/internal/store")
func(s *Server)handleList(w http.ResponseWriter,r *http.Request){list,_:=s.db.List();if list==nil{list=[]store.Announcement{}};writeJSON(w,200,list)}
func(s *Server)handleCreate(w http.ResponseWriter,r *http.Request){var a store.Announcement;json.NewDecoder(r.Body).Decode(&a);if a.Title==""{writeError(w,400,"title required");return};s.db.Create(&a);writeJSON(w,201,a)}
func(s *Server)handleRead(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var req struct{Name string `json:"name"`};json.NewDecoder(r.Body).Decode(&req);if req.Name==""{req.Name="anonymous"};s.db.MarkRead(id,req.Name);writeJSON(w,200,map[string]string{"status":"read"})}
func(s *Server)handleDelete(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Delete(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
