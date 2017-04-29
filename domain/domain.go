package domain

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type AdspaceRepository interface {
	Store(adspace Adspace)
	FindById(id string) Adspace
}

type CampaignRepository interface {
	Store(campaign Campaign)
	FindById(id string) Campaign
	FindByAdspaceId(adspace string) []Campaign
}

type CampaignInternalRepository interface {
	Store(campaignInternal CampaignInternal)
	FindById(id string) CampaignInternal
	FindByNetworkId(netrowkId string) []InternalCampaign
}

type PlaceholderTemplateRepository interface {
	Store(placeholderTemplate PlaceholderTemplate)
	FindById(id string) PlaceholderTemplate
}

type AdnetworkRepository interface {
	Store(adnetwork Adnetwork)
	FindById(id string) Adnetwork
}

type Adspace struct {
	Id            string
	Title         string
	Description   string
	Width         int32
	Height        int32
	DeviceType    string
	ContentFormat string
	Position      string
	FloorPrice    float64
}

type Campaign struct {
	Id            string
	Name          string
	ContentPath   string //add content utl
	ContentType   string
	Desription    string
	AdspaceId     string
	FloorPrice    float64
	AdnetworkId   string
	AdnetworkUrl  string //ads content api/url
	AdnetworkType string //ads content type
}

type AdspaceCampaign struct {
	Code     string `json:"code"`
	Iurl     string `json:"iurl"`
	ClickUrl string `json:"clickUrl"`
	PixelUrl string `json:"pixelUrl"`
}

type CampaignInternal struct {
	Id                    string
	Name                  string
	ContentPath           string
	ContentType           string
	ActiveStatus          string
	PlaceholderTemplateId string
	UpdatedDateTime       string
	CreatedDateTime       string
	TargetClick           string
	AdnetworkId           string
}

type PlaceholderTemplate struct {
	Id              string
	AdnetworkId     string
	Width           int32
	Height          int32
	UpdatedDateTime string
	CreatedDateTime string
}

type Adnetwork struct {
	Id              string
	Name            string
	Url             string
	Type            string
	UpdatedDateTime string
	CreatedDateTime string
}

func (camp *Campaign) GetContentUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("Cannot get reply from designated url")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

}

func (adspaceCamp *AdspaceCampaign) GetClickUrl(adTransactionId string, adspaceId string, campaignId string,
	contentType string, placeholderTemplateId string, targetClick string) string {
	var clickUrl string
	contextPath := http.
	return clickUrl
}

/*
func (order *Order) Add(item Item) error {
	if !item.Available {
		return errors.New("Cannot add unavailable items to order")
	}
	if order.value()+item.Value > 250.00 {
		return errors.New(`An order may not exceed
			a total value of $250.00`)
	}
	order.Items = append(order.Items, item)
	return nil
}

func (order *Order) value() (sum float64) {
	for i := range order.Items {
		sum += order.Items[i].Value
	}
	return
}

*/
