package controllers

import play.api.mvc._
import play.api.libs.json._
import javax.inject._
import scala.collection.mutable.ListBuffer

import models.{Category, NewCategory}

@Singleton
class CategoryController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {
    private val categories = ListBuffer[Category]()

    implicit val newCategoryFormat = Json.format[NewCategory]
    implicit val categoryFormat  = Json.format[Category]

    def getAllCategories: Action[AnyContent] = Action {
        if (categories.isEmpty) NoContent else Ok(Json.toJson(categories))
    }

    def getCategory(id: Int) = Action {
        categories.find(_.id == id) match {
            case Some(category) => Ok(Json.toJson(category))
            case None => NotFound
        }
    }

    def addCategory = Action(parse.json) { request =>
    request.body.validate[NewCategory].asOpt.fold(
            BadRequest("No category added")
        ) { response =>
            val nextId = if (categories.isEmpty) 1 else categories.map(_.id).max + 1
            val newCategory = Category(nextId, response.name)
            categories += newCategory
            Created(Json.toJson(newCategory))
        }
    }

    def updateCategory(id: Int) = Action(parse.json) { request =>
        categories.find(_.id == id) match {
            case Some(existingCategory) =>
                request.body.validate[NewCategory].asOpt.fold(
                    BadRequest("Invalid category data")
                ) { updatedData =>
                    val updatedCategory = Category(id, updatedData.name)
                    categories -= existingCategory
                    categories += updatedCategory
                    Ok(Json.toJson(updatedCategory))
                }
            case None => NotFound(Json.obj("error" -> s"Category with ID $id not found"))
        }
    }

    def deleteCategory(id: Int): Action[AnyContent] = Action {
        categories.find(_.id == id) match {
            case Some(category) =>
                categories -= category
                Ok(Json.obj("message" -> s"Category with ID $id deleted"))
            case None =>
                NotFound(Json.obj("error" -> s"Category with ID $id not found"))
        }
    }
}