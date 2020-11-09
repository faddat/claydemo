package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers claydemo-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/claydemo/post", createPostHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/claydemo/post", listPostHandler(cliCtx, "claydemo")).Methods("GET")
		r.HandleFunc("/claydemo/post/{key}", getPostHandler(cliCtx, "claydemo")).Methods("GET")
		r.HandleFunc("/claydemo/post", setPostHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/claydemo/post", deletePostHandler(cliCtx)).Methods("DELETE")

		
}
