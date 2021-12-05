

type NFTAccount struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	ProxyID   int       `json:"proxy_id" db:"proxy_id"`
	DisLogin  string    `json:"dis_login" db:"dis_login"`
	DisPass   string    `json:"dis_pass" db:"dis_pass"`
	DisNumber string    `json:"dis_number" db:"dis_number"`
	DisToken  string    `json:"dis_token" db:"dis_token"`
	CreatedAT time.Time `json:"created_at" db:"created_at"`
	UpdatedAT time.Time `json:"updated_at" db:"updated_at"`
}

create table if not exists nft_accounts (
     id serial not null unique,
     name varchar,
     proxy_id varchar,
     dis_login varchar,
     dis_pass varchar,
     dis_number varchar,
     dis_token varchar,
     created_at timestamp with time zone,
     updated_at timestamp with time zone,
     unique (dis_login));