package handlers

import (
	"go-gin-postgres-template/config"
	"go-gin-postgres-template/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProject(c *gin.Context) {
	var projects []models.Project
	var count int

	rows, err := config.Db.Query("SELECT project_id,title,user_id from projects")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Project not found"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.Project_id, &project.Title, &project.User_id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning user data"})
			return
		}
		projects = append(projects, project)
	}
	err = config.Db.QueryRow("SELECT COUNT(*) FROM projects").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching project count"})
		return
	}

	// Adding data Count
	c.JSON(http.StatusOK, gin.H{"Data Rows": count, "Project": projects})
}

func PostProject(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	project.User_id = userID.(string)

	_, err := config.Db.Exec("INSERT INTO projects (title, description, date_completed,technologies_used,image_url,link,user_id) VALUES ($1,$2,$3,$4::jsonb,$5,$6,$7)", project.Title,
		project.Description,
		project.Date_completed,
		project.Technologies_used,
		project.Image_url,
		project.Link,
		project.User_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Log detailed error
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project Created successfully", "project": project})
}
func UpdateProject(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error json": err.Error()})
		return
	}
	projectID := c.Param("project_id")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID = userID.(string)

	result, err := config.Db.Exec("UPDATE projects SET title=$1, description=$2, date_completed=$3,technologies_used=$4::jsonb,image_url=$5,link=$6 WHERE user_id = $7 AND project_id = $8", project.Title,
		project.Description,
		project.Date_completed,
		project.Technologies_used,
		project.Image_url,
		project.Link,
		userID,
		projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error internal": err.Error()}) // Log detailed error
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found or not authorized to update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project successfully Update", "project": project})
}
func DeleteProject(c *gin.Context) {
	projectID := c.Param("project_id")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	result, err := config.Db.Exec("DELETE FROM projects WHERE project_id = $1 AND user_id = $2", projectID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found or not authorized to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
