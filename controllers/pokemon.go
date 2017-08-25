package controllers

import (
	"api-tutorial01/models"

	"github.com/astaxie/beego"
)

// PokemonController ...
type PokemonController struct {
	beego.Controller
}

// GetAll ...
func (p *PokemonController) GetAll() {
	pokes := models.GetAllPokemon()
	p.Data["json"] = pokes
	p.ServeJSON()
}
