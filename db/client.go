package db

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/rwxd/new-newt/utils"
)

var ctx = context.Background()

type RedisClient struct {
	connection *redis.Client
}

func (r RedisClient) GetDomainsToCheck() ([]string, error) {
	log.Println("Getting all items in key \"domains_to_check\"")
	domains, err := r.connection.LRange(ctx, "domains_to_check", 0, -1).Result()
	if err == redis.Nil {
		log.Println("No domains to check")
		return []string{}, nil
	} else if err != nil {
		log.Fatal(err)
	}
	return domains, nil
}

func (r RedisClient) AddDomainToCheck(domain string) error {
	return r.AddDomainsToCheck([]string{domain})
}

func (r RedisClient) AddDomainsToCheck(domains []string) error {
	exists, err := r.connection.Exists(ctx, "domains_to_check").Result()
	if err != nil {
		return err
	}
	if exists == 1 {
		r.connection.LPush(ctx, "domains_to_check", domains)
		return nil
	}
	_, err = r.connection.LPush(ctx, "domains_to_check", domains).Result()
	return err
}

func (r RedisClient) DeleteDomainToCheck(domain string) error {
	domains, err := r.GetDomainsToCheck()
	if err != nil {
		return err
	}
	for i, d := range domains {
		if d == domain {
			log.Printf("Deleting \"%s\" at index %d\n", domain, i)
			_, err = r.connection.LRem(ctx, "domains_to_check", int64(i), d).Result()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("domain \"%s\" not found in list", domain)
}

func (r RedisClient) ClearDomains() error {
	for _, pattern := range []string{"domains_to_check*", "domain_status_*"} {
		keys, err := r.connection.Keys(ctx, pattern).Result()
		if err != nil {
			return err
		}
		for _, key := range keys {
			exists, err := r.connection.Exists(ctx, key).Result()
			if err != nil {
				return err
			}
			if exists == 1 {
				log.Printf("Deleting key \"%s\"\n", key)
				_, err := r.connection.Del(ctx, key).Result()
				if err != nil {
					return err
				}
			}
		}

	}
	return nil
}

func (r RedisClient) SetDomainStatus(domain string, status DomainStatus) (err error) {
	for key, value := range status.ToHashMap() {
		_, err = r.connection.HSet(ctx, fmt.Sprintf("domain_status_%s", domain), key, value).Result()
		if err != nil {
			return
		}
	}
	return
}

func (r RedisClient) GetDomainStatus(domain string) (status DomainStatus, err error) {
	result, err := r.connection.HGetAll(ctx, fmt.Sprintf("domain_status_%s", domain)).Result()
	if err != nil {
		return DomainStatus{}, err
	}
	status, err = NewDomainStatusFromHashMap(result)
	return
}

func NewRedisClient() RedisClient {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return RedisClient{
		connection: client,
	}
}

var Rdb RedisClient = NewRedisClient()
