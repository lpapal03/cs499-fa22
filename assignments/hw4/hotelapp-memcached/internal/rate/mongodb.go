//go:build mongodb

package rate

import (
	log "github.com/sirupsen/logrus"

	"github.com/ucy-coast/hotel-app/pkg/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"time"
	"github.com/bradfitz/gomemcache/memcache"
)

type DatabaseSession struct {
	MongoSession *mgo.Session
	MemcClient   *memcache.Client
}

func NewDatabaseSession(db_addr string, memc_addr string) *DatabaseSession {
	session, err := mgo.Dial(db_addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("New session successfull...")

	memc_client := memcache.New(memc_addr)
	memc_client.Timeout = time.Second * 2
	memc_client.MaxIdleConns = 512

	return &DatabaseSession{
		MongoSession: session,
		MemcClient: memc_client,
  }
}

func (db *DatabaseSession) LoadDataFromJsonFile(rateJsonPath string) {
	util.LoadDataFromJsonFile(db.MongoSession, "rate-db", "inventory", rateJsonPath)
}

// GetRates gets rates for hotels for specific date range.
func (db *DatabaseSession) GetRates(hotelIds []string) (RatePlans, error) {
	ratePlans := make(RatePlans, 0)



	// for _, hotelID := range hotelIds {
	// 	tmpRatePlans := make(RatePlans, 0)
	// 	err := c.Find(&bson.M{"hotelId": hotelID}).All(&tmpRatePlans)
	// 	if err != nil {
	// 		log.Fatalf("Tried to find hotelId [%v], but got error", hotelID, err.Error())
	// 	} else {
	// 		for _, r := range tmpRatePlans {
	// 			ratePlans = append(ratePlans, r)
	// 		}
	// 	}
	// }

	for _, hotelID := range hotelIds {
		tmpRatePlans := make(RatePlans, 0)
		// first check memcached
		item, err := db.MemcClient.Get(hotelID)
		if err == nil {
			// memcached hit
			log.Infof("Memcached hit: hotel_id == %v\n", hotelID)
			if err = json.Unmarshal(item.Value, item); err != nil {
				log.Warn(err)
			}
			for _, r := range tmpRatePlans {
				ratePlans = append(ratePlans, r)
			}
		} else if err == memcache.ErrCacheMiss {
				// memcached miss, set up mongo connection
				log.Infof("Memcached miss: hotel_id == %v\n", hotelID)
				session := db.MongoSession.Copy()
				defer session.Close()
				c := session.DB("rate-db").C("inventory")
				
				tmpRatePlans := make(RatePlans, 0)
				err := c.Find(&bson.M{"hotelId": hotelID}).All(&tmpRatePlans)
				if err != nil {
					log.Fatalf("Tried to find hotelId [%v], but got error", hotelID, err.Error())
				} else {
					for _, r := range tmpRatePlans {
						ratePlans = append(ratePlans, r)
					}
				}

				rate_json, err := json.Marshal(ratePlans)
				if err != nil {
					log.Errorf("Failed to marshal hotel [id: %v] with err:", hotelID, err)
				}
				memc_str := string(rate_json)

				// write to memcached
				err = db.MemcClient.Set(&memcache.Item{Key: hotelID, Value: []byte(memc_str)})
				if err != nil {
					log.Warn("MMC error: ", err)
			}
		} else{
				log.Errorf("Memcached error = %s\n", err)
				panic(err)
		}
	}
	return ratePlans, nil
}
