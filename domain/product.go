package domain

import (
	"context"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Price struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type Stock struct {
	Available int64 `json:"available"`
	Reserved  int64 `json:"reserved"`
}

type Image struct {
	Url string `json:"url"`
	Alt string `json:"alt"`
}

type Specification struct {
	Weight          string `json:"weight"`
	Dimensions      string `json:"dimensions"`
	BatteryLife     string `json:"battery_life"`
	WaterResistance string `json:"water_resistance"`
}

type Product struct {
	ProductId      string        `json:"product_id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Price          Price         `json:"price"`
	Category       string        `json:"category"`
	Brand          string        `json:"brand"`
	Stock          Stock         `json:"stock"`
	Sku            string        `json:"sku"`
	Tags           []string      `json:"tags"`
	Images         []Image       `json:"images"`
	Specifications Specification `json:"specifications"`
	CreatedAt      string        `json:"created_at"`
	UpdatedAt      string        `json:"updated_at"`
	Index          string        `json:"index"`
	StoreId        string        `json:"store_id"`
}

func GenerateNewProducts(ctx context.Context, cnt int, productOutCh chan<- *Product) {
	f := gofakeit.New(0)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < cnt; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				p := &Product{
					f.ProductSuffix(),
					f.ProductName(),
					f.ProductDescription(),
					Price{
						int64(f.Number(1000, 10000)),
						"RUB",
					},
					f.ProductCategory(),
					f.Name(),
					Stock{
						int64(f.Number(100, 10000)),
						int64(f.Number(1, 1000)),
					},
					f.ProductSuffix(),
					[]string{f.Name(), f.Name(), f.Name()},
					[]Image{
						Image{
							"http://localhost",
							"alt text",
						},
					},
					Specification{
						f.ProductFeature(),
						f.ProductDimension(),
						f.ProductBenefit(),
						f.ProductUPC(),
					},
					time.Now().String(),
					time.Now().String(),
					f.Zip(),
					f.Zip(),
				}
				productOutCh <- p
			}
		}
	}()
	wg.Wait()
}
