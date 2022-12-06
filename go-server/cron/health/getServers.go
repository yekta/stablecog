package cronHealth

import (
	"log"

	"github.com/supabase/postgrest-go"
	shared "github.com/yekta/stablecog/go-server/shared"
)

func GetServers() []shared.SDBServer {
	var servers []shared.SDBServer
	_, err := shared.Supabase.From("server").
		Select("id,url,enabled,healthy,last_health_check_at,created_at,updated_at", "", false).
		Order("created_at", &postgrest.OrderOpts{Ascending: false}).
		ExecuteTo(&servers)
	if err != nil {
		log.Println(err)
		return nil
	}
	return servers
}
