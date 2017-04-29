package usecases

import (
	"domain"
	"fmt"
	"math/rand"
	"sort"
)

type Logger interface {
	Log(message string) error
}

func (a []domain.Campaign) Len() int           { return len(a) }
func (a []domain.Campaign) Less(i, j int) bool { return a[i].FloorPrice < a[j].FloorPrice }
func (a []domain.Campaign) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type CampaignInteractor struct {
	AdspaceRepository          domain.AdspaceRepository
	AdspaceCampaign			   domain.AdspaceCampaign
	CampaignRepository         domain.CampaignRepository
	CampaignInternalRepository domain.CampaignInternalRepository
	Logger                     Logger
}

func (interactor *CampaignInteractor) GetCampaignUrl(adspaceId string) (CampaignInteractor.AdspaceCampaign, error) {
	var finalCampaigns []domain.Campaign
	var finalCampaign domain.Campaign
	adspaceCampaign := CampaignInteractor.AdspaceCampaign{}
	adspace := interactor.AdspaceRepository.FindById(adspaceId)
	if adspace.Id == nil {
		message := "Adspace #%s is not available"
		err := fmt.Errorf(message, adspaceId)
		interactor.Logger.Log(err.Error())
		return nil, err
	}
	campaigns := interactor.CampaignRepository.FindByAdspaceId(adspaceId)
	if (domain.Campaign{}) == campaign {
		message := "Can't find any suitable campaign for AdspaceId : #%s"
		err := fmt.Errorf(message, adspaceId)
		interactor.Logger.Log(err.Error())
		return nil, err
	}
	//sort result
	sort.Sort(campaigns)
	//check highest floorprice from available campaigns
	highestPrice := campaigns[0].FloorPrice
	idx := 0
	curIdx []int
	for _, camp := range campaigns {
		if camp.FloorPrice >= highestPrice {
			append(finalCampaigns, camp)
			append(curIdx, idx) 
		}
		idx++
	}

	//decide by random if it has more than 1 result
	if len(finalCampaigns) == 1 {
		finalCampaign = finalCampaigns[0]
		campaigns = append(campaigns[:curIdx[0]], campaigns[curIdx[0]+1:]...)
	} else {
		randIdx = rand.Intn(len(finalCampaigns))
		finalCampaign = finalCampaigns[randIdx]
		campaigns = append(campaigns[:curIdx[randIdx]], campaigns[curIdx[randIdx]+1:]...)
	}

	//get url, check if it is internal or outside
	//if internal, get internal campaign url, if outside hit it to get ads content url
	if finalCampaign.AdnetworkType == "INTERNAL"{
		internalCampaigns := interactor.CampaignInternalRepository.FindByNetworkId(finalCampaign.AdnetworkId)
		if len(internalCampaigns) == 1 {
			//return internalCampaigns[0].ContentPath, nil
		} else if len(internalCampaign) > 1{
			//return internalCampaigns[rand.Intn(len(internalCampaigns))].ContentPath, nil
		} else {
			//repeat the process
		}
	} else {
		//hit DSP URL then get the return
		resp, err := http.Get(finalCampaign.Url)
		if err != nil {
			message := "Can't get any reply from Adnetwork URL for Campaign id #%s"
			err := fmt.Errorf(message, finalCampaign.Id)
			interactor.Logger.Log(err.Error())
		} else {
			//we need to parse JSON to object
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			adspaceCampaign := CampaignInteractor.AdspaceCampaign{}
			json.Unmarshal(byte[](body), &adspaceCampaign)
			if adspaceCampaign.ClickUrl != nil {
				return adspaceCampaign, nil
			}
		}

	}

	return nil, nil

}
