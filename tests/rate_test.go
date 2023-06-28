package easypost_test

import (
	"encoding/json"
	easypost2 "github.com/saltboxteam/easypost-go/v2"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"

	"github.com/EasyPost/easypost-go/v2"
)

func (c *ClientTests) TestRateRetrieve() {
	client := c.TestClient()
	assert, require := c.Assert(), c.Require()

	shipment, err := client.CreateShipment(c.fixture.BasicShipment())
	require.NoError(err)

	rate, err := client.GetRate(shipment.Rates[0].ID)
	require.NoError(err)

	assert.Equal(reflect.TypeOf(&easypost.Rate{}), reflect.TypeOf(rate))
	assert.True(strings.HasPrefix(rate.ID, "rate_"))
}

func (c *ClientTests) TestBetaStatelessRateRetrieve() {
	client := c.TestClient()
	assert, require := c.Assert(), c.Require()

	rates, err := client.BetaGetStatelessRates(c.fixture.BasicShipment())
	require.NoError(err)

	assert.Equal(reflect.TypeOf([]*easypost.StatelessRate{}), reflect.TypeOf(rates))
}

func (c *ClientTests) TestBetaStatelessRateGetLowest() {
	client := c.TestClient()
	assert, require := c.Assert(), c.Require()

	rates, err := client.BetaGetStatelessRates(c.fixture.BasicShipment())
	require.NoError(err)

	lowestRate, err := client.LowestStatelessRate(rates)
	require.NoError(err)

	assert.Equal(reflect.TypeOf(easypost.StatelessRate{}), reflect.TypeOf(lowestRate))
	assert.Equal("First", lowestRate.Service)
}

func (c *ClientTests) TestBetaStatelessRateGetLowestError() {
	client := c.TestClient()
	require := c.Require()

	rates, err := client.BetaGetStatelessRates(c.fixture.BasicShipment())
	require.NoError(err)

	// Bad carrier
	_, err = client.LowestStatelessRateWithCarrier(rates, []string{"BadCarrier"})
	require.Error(err)

	// Bad service
	_, err = client.LowestStatelessRateWithCarrierAndService(rates, []string{"USPS"}, []string{"BadService"})
	require.Error(err)
}

func TestSmartRateUpsDeliveryDateJson(t *testing.T) {
	smartRateUpsDateJson := `
{
                "carrier": "UPS",
                "carrier_account_id": "ca_a4be3a46b7bc4657ab2f6c898e1d7f6e",
                "created_at": "2023-06-28T14:52:39Z",
                "currency": "USD",
                "delivery_date": "Thu, 29 Jun 2023 10:30:00 GMT",
                "delivery_date_guaranteed": true,
                "delivery_days": 1,
                "est_delivery_days": 1,
                "id": "rate_7f560e1598024218a94aad5eca72851e",
                "list_currency": "USD",
                "list_rate": 70.0,
                "mode": "production",
                "object": "Rate",
                "rate": 45.57,
                "retail_currency": "USD",
                "retail_rate": 71.31,
                "service": "NextDayAir",
                "shipment_id": "shp_087a68cd07074e9da840971ced82cca5",
                "updated_at": "2023-06-28T14:52:39Z"
            }
`

	var smartRateUpsDate easypost2.SmartRate
	err := json.Unmarshal([]byte(smartRateUpsDateJson), &smartRateUpsDate)
	assert.NoError(t, err)

	smartRateStandardDateJson := `
{
                "carrier": "UPS",
                "carrier_account_id": "ca_a4be3a46b7bc4657ab2f6c898e1d7f6e",
                "created_at": "2023-06-28T14:52:39Z",
                "currency": "USD",
                "delivery_date": "2023-06-28T14:52:39Z",
                "delivery_date_guaranteed": true,
                "delivery_days": 1,
                "est_delivery_days": 1,
                "id": "rate_7f560e1598024218a94aad5eca72851e",
                "list_currency": "USD",
                "list_rate": 70.0,
                "mode": "production",
                "object": "Rate",
                "rate": 45.57,
                "retail_currency": "USD",
                "retail_rate": 71.31,
                "service": "NextDayAir",
                "shipment_id": "shp_087a68cd07074e9da840971ced82cca5",
                "updated_at": "2023-06-28T14:52:39Z"
            }
`

	var smartRateStandardDate easypost2.SmartRate
	err = json.Unmarshal([]byte(smartRateStandardDateJson), &smartRateStandardDate)
	assert.NoError(t, err)
}
