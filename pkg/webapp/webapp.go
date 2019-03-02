package webapp

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"syscall/js"
	"time"

	"github.com/larwef/cert-monitor/pkg/config"

	"github.com/larwef/cert-monitor/pkg/cert"
)

var document = js.Global().Get("document")
var tableBody = js.Global().Get("document").Call("getElementById", "tableBody")

// Webapp object used to run the webapp.
type Webapp struct {
	client *cert.Client
}

// New returns a new Webapp object.
func New(conf *config.Config) *Webapp {
	return &Webapp{
		client: cert.NewClient(conf),
	}
}

// Run runs the webapp.
func (w *Webapp) Run() {
	c := make(chan struct{}, 0)
	log.Println("Started webapp")

	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.searchCerts()
		return nil
	})

	js.Global().Get("document").Call("getElementById", "searchButton").Call("addEventListener", "click", cb)
	<-c
}

type reqWrap struct {
	req *cert.Request
	env string
}

func (w *Webapp) searchCerts() {
	go func() {
		tableBody.Set("innerHTML", "")

		orgNo := js.Global().Get("document").Call("getElementById", "orgNo").Get("value").String()

		log.Printf("OrgNo: %s\n", orgNo)

		reqs := []reqWrap{
			reqWrap{cert.BuypassProdRequest(orgNo), "prod"},
			reqWrap{cert.BuypassTestRequest(orgNo), "test"},
			reqWrap{cert.CommfidesProdRequest(orgNo), "prod"},
			reqWrap{cert.CommfidesTestRequest(orgNo), "test"},
		}

		for _, req := range reqs {
			res, err := w.client.Search(req.req)
			if err != nil {
				log.Printf("Error searching for cert: %v", err)
				return
			}

			for _, elem := range res.Certs {
				addRow(tableBody, elem, req.env)
			}
		}

		return
	}()
}

func addRow(parent js.Value, crt *cert.Cert, env string) {
	tableRow := document.Call("createElement", "tr")

	addCell(tableRow, crt.Organization)
	addCell(tableRow, crt.SerialNumber)
	addCell(tableRow, crt.Issuer)
	addCell(tableRow, crt.ValidFrom.Format(time.RFC3339))
	addCell(tableRow, crt.ValidTo.Format(time.RFC3339))
	addCell(tableRow, strconv.Itoa(int(math.Floor(crt.ValidTo.Sub(time.Now()).Hours()/24))))
	addCell(tableRow, fmt.Sprintf("%t", crt.NonRepudiation))
	addCell(tableRow, env)

	parent.Call("appendChild", tableRow)
}

func addCell(parent js.Value, value string) {
	cell := document.Call("createElement", "td")
	cellVal := document.Call("createTextNode", value)
	cell.Call("appendChild", cellVal)
	parent.Call("appendChild", cell)
}
