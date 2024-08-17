package swagger

// swagger:operation POST /motorista Motorista postMotorista
//
// ---
// produces:
// - application/json
// parameters:
// - in: body
//   schema:
//     "$ref": "#/definitions/InsertMotorista"
// responses:
//   '201':
//     schema:
//       "$ref": "#/definitions/ReturnID"

// swagger:operation GET /motorista Motorista listMotorista
//
// ---
// produces:
// - application/json
// parameters:
// - in: query
// responses:
//   '200':
//     schema:
//       "$ref": "#/definitions/ListMotorista"

// swagger:operation PUT /motorista/{id} Motorista putMotorista
//
// ---
// produces:
// - application/json
// parameters:
// - in: body
//   schema:
//     "$ref": "#/definitions/UpdateMotorista"
// responses:
//   '200':
//     schema:
//       "$ref": "#/definitions/ReturnID"

// swagger:operation DELETE /motorista/{id} Motorista deleteMotorista
//
// ---
// produces:
// - application/json
// responses:
//   '204':

// swagger:operation GET /motorista/{id} Motorista getMotorista
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     schema:
//       "$ref": "#/definitions/Motorista"

// swagger:operation PATCH /motorista/{motoristaId}/veiculo/{veiculoId} Motorista setMotoristaVeiculo
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     schema:
//       "$ref": "#/definitions/ReturnID"
