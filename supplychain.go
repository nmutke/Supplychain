package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Asset struct {
	TransactionID     	string `json:"TransactionID"`
	TransactionEntity 	string `json:"TransactionEntity"`
	Branch            	string `json:"Branch"`
	Module            	string `json:"Module"`
	TransactionType   	string `json:"TransactionType"`
	PartnerEntity     	string `json:"PartnerEntity"`
	PartnerBranch     	string `json:"PartnerBranch"`
	AssetID           	string `json:"AssetID"`
	OrderID           	string `json:"OrderID"`
	AssetName         	string `json:"AssetName"`
	AssetLocation     	string `json:"AssetLocation"`
	AssetType	  		string `json:"AssetType"`
	BalanceQty	  		string `json:"BalanceQty"`
	Qty               	string `json:"Qty"`
	UOM               	string `json:"UOM"`
	EffectiveDate     	string `json:"EffectiveDate"`
	ExpiryDate        	string `json:"ExpiryDate"`
	ReferenceAsset    	string `json:"ReferenceAsset"`
	ReferenceOrder    	string `json:"ReferenceOrder"`
	Status            	string `json:"Status"`
	AllValues         	string `json:"AllValues"`
	Acknowledgement   	string `json:"Acknowledgement"`
	Transaction	  		string `json:"Transaction"`
}

// AssetPrivateDetails describes details that are private to owners
//type AssetPrivateDetails struct {
//	ID             string `json:"assetID"`
//	AppraisedValue int    `json:"appraisedValue"`
//}

func (s *SmartContract) InsertAssetRecords(ctx contractapi.TransactionContextInterface, AssetData string) (string, error) {

	if len(AssetData) == 0 {
		return "", fmt.Errorf("Please pass the correct Asset data")
	}

	var asset Asset
	err := json.Unmarshal([]byte(AssetData), &asset)
	if err != nil {
		return "", fmt.Errorf("Failed while unmarshling Asset records %s", err.Error())
	}

	assetAsBytes, err := json.Marshal(asset)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling kyc records %s", err.Error())
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(asset.TransactionID, assetAsBytes)

}

func (s *SmartContract) InsertOrderRecords(ctx contractapi.TransactionContextInterface, AssetData string) (string, error) {

	if len(AssetData) == 0 {
		return "", fmt.Errorf("Please pass the correct Asset data")
	}

	var asset Asset
	err := json.Unmarshal([]byte(AssetData), &asset)
	if err != nil {
		return "", fmt.Errorf("Failed while unmarshling Asset records %s", err.Error())
	}

	assetAsBytes, err := json.Marshal(asset)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling kyc records %s", err.Error())
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(asset.TransactionID, assetAsBytes)

}

func (s *SmartContract) GetAssetByTransactionID(ctx contractapi.TransactionContextInterface, TransactionID string) (*Asset, error) {
	if len(TransactionID) == 0 {
		return nil, fmt.Errorf("Please provide correct citizen Id")
	}

	assetAsBytes, err := ctx.GetStub().GetState(TransactionID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if assetAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", TransactionID)
	}

	asset := new(Asset)
	_ = json.Unmarshal(assetAsBytes, asset)

	return asset, nil

}

func (s *SmartContract) GetAssetForQuery(ctx contractapi.TransactionContextInterface, queryString string) ([]Asset, error) {
	queryResults, err := s.getQueryResultForQueryString(ctx, queryString)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	return queryResults, nil

}

func (s *SmartContract) getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]Asset, error) {

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []Asset{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		newAsset := new(Asset)

		err = json.Unmarshal(response.Value, newAsset)
		if err != nil {
			return nil, err
		}

		results = append(results, *newAsset)
	}
	return results, nil
}

func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, transactionID string, transactionentity string ,branch string,module string,transactiontype string,partnerentity string, partnerbranch string, assetID string, orderID string, assetname string,assettype string,assetlocation string,balanceqty string,qty string,uom string ,effectivedate string, expirydate string,referenceasset string,referenceorder string,status string,allvalues string,acknowledgement string,transaction string ) (string, error) {

		if len(transactionID) == 0 {
			return "", fmt.Errorf("Please pass the correct Asset data")
        }
		assetAsBytes, err := ctx.GetStub().GetState(transactionID)

		if err != nil {
			return "", fmt.Errorf("Failed while unmarshling Asset records %s", err.Error())
        }

		if assetAsBytes == nil {
			return "", fmt.Errorf("the asset %s does not exist", transactionID)
		}

		asset := new(Asset)
		_ = json.Unmarshal(assetAsBytes, asset)
		
		asset.TransactionEntity = transactionentity
		asset.Branch = branch
		asset.Module = module
		asset.TransactionType = transactiontype
		asset.PartnerEntity = partnerentity
		asset.PartnerBranch = partnerbranch
		asset.AssetID = assetID
		asset.OrderID = orderID
		asset.AssetName = assetname
		asset.AssetLocation = assetlocation
		asset.AssetType = assettype
		asset.BalanceQty = balanceqty
		asset.Qty = qty
		asset.UOM = uom
		asset.EffectiveDate = effectivedate
		asset.ExpiryDate = expirydate
		asset.ReferenceAsset = referenceasset
		asset.ReferenceOrder = referenceorder
		asset.Status = status
		asset.AllValues = allvalues
		asset.Acknowledgement = acknowledgement
		asset.Transaction =	transaction

		assetAsBytes, err = json.Marshal(asset)
		if err != nil {

			return "", fmt.Errorf("Failed marshal %s", err.Error())
		}

		return  ctx.GetStub().GetTxID(), ctx.GetStub().PutState(asset.TransactionID, assetAsBytes)

}

func (s *SmartContract) DeleteAssetByTransactionId(ctx contractapi.TransactionContextInterface, TransactionID string) (string, error) {
	if len(TransactionID) == 0 {
		return "", fmt.Errorf("Please provide correct contract Id")
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().DelState(TransactionID)
}

func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, TransactionID string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(TransactionID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

func (s *SmartContract) GetHistoryForAsset(ctx contractapi.TransactionContextInterface, carID string) (string, error) {

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(carID)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return string(buffer.Bytes()), nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create Snapchain chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting Snapchain chaincode: %s", err.Error())
	}

}


