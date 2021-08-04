package selector


import (
	"cloud.google.com/go/bigquery"	
	"context"
	"fmt"
	"os"
	"log"
	// secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/api/option"
)

type BigQuery interface {
	AddGoldernPapers(preobject []*bigquery.StructSaver, client *bigquery.Client) (error)
	// DisplayGoldernPapers()
	NewRequestObject() (*bigquery.Client)
	StopClient(Client *bigquery.Client)
	NewDataset(dataset string, client *bigquery.Client)(*bigquery.Dataset)
	CreateDataset(data *bigquery.Dataset)(error) 
	CreateTable(data *bigquery.Dataset, table string)( *bigquery.Table)
}

type History struct {
	Biography string
	Location string
}

func NewObject() BigQuery {
	return &History{}
}

func (prehistory *History) AddGoldernPapers(preobject []*bigquery.StructSaver, client *bigquery.Client) (error){

	historyIndia := prehistory.NewDataset("mughals", client)
	fmt.Println("New Dataset:", historyIndia)
	err := prehistory.CreateDataset(historyIndia); if err != nil {
		log.Fatalln("Something went wrong", err.Error())
		return err 
	}
	
	
	indianempire := prehistory.CreateTable(historyIndia, "India")
	
	fmt.Println("New Table:", indianempire)
	empire := indianempire.Inserter()
	
	// fmt.Println("Insert:", insertpoint)
	err = client.Dataset("mughals").Delete(context.Background()); if err != nil {
		log.Fatalln("Something went wrong", err.Error())
		return err 
	}
	return empire.Put(context.Background(), preobject)
}

// func (prehistory *History) DisplayGoldernPapers()  {
	
// }


func (prehistory *History) NewRequestObject() *bigquery.Client {
	
	_, err := os.Stat("config/" + "urban-octo-sniffle-28d6031a3174.json")
	if os.IsExist(err) {
		fmt.Println("File Doesn't exist...", err)
		return &bigquery.Client{}
	}

	// cred, err := secretmanager.NewClient(context.Background(),
		
	// if err != nil {
	// 	return &bigquery.Client{}
	// }

	bigClient, err := bigquery.NewClient(context.Background(), "urban-octo-sniffle", option.WithCredentialsFile("config/"+"urban-octo-sniffle-28d6031a3174.json")); if err != nil {
		fmt.Println("Client:", err.Error())
		return &bigquery.Client{}  
	}
	return bigClient
}

func (prehistory *History) StopClient(Client *bigquery.Client)  {
	defer Client.Close()
}

func (prehistory *History) NewDataset(dataset string, client *bigquery.Client) (*bigquery.Dataset) {
	return client.Dataset(dataset)
}

func (prehistory *History) CreateDataset(data *bigquery.Dataset)(error) {
	return data.Create(context.Background(), &bigquery.DatasetMetadata{Location: "EU"})
}

func (prehistory *History) CreateTable(data *bigquery.Dataset, table string)( *bigquery.Table) {
	return data.Table(table)
}