package handlers

import (
	"mi-amorito-backend/internal/models"
	"mi-amorito-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CharacterHandler struct {
	characterService *services.CharacterService
}

func NewCharacterHandler(characterService *services.CharacterService) *CharacterHandler {
	return &CharacterHandler{
		characterService: characterService,
	}
}

// GetCharacters 获取角色列表
func (h *CharacterHandler) GetCharacters(c *gin.Context) {
	characters, err := h.characterService.GetActiveCharacters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get characters",
			"code":  "CHARACTERS_FETCH_FAILED",
		})
		return
	}

	response := models.GetCharactersResponse{
		Characters: characters,
	}

	c.JSON(http.StatusOK, response)
}

// GetCharacter 获取单个角色详情
func (h *CharacterHandler) GetCharacter(c *gin.Context) {
	characterID := c.Param("character_id")
	if characterID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "character ID is required",
			"code":  "CHARACTER_ID_REQUIRED",
		})
		return
	}

	character, err := h.characterService.GetCharacterByID(characterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "character not found",
			"code":  "CHARACTER_NOT_FOUND",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"character": character,
	})
}