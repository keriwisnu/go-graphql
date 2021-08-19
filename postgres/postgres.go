
package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


type Db struct {
	*sql.DB
}

func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func ConnString(host string, port int, user string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbName,
	)
}


type Cart struct {
	ID         int
	Name       string
	Quantity   int
	Price	   int
	Totprice   int
}

func (d *Db) SaveCart(name string) []Cart {
	stmt, err := d.Prepare("INSERT INTO cart (id, name, quantity, price, totprice) values ($1, $2, $3, $4, $5)")
	if err != nil {
		fmt.Println(" Insert Preperation Err: ", err)
	}

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("SaveCart Query Err: ", err)
	}


	cart := []Cart{}
	for rows.Next() {
		var r Cart
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Quantity,
			&r.Price,
		)

		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		r.Totprice = r.Quantity * r.Price
		cart = append(cart, r)

	}

	return cart
}