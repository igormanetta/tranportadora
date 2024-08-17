package swagger

// swagger:operation POST /veiculo Veiculo postVeiculo
//
// ---
// produces:
// - application/json
// parameters:
// - in: body
//   schema:
//     "$ref": "#/definitions/InsertVeiculo"
// responses:
//   '201':
//     schema:
//       "$ref": "#/definitions/ReturnID"

// swagger:operation GET /veiculo Veiculo listVeiculo
//
// ---
// produces:
// - application/json
// parameters:
// - in: query
// responses:
//   '200':
//     schema:
//       "$ref": "#/definitions/ListVeiculo"

// swagger:operation PUT /veiculo/{id} Veiculo putVeiculo
//
// ---
// produces:
// - application/json
// parameters:
// - in: body
//   schema:
//     "$ref": "#/definitions/UpdateVeiculo"
// responses:
//   '200':
//     schema:
//       "$ref": "#/definitions/ReturnID"

// swagger:operation DELETE /veiculo/{id} Veiculo deleteVeiculo
//
// ---
// produces:
// - application/json
// responses:
//   '204':

// swagger:operation GET /veiculo/{id} Veiculo getVeiculo
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     schema:
//       "$ref": "#/definitions/Veiculo"
