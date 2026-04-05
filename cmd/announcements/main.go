package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-announcements/internal/server";"github.com/stockyard-dev/stockyard-announcements/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="9700"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./announcements-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("announcements: %v",err)};defer db.Close();srv:=server.New(db,server.DefaultLimits())
fmt.Printf("\n  Announcements — Self-hosted internal announcement board\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n  Questions? hello@stockyard.dev — I read every message\n\n",port,port)
log.Printf("announcements: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
