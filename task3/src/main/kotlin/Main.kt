import io.ktor.server.application.*
import io.ktor.server.engine.embeddedServer
import io.ktor.server.netty.Netty
import io.ktor.server.response.*
import io.ktor.server.request.*
import io.ktor.server.routing.*
import io.ktor.serialization.kotlinx.json.*
import io.ktor.server.plugins.cors.routing.*
import io.ktor.server.plugins.contentnegotiation.*
import io.ktor.http.ContentType
import kotlinx.coroutines.launch
import io.github.cdimascio.dotenv.dotenv
import kotlinx.serialization.Serializable

data class Category(val name: String)

data class Product(val name: String, val category: String, val price: Double)

val categories = listOf(
    Category("Electronics"),
    Category("Clothing"),
    Category("House")
)

val products = listOf(
    Product("Smartphone", "Electronics", 999.99),
    Product("Laptop", "Electronics", 2449.99),
    Product("T-shirt", "Clothing", 42.00),
    Product("Jacket", "Clothing", 69.69),
    Product("Chair", "House", 84.91)
)

@Serializable
data class MessageRequest(val message: String)

fun main() {
    val dotenv = dotenv()
    val token = dotenv["DISCORD_BOT_TOKEN"]
    val bot = DcBot(token)

    embeddedServer(Netty, port = 8080) {
        module(bot)
    }.start(wait = false)

    kotlinx.coroutines.runBlocking {
        bot.start()
    }
}

fun Application.module(bot: DcBot) {
    install(ContentNegotiation) {
        json()
    }

    routing {
        get("/") {
            call.respondText(
                """
                <h1>Hello there!</h1>
                """.trimIndent(),
                ContentType.Text.Html
            )
        }
        post("/send-message") {
            val channelId = dotenv()["DISCORD_CHANNEL_ID"]
            val request = call.receive<MessageRequest>()
            val message = request.message.takeIf { it.isNotBlank() } ?: "Empty message"
            bot.sendMessage(channelId, message)
            call.respondText("Message send: $message")
        }
    }
}