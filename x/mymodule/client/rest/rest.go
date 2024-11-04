package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
    r.HandleFunc("/mymodule/something", createSomethingHandler(cliCtx)).Methods("POST")
}

func createSomethingHandler(cliCtx context.CLIContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Implement your handler logic here
    }
}
