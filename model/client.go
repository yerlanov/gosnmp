package model

import (
	"database/sql"
	"test/database"
)

type Client struct {
	ClientIp string
	IpTkd    string
	IpAgu    string
	NameAgu  string
	NameTkd  string
	Port     string
}

type ClientModel struct {
	ClientIp sql.NullString `db:"ip_clients"`
	IpTkd    sql.NullString `db:"ip_tkd"`
	IpAgu    sql.NullString `db:"ip_agu"`
	NameAgu  sql.NullString `db:"hostname_agu"`
	NameTkd  sql.NullString `db:"hostname_tkd"`
	Port     sql.NullString `db:"port_clients"`
}

func GetClientByLogin(login string) (Client, error) {
	var (
		client Client
		model  ClientModel
	)
	query := `SELECT fq.ip_clients, fq.ip_tkd, fq.ip_agu, fq.hostname_agu, fq.hostname_tkd, fq."port_clients"
			  FROM fix_quality.fix_quality fq
			  WHERE fq.ip_tkd <> '0.0.0.0'
			  AND fq.username = $1
			  ORDER BY fq.radius_end_session desc, fq.dhcp_end_session desc
			  LIMIT 1`

	err := database.DB.Get(&model, query, login)
	if err != nil {
		return client, err
	}
	return client, nil
}

func (model ClientModel) ConvertToGoType() Client {
	return Client{
		ClientIp: model.ClientIp.String,
		IpTkd:    model.IpTkd.String,
		IpAgu:    model.IpAgu.String,
		NameAgu:  model.NameAgu.String,
		NameTkd:  model.NameTkd.String,
		Port:     model.Port.String,
	}
}
