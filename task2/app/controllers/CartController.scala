package controllers

import play.api.mvc._
import play.api.libs.json._
import javax.inject._
import scala.collection.mutable.ListBuffer
import models.{CartItem, NewCartItem}

@Singleton
class CartController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {
    private val cartItems = ListBuffer[CartItem]()

    implicit val newCartItemFormat = Json.format[NewCartItem]
    implicit val cartItemFormat = Json.format[CartItem]

    def getAllCartItems = Action {
        if (cartItems.isEmpty) NoContent else Ok(Json.toJson(cartItems))
    }

    def getCartItem(id: Int) = Action {
        cartItems.find(_.id == id) match {
            case Some(item) => Ok(Json.toJson(item))
            case None => NotFound
        }
    }

    def addCartItem = Action(parse.json) { request =>
        request.body.validate[NewCartItem].asOpt.fold(
            BadRequest("No cart item added")
        ) { response =>
            val nextId = if (cartItems.isEmpty) 1 else cartItems.map(_.id).max + 1
            val newItem = CartItem(nextId, response.productId, response.quantity)
            cartItems += newItem
            Created(Json.toJson(newItem))
        }
    }

    def updateCartItem(id: Int) = Action(parse.json) { request =>
        cartItems.find(_.id == id) match {
            case Some(existingItem) =>
                request.body.validate[NewCartItem].asOpt.fold(
                    BadRequest("Invalid cart item data")
                ) { updatedData =>
                    val updatedItem = CartItem(id, updatedData.productId, updatedData.quantity)
                    cartItems -= existingItem
                    cartItems += updatedItem
                    Ok(Json.toJson(updatedItem))
                }
            case None => NotFound(Json.obj("error" -> s"Cart item with ID $id not found"))
        }
    }

    def deleteCartItem(id: Int) = Action {
        cartItems.find(_.id == id) match {
            case Some(item) =>
                cartItems -= item
                Ok(Json.obj("message" -> s"Cart item with ID $id deleted"))
            case None => NotFound(Json.obj("error" -> s"Cart item with ID $id not found"))
        }
    }
}