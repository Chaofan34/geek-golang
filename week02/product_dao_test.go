package week02

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/Chaofan34/geek-golang/pkg/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestFindProductByCode(t *testing.T) {
	db, err := mysql.OpenDB("172.17.0.1", 3306, "root", "123456", "rec_movie")
	if err != nil {
		panic(fmt.Sprintf("open db failed %v", err))
	}
	product, err := FindProductByCode(db, "D40") // not find
	log.Printf("product = %+v, err = %+v\n", product, err)
	assert.Equal(t, errors.Is(err, gorm.ErrRecordNotFound), true, "err should be errNotFound")
}
