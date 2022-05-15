package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/milhamh95/checkr/domain"
	"github.com/milhamh95/checkr/graph"
	"github.com/milhamh95/checkr/graph/generated"
	"github.com/milhamh95/checkr/service"
	"github.com/milhamh95/checkr/storage"
)

func main() {
	promoStorage := storage.NewPromoStorage()
	productStorage := storage.NewProductStorage()
	cartStorage := storage.NewCartStorage()

	bundlingPromo := service.NewBundlingPromoOffer(
		productStorage,
		cartStorage,
	)

	bulkDiscountPromo := service.NewBulkDiscountPromoOffer(cartStorage)
	buyXPayYPromo := service.NewBuyXPayYPromoOffer(cartStorage)

	promoOffer := map[domain.PromoType]service.ApplyPromoOffer{
		domain.BulkDiscountPromo: bulkDiscountPromo,
		domain.BundlingPromo:     bundlingPromo,
		domain.BuyXPayYPromo:     buyXPayYPromo,
	}

	promoService := service.NewPromoService(
		productStorage,
		promoStorage,
		promoOffer,
	)

	cartService := service.NewCartService(
		cartStorage,
		productStorage,
		promoService,
	)

	srvResolver := graph.NewResolver(cartService)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: srvResolver,
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := "8080"
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
