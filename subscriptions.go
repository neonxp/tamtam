package tamtam

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/neonxp/tamtam/schemes"
)

type subscriptions struct {
	client *client
}

func newSubscriptions(client *client) *subscriptions {
	return &subscriptions{client: client}
}

//GetSubscriptions returns list of all subscriptions
func (a *subscriptions) GetSubscriptions() (*schemes.GetSubscriptionsResult, error) {
	result := new(schemes.GetSubscriptionsResult)
	values := url.Values{}
	body, err := a.client.request(http.MethodGet, "subscriptions", values, nil)
	if err != nil {
		return result, err
	}
	defer func() {
		if err := body.Close(); err != nil {
			log.Println(err)
		}
	}()
	return result, json.NewDecoder(body).Decode(result)
}

//Subscribe subscribes bot to receive updates via WebHook
func (a *subscriptions) Subscribe(subscribeURL string, updateTypes []string) (*schemes.SimpleQueryResult, error) {
	subscription := &schemes.SubscriptionRequestBody{
		Url:         subscribeURL,
		UpdateTypes: updateTypes,
		Version:     a.client.version,
	}
	result := new(schemes.SimpleQueryResult)
	values := url.Values{}
	body, err := a.client.request(http.MethodPost, "subscriptions", values, subscription)
	if err != nil {
		return result, err
	}
	defer func() {
		if err := body.Close(); err != nil {
			log.Println(err)
		}
	}()
	return result, json.NewDecoder(body).Decode(result)
}

//Unsubscribe unsubscribes bot from receiving updates via WebHook
func (a *subscriptions) Unsubscribe(subscriptionURL string) (*schemes.SimpleQueryResult, error) {
	result := new(schemes.SimpleQueryResult)
	values := url.Values{}
	values.Set("url", subscriptionURL)
	body, err := a.client.request(http.MethodDelete, "subscriptions", values, nil)
	if err != nil {
		return result, err
	}
	defer func() {
		if err := body.Close(); err != nil {
			log.Println(err)
		}
	}()
	return result, json.NewDecoder(body).Decode(result)
}
