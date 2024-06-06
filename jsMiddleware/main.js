const { Client, LocalAuth } = require('whatsapp-web.js');
const qrcode = require('qrcode-terminal');
const palabras = [
    "Serendipia",
    "Melodía",
    "Efémero",
    "Susurro",
    "Atardecer",
    "Estrellas",
    "Bosque",
    "Mariposa",
    "Río",
    "Montaña",
    "Horizonte",
    "Viento",
    "Nube",
    "Aurora boreal",
    "Cascada",
    "Océano",
    "Galaxia",
    "Universo",
    "Imaginación",
    "Sueño",
    "Serendipia",
    "Melodía",
    "Efémero",
    "Susurro",
    "Atardecer",
    "Estrellas",
    "Bosque",
    "Mariposa",
    "Río",
    "Montaña",
    "Horizonte",
    "Viento",
    "Nube",
    "Aurora boreal",
    "Cascada",
    "Océano",
    "Galaxia",
    "Universo",
    "Imaginación",
    "Sueño"
  ];

  async function sendGPTMessage(mensaje) {
    const response = await fetch("http://localhost:8888/baia/askGPT/question", {
        method: 'POST',
        body: JSON.stringify({ // Convert data to JSON string
            "question": mensaje
        }),
        headers: { // Set Content-Type header for JSON data
            'Content-Type': 'application/json'
        }
    });

    if (!response.ok) {
        console.log("Error asking gpt: " + response.status);
        return "Hubo un error"
    } else {
        const responseData = await response.json(); // Parse JSON response
        console.log(responseData["Answer"]); 
        return responseData["Answer"] // Print the parsed JSON data
    }
}

const client = new Client({
    authStrategy: new LocalAuth(),
    webVersion: "2.2412.54",
    webVersionCache: {
        type: "remote",
        remotePath:
            "https://raw.githubusercontent.com/wppconnect-team/wa-version/main/html/2.2412.54.html",
    },
});

client.on('message', async message => {
    console.log(message.from)
    console.log(message.body + "\n \n")
    if(message.from === "5212222150794@c.us"){
        message.reply("caca")
        
    }
    
});

client.on('ready', () => {
    console.log('Client is ready!');
    for(let i=0;i<20;i++){
    client.sendMessage("5212222150794@c.us", "CACA")
    }
});

client.on('qr', qr => {
    qrcode.generate(qr, { small: true });
});

client.initialize();
