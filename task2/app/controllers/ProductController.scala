package controllers

import play.api.mvc._
import javax.inject._
import play.api.libs.json._
import scala.collection.mutable.ListBuffer

import models.{Product, NewProduct}

@Singleton
class ProductController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {
    private val products = ListBuffer[Product]()

    implicit val newProductFormat = Json.format[NewProduct]
    implicit val productFormat = Json.format[Product]

    def getAllProducts = Action {
        if (products.isEmpty) NoContent else Ok(Json.toJson(products))
    }

    def getProduct(id: Int) = Action {
        products.find(_.id == id) match {
            case Some(product) => Ok(Json.toJson(product))
            case None => NotFound
        }
    }

    def addProduct = Action(parse.json) { request =>
        request.body.validate[NewProduct].asOpt.fold(
            BadRequest("No product added")
        )
        { response =>
            val nextId = if (products.isEmpty) 1 else products.map(_.id).max + 1
            val newProduct = Product(nextId, response.categoryId, response.name, response.price)
            products += newProduct
            Created(Json.toJson(newProduct))
        }
    }

    def updateProduct(id: Int) = Action(parse.json) { request =>
        products.find(_.id == id) match {
            case Some(existingProduct) =>
                request.body.validate[NewProduct].asOpt.fold(
                    BadRequest("Invalid product data")
                ) { updatedData =>
                    val updatedProduct = Product(id, updatedData.categoryId, updatedData.name, updatedData.price)
                    products -= existingProduct
                    products += updatedProduct
                    Ok(Json.toJson(updatedProduct))
                }
            case None =>
                NotFound(Json.obj("error" -> s"Product with ID $id not found"))
        }
    }

    def deleteProduct(id: Int) = Action {
        products.find(_.id == id) match {
        case Some(product) =>
            products -= product
            Ok(Json.obj("message" -> s"Product with ID $id deleted"))
        case None =>
            NotFound(Json.obj("error" -> s"Product with ID $id not found"))
        }
    }
}