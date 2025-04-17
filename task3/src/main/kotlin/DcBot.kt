import dev.kord.common.entity.Snowflake
import dev.kord.core.Kord
import dev.kord.core.entity.channel.MessageChannel
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent

class DcBot(private val token: String) {
    private lateinit var client: Kord

    suspend fun start() {
        client = Kord(token)

        client.on<MessageCreateEvent> {
            if (message.content == "!hello") {
                message.channel.createMessage("Hey Hi Hello!")
                return@on
            }

            if (message.content == "!categories") {
                val categoryList = categories.joinToString("\n") { "- ${it.name}" }
                message.channel.createMessage("Categories:\n$categoryList")
                return@on
            }

            if (message.content.startsWith("!products ")) {
                val category = message.content.removePrefix("!products ").trim()
                val filteredProducts = products.filter { it.category.equals(category, ignoreCase = true) }
                val productList = filteredProducts.joinToString("\n") { 
                    "- ${it.name} ${it.price}" 
                }
                message.channel.createMessage("Products in $category:\n$productList")
                return@on
            }
        }

        client.login {
            @OptIn(PrivilegedIntent::class)
            intents += Intent.MessageContent
        }
    }

    suspend fun sendMessage(channelId: String, content: String) {
        val channel = client.getChannel(Snowflake(channelId)) as? MessageChannel
        channel?.createMessage(content)
    }
}